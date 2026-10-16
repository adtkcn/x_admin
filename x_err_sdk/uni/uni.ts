import type {
  LogWithError,
  LogWithEnv,
  ListenCallbackFn,
  IErrorEvent,
} from "../types";

interface LoggerProps {}

// uniapp 平台错误采集实现：
// 存储用 uni.*StorageSync，上报用 uni.request（GET beacon 同款语义），
// 错误监听通过包装全局 App 注入 onError / onUnhandledRejection 生命周期。
class Uni implements IErrorEvent {
  props: LoggerProps;
  private _origApp: any = null;
  private _hooked = false;

  constructor(props?: LoggerProps) {
    this.props = {
      ...props,
    };
  }

  public upload(url: string, data: any): Promise<void> {
    return new Promise((resolve) => {
      try {
        uni.request({
          url: url + "?data=" + encodeURIComponent(JSON.stringify(data)),
          method: "GET",
          success: () => resolve(),
          fail: () => resolve(),
        });
      } catch (error) {
        resolve();
      }
    });
  }

  public setCache(key: string, info: any): void {
    try {
      uni.setStorageSync(key, JSON.stringify(info));
    } catch (error) {}
  }

  public getCache(key: string): any {
    try {
      let list = uni.getStorageSync(key);
      if (list) {
        return JSON.parse(list);
      } else {
        return null;
      }
    } catch (error) {
      return null;
    }
  }

  public delCache(key: string): void {
    try {
      uni.removeStorageSync(key);
    } catch (error) {}
  }

  public getEnvInfo(): LogWithEnv {
    const env: LogWithEnv = {
      height: 0,
      width: 0,
    };
    try {
      const info = uni.getSystemInfoSync();
      env.height = info.windowHeight || info.screenHeight || 0;
      env.width = info.windowWidth || info.screenWidth || 0;
    } catch (error) {}
    return env;
  }

  private callback(err: LogWithError): void {}

  // 当前页面路由，作为 path
  private getPath(): string {
    try {
      const pages = getCurrentPages();
      const page = pages[pages.length - 1];
      return page?.route || "";
    } catch (error) {
      return "";
    }
  }

  private handleStack(stack: string): string {
    let newStack: string[] = [];
    if (stack) {
      stack.split("\n").map((item, index) => {
        if (index < 4) {
          newStack.push(item);
        }
      });
    }
    return newStack.join("\n");
  }

  private onError = (err: any): void => {
    if (err && String(err.message || err).includes("ResizeObserver")) {
      return;
    }
    const message = typeof err === "string" ? err : err?.message || "";
    const stack =
      typeof err === "string" ? "" : this.handleStack(err?.stack || "");
    this.callback({
      type: "error",
      event_type: "onError",
      path: this.getPath(),
      message,
      stack,
      ...this.getEnvInfo(),
    });
  };

  private onUnhandledRejection = (res: any): void => {
    const reason = res?.reason ?? res;
    if (typeof reason === "string") {
      this.callback({
        type: "error",
        event_type: "unhandledrejection",
        path: this.getPath(),
        message: reason,
        stack: "",
        ...this.getEnvInfo(),
      });
    } else {
      this.callback({
        type: "error",
        event_type: "unhandledrejection",
        path: this.getPath(),
        message: reason?.message || "",
        stack: this.handleStack(reason?.stack || ""),
        ...this.getEnvInfo(),
      });
    }
  };

  // 包装全局 App，把 SDK 的处理器注入到应用生命周期
  private hookApp(): void {
    if (this._hooked || typeof App !== "function") return;
    this._hooked = true;
    this._origApp = App;
    const self = this;
    // @ts-ignore 全局构造函数重赋值
    App = function (options: any) {
      options = options || {};
      const origOnError = options.onError;
      const origOnUnhandledRejection = options.onUnhandledRejection;
      options.onError = function (err: any) {
        self.onError(err);
        if (typeof origOnError === "function") origOnError.call(this, err);
      };
      options.onUnhandledRejection = function (res: any) {
        self.onUnhandledRejection(res);
        if (typeof origOnUnhandledRejection === "function")
          origOnUnhandledRejection.call(this, res);
      };
      // @ts-ignore 调用原始 App
      return self._origApp(options);
    };
  }

  public listen(callback: ListenCallbackFn): void {
    this.callback = callback;
    this.hookApp();
  }

  public unListen(): void {
    this.callback = () => {};
    if (this._hooked && this._origApp) {
      // @ts-ignore 还原全局 App
      App = this._origApp;
      this._hooked = false;
    }
  }
}

export default Uni;
