import { Dimensions } from "react-native";
import type {
  LogWithError,
  LogWithEnv,
  ListenCallbackFn,
  IErrorEvent,
} from "../types";

// 同步存储接口（可注入 mmkv 等同步存储；默认使用内存缓存）
interface SyncStorage {
  getItem(key: string): string | null;
  setItem(key: string, value: string): void;
  removeItem(key: string): void;
}

interface LoggerProps {
  storage?: SyncStorage;
}

// 内存版同步存储：RN 无原生同步存储，默认仅在当前会话内缓存
class MemoryStorage implements SyncStorage {
  private map: Record<string, string> = {};
  getItem(key: string): string | null {
    return Object.prototype.hasOwnProperty.call(this.map, key)
      ? this.map[key]
      : null;
  }
  setItem(key: string, value: string): void {
    this.map[key] = value;
  }
  removeItem(key: string): void {
    delete this.map[key];
  }
}

// React Native 平台错误采集实现：
// 上报用 fetch（GET beacon 同款语义），错误监听通过 ErrorUtils.setGlobalHandler。
class ReactNative implements IErrorEvent {
  props: LoggerProps;
  private storage: SyncStorage;
  private prevHandler: any = null;
  private _hooked = false;

  constructor(props?: LoggerProps) {
    this.props = {
      ...props,
    };
    this.storage = props?.storage || new MemoryStorage();
  }

  public upload(url: string, data: any): Promise<void> {
    return new Promise((resolve) => {
      try {
        fetch(
          url + "?data=" + encodeURIComponent(JSON.stringify(data)),
        ).then(
          () => resolve(),
          () => resolve(),
        );
      } catch (error) {
        resolve();
      }
    });
  }

  public setCache(key: string, info: any): void {
    try {
      this.storage.setItem(key, JSON.stringify(info));
    } catch (error) {}
  }

  public getCache(key: string): any {
    try {
      let list = this.storage.getItem(key);
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
      this.storage.removeItem(key);
    } catch (error) {}
  }

  public getEnvInfo(): LogWithEnv {
    const env: LogWithEnv = {
      height: 0,
      width: 0,
    };
    try {
      const win = Dimensions.get("window");
      env.height = win.height || 0;
      env.width = win.width || 0;
    } catch (error) {}
    return env;
  }

  private callback(err: LogWithError): void {}

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

  private onError = (error: any, isFatal?: boolean): void => {
    if (error && String(error.message || error).includes("ResizeObserver")) {
      return;
    }
    const message =
      typeof error === "string" ? error : error?.message || String(error);
    const stack =
      typeof error === "string" ? "" : this.handleStack(error?.stack || "");
    this.callback({
      type: "error",
      event_type: isFatal ? "fatal" : "error",
      path: "",
      message,
      stack,
      ...this.getEnvInfo(),
    });
  };

  public listen(callback: ListenCallbackFn): void {
    this.callback = callback;
    if (this._hooked) return;
    const g = globalThis as any;
    const ErrorUtils = g.ErrorUtils;
    if (ErrorUtils && typeof ErrorUtils.setGlobalHandler === "function") {
      this._hooked = true;
      this.prevHandler =
        typeof ErrorUtils.getGlobalHandler === "function"
          ? ErrorUtils.getGlobalHandler()
          : null;
      ErrorUtils.setGlobalHandler((error: any, isFatal: any) => {
        this.onError(error, isFatal);
        if (typeof this.prevHandler === "function") {
          this.prevHandler(error, isFatal);
        }
      });
    }
  }

  public unListen(): void {
    this.callback = () => {};
    const g = globalThis as any;
    const ErrorUtils = g.ErrorUtils;
    if (this._hooked && ErrorUtils && this.prevHandler) {
      ErrorUtils.setGlobalHandler(this.prevHandler);
    }
    this._hooked = false;
    this.prevHandler = null;
  }
}

export default ReactNative;
