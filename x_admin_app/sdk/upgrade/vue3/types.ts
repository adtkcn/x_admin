/**
 * 升级 SDK 协议类型定义（与后端 fabu CheckUpdate 接口契约一一对应）
 * 平台无关：react-native 等新端接入时直接复用本文件的类型语义
 */

/** 全量安装包更新信息 */
export interface AppUpdateInfo {
	version: string;
	version_code: number;
	download_url: string;
	install_url: string;
	update_mode: number; // 0普通 1强制（2静默后端已废弃，客户端按普通处理）
}

/** wgt 热更新包更新信息（仅 uni-app vue3 支持，uni-app x 不支持 wgt） */
export interface WgtUpdateInfo {
	version: string;
	version_code: number;
	download_url: string;
	md5: string;
	size: number;
}

/**
 * CheckUpdate 接口返回，三种固定形态：
 * {"update":true,"type":"app","app":{...}} / {"update":true,"type":"wgt","wgt":{...}} / {"update":false}
 */
export interface CheckUpdateResult {
	update: boolean;
	type?: string; // "app" | "wgt"
	app?: AppUpdateInfo;
	wgt?: WgtUpdateInfo;
}

/** 后端统一响应信封 */
export interface ApiEnvelope<T> {
	code: number;
	message?: string;
	data: T;
}

/** 当前客户端应用信息（检查更新的请求参数） */
export interface ClientAppInfo {
	bundle_id: string;
	platform: string; // ios / android
	version_code: number; // 客户端安装包 version code
	wgt_version: string; // 客户端当前 wgt 资源版本 name（如 1.0.3）；uni-app x 不检查 wgt，传空串
}

/** 弹窗展示的更新信息（经 flow 判定后的待更新内容） */
export interface UpgradeInfo {
	force: boolean; // 是否强制更新
	isWgt: boolean; // 是否 wgt 热更包
	version: string;
	version_code: number;
	download_url: string;
	install_url: string; // iOS itms-services 安装地址（仅全量包）
}
