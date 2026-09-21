/**
 * 升级 SDK 核心流程（vue3 端）：检查 → 判定 → 下载 → 安装
 * flow() 为平台无关决策逻辑；plus.* 为 5+ Runtime 专用 API，仅 uni-app vue3 可用，
 * 集中于本文件桥接层，uni-app x / react-native 各自实现桥接（见 sdk/upgrade/README.md）
 */
import { config, state, Phase, registerCheckHandler } from "./config";
import type { ApiEnvelope, CheckUpdateResult, ClientAppInfo } from "./types";

const SKIP_KEY = "upgrade_skipped_version";

// 当前下载任务（模块级，供关闭/跳过时中断）
let downloadTask: UniApp.DownloadTask | null = null;

/** 拼接绝对下载地址（后端返回的 /api/uploads/:id/:name 等为相对路径） */
export function absUrl(url: string): string {
  if (/^(https?:|itms-services:)/.test(url)) return url;
  return `${config.baseUrl}${url}`;
}

/** 读取客户端应用信息（包名、version code、wgt 资源版本 name） */
function getClientAppInfo(): Promise<ClientAppInfo> {
  return new Promise((resolve, reject) => {
    // #ifdef APP-PLUS
    const os = uni.getSystemInfoSync().platform; // ios / android
    plus.runtime.getProperty(plus.runtime.appid, (widget) => {
      // plus 类型声明不全，packagename 为运行时实际存在的字段，统一转 any 访问
      const info = widget as any;
      // plus.runtime.version 为当前运行的 wgt 资源版本 name（如 1.0.3），上报给服务端按段比较
      resolve({
        bundle_id: String(info.packagename ?? ""),
        platform: os,
        version_code: Number(info.versionCode ?? 0),
        wgt_version: String(plus.runtime.version ?? ""),
      });
    });
    return;
    // #endif
    // #ifndef APP-PLUS
    reject(new Error("[upgrade] 仅支持 App 平台"));
    // #endif
  });
}

/** 请求检查更新接口（公开接口，无需登录） */
function requestCheckUpdate(params: ClientAppInfo): Promise<CheckUpdateResult> {
  return new Promise((resolve, reject) => {
    uni.request({
      url: `${config.baseUrl}${config.checkPath}`,
      method: "GET",
      data: {
        bundle_id: params.bundle_id,
        platform: params.platform,
        version_code: params.version_code,
        wgt_version: params.wgt_version,
      },
      success: (res) => {
        const body = res.data as ApiEnvelope<CheckUpdateResult>;
        if (body == null || body.code !== 0) {
          reject(new Error(body?.message ?? "检查更新失败"));
        } else {
          resolve(body.data);
        }
      },
      fail: (err) => reject(new Error(err.errMsg)),
    });
  });
}

/**
 * 判定更新（平台无关决策，react-native 移植时可直接复用本函数语义）：
 * 1. 无更新 → 隐藏
 * 2. wgt 更新 → 静默下载安装，不弹窗（仅 android；iOS 不允许热更自己的 JS）
 * 3. 全量包更新 → 弹窗；普通模式用户可选择"以后再说"（同版本不再提示）
 */
export function flow(
  result: CheckUpdateResult,
  client: ClientAppInfo,
): "dialog" | "wgt" | "none" {
  if (!result.update) return "none";

  if (result.type === "wgt" && result.wgt != null) {
    // uni-app x 不支持 wgt；iOS 全量包走商店/OTA，同样跳过热更
    if (!config.enableWgt || client.platform !== "android") return "none";
    state.info = {
      force: true, // wgt 静默流程无用户确认环节
      isWgt: true,
      version: result.wgt.version,
      version_code: result.wgt.version_code,
      download_url: result.wgt.download_url,
      install_url: "",
    };
    return "wgt";
  }

  if (result.type === "app" && result.app != null) {
    // 普通模式下用户曾跳过该版本则不再提示；强制模式始终提示
    const force = result.app.update_mode === 1;
    if (
      !force &&
      uni.getStorageSync(SKIP_KEY) === String(result.app.version_code)
    ) {
      return "none";
    }
    state.info = {
      force,
      isWgt: false,
      version: result.app.version,
      version_code: result.app.version_code,
      download_url: result.app.download_url,
      install_url: result.app.install_url,
    };
    return "dialog";
  }
  return "none";
}

/** 失败处理：静默流程仅复位；弹窗流程回到确认阶段允许重试 */
function handleFail(msg: string, silent: boolean) {
  console.warn("[upgrade]", msg);
  if (silent) {
    close();
    return;
  }
  state.hint = msg;
  state.phase = Phase.CONFIRM;
}

