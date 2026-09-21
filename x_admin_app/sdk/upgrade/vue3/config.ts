/**
 * 升级 SDK 全局配置与共享状态（vue3 端）
 * 状态为模块级单例：check() 在页面外发起，弹窗组件放置于首页等常驻页面渲染
 */
import { reactive } from "vue";
import type { UpgradeInfo } from "./types";
import env from "@/utils/env";

/** 弹窗交互阶段（状态机，流转见 flow.ts） */
export enum Phase {
  HIDDEN = "hidden", // 未展示
  CONFIRM = "confirm", // 提示升级，等待用户选择
  DOWNLOADING = "downloading", // 下载中（强制模式不可关闭）
  READY = "ready", // 下载完成，等待安装
}

export interface UpgradeConfig {
  baseUrl: string; // 服务端地址
  checkPath: string; // 检查更新接口路径（公开接口，无需 Token）
  enableWgt: boolean; // 是否启用 wgt 热更（uni-app x 不支持，恒为 false）
}

export const config: UpgradeConfig = {
  baseUrl: env.baseUrl,
  checkPath: "/api/web/fabu/version/checkupdate",
  enableWgt: true,
};

/** 共享状态：驱动弹窗组件渲染 */
export const state = reactive({
  info: null as UpgradeInfo | null,
  phase: Phase.HIDDEN as string,
  progress: 0, // 下载进度 0-100
  hint: "", // 阶段提示文案
});

/** 发起检查（由 App.vue onLaunch 调用；弹窗组件由常驻页面渲染） */
let checkHandler: (() => void) | null = null;
// onLaunch 早于组件 setup 时暂存调用，组件注册后自动补执行
let pendingCheck = false;

export function registerCheckHandler(handler: () => void) {
  checkHandler = handler;
  if (pendingCheck) {
    pendingCheck = false;
    handler();
  }
}

export function check(options?: Partial<UpgradeConfig>) {
  if (options) Object.assign(config, options);
  if (checkHandler == null) {
    pendingCheck = true; // 等待 <x-upgrade /> 所在页面加载后执行
    return;
  }
  checkHandler!();
}