/** 安装升级包：wgt 静默安装后提示重启；apk 拉起系统安装器 */
function install(filePath: string, isWgt: boolean, silent: boolean) {
  // #ifdef APP-PLUS
  const rt = plus.runtime as any;
  if (isWgt) {
    rt.install(
      filePath,
      { force: true },
      () => {
        // 热更安装完成，重启生效（唯一需要用户感知的环节）
        uni.showModal({
          title: "升级完成",
          content: "热更新已安装，重启后生效",
          confirmText: "立即重启",
          cancelText: "稍后",
          showCancel: true,
          success: (r) => {
            if (r.confirm) rt.restart();
            else close();
          },
        });
      },
      (err: any) => handleFail(`热更新安装失败(${err?.code ?? ""})`, silent),
    );
    return;
  }
  // android 全量包：拉起系统安装器，success 仅代表安装器启动；
  // 安装器被用户取消时停留在 READY 阶段，可再次点击（重走 accept）
  rt.install(
    filePath,
    { force: true },
    () => {
      if (!silent) state.phase = Phase.READY;
    },
    (err: any) => handleFail(`安装失败(${err?.code ?? ""})`, silent),
  );
  // #endif
}

/** 下载升级包（android 全量 apk / wgt），进度写入共享状态；silent 为 wgt 静默模式 */
function startDownload(silent: boolean) {
  const info = state.info;
  if (info == null) return;
  if (!silent) {
    state.phase = Phase.DOWNLOADING;
    state.progress = 0;
    state.hint = "";
  }
  downloadTask = uni.downloadFile({
    url: absUrl(info.download_url),
    success: (res) => {
      downloadTask = null;
      if (res.statusCode !== 200 || res.tempFilePath == null) {
        handleFail(`下载失败(${res.statusCode})`, silent);
        return;
      }
      install(res.tempFilePath, info.isWgt, silent);
    },
    fail: () => {
      downloadTask = null;
      handleFail("下载失败，请检查网络", silent);
    },
  });
  downloadTask.onProgressUpdate((p) => {
    state.progress = p.progress;
  });
}

/**
 * 执行一次检查（组件 setup 时注册为共享 handler）
 * wgt 走静默下载安装；全量包进入弹窗确认阶段
 */
export async function runCheck() {
  try {
    state.hint = "";
    const client = await getClientAppInfo();
    if (client.bundle_id === "") throw new Error("无法获取应用包名");
    const result = await requestCheckUpdate(client);
    const action = flow(result, client);
    if (action === "wgt") {
      // 热更包静默下载安装，不展示弹窗
      startDownload(true);
    } else if (action === "dialog") {
      state.phase = Phase.CONFIRM;
    }
  } catch (e: any) {
    // 启动自动检查场景失败不打扰用户
    console.warn("[upgrade] 检查更新失败:", e?.message ?? e);
    state.phase = Phase.HIDDEN;
  }
}

/** 用户同意升级：android 进入下载；iOS 打开商店/OTA */
export function accept() {
  // #ifdef APP-IOS
  openIosStore();
  // #endif
  // #ifdef APP-ANDROID
  startDownload(false);
  // #endif
}

/** 下载就绪后重新拉起安装（安卓安装器被取消时的重试入口，等价于 accept） */
export function installReady() {
  // #ifdef APP-ANDROID
  startDownload(false);
  // #endif
}

/** 打开 iOS 安装入口：优先 itms-services OTA 地址，缺省回外置浏览器打开下载地址 */
function openIosStore() {
  const info = state.info;
  if (info == null) return;
  // #ifdef APP-PLUS
  const url = info.install_url !== "" ? info.install_url : info.download_url;
  plus.runtime.openURL(absUrl(url));
  if (!info.force) close(); // 强制模式停留在弹窗，用户从商店返回后可再次点击
  // #endif
}

/** 用户跳过（仅普通模式）：记录该版本 code，同版本不再提示 */
export function refuse() {
  if (state.info != null) {
    uni.setStorageSync(SKIP_KEY, String(state.info.version_code));
  }
  close();
}

/** 关闭弹窗并中断下载 */
export function close() {
  downloadTask?.abort();
  downloadTask = null;
  state.info = null;
  state.phase = Phase.HIDDEN;
  state.progress = 0;
  state.hint = "";
}

/** 弹窗遮罩点击：强制模式或下载中禁止关闭 */
export function onMaskTap() {
  const info = state.info;
  if (info == null) return;
  if (info.force || state.phase === Phase.DOWNLOADING) return;
  refuse();
}

/** 组件 setup 时注册检查入口，供 config.check() 调用 */
export function setupCheck() {
  registerCheckHandler(runCheck);
}
