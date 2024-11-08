import type {
  LogWithError,
  LogWithEnv,
  ListenCallbackFn,
  IErrorEvent,ISlow
} from "../types";

interface LoggerProps {
  // timeout:number
  onloadTimeOut?: number;
}
class Web implements IErrorEvent {
  props: LoggerProps;
  constructor(props?: LoggerProps) {
    this.props = {
      onloadTimeOut: 5000,
      ...props,
    };
  }
  public upload(url: string, data: any): Promise<void> {
    return new Promise((resolve) => {
      try {
        let h = new Image();
        h.onload = function (event) {
          var e = event as Event;
          e.preventDefault();
        };
        h.onerror = function () {};
        h.src = url + "?data=" + encodeURIComponent(JSON.stringify(data));
        resolve();
      } catch (error) {
        resolve();
      }
    });
  }

  public setCache(key: string, info: any): void {
    localStorage.setItem(key, JSON.stringify(info));
  }

  public getCache(key: string): any {
    try {
      let list = localStorage.getItem(key);
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
    localStorage.removeItem(key);
  }

  public getEnvInfo(): LogWithEnv {
    const env: LogWithEnv = {
      Type: "env",
      ScreenHeight: 0,
      ScreenWidth: 0,
    };
    if (window) {
      env.ScreenHeight = window.innerHeight || 0; // 获取显示屏信息
      env.ScreenWidth = window.innerWidth || 0;
    }
    return env;
  }
  private callback(err: LogWithError|ISlow): void {}
  private listenError = (err: any) => {
    console.error([err]);
    let target = err.target;
    if (target?.localName) {
      if (target?.localName === "img" || target?.localName === "script") {
        this.callback({
          Type: "resources",
          EventType: target?.localName,
          Path: target.src,
          Message: "",
          Stack: "",
        });
      } else if (target?.localName === "link") {
        this.callback({
          Type: "resources",
          EventType: target?.localName,
          Path: target.href,
        });
      }
    } else {
      this.callback({
        Type: "error",
        EventType: err.type,
        Path: window.location.href,
        Message: err.message,
        Stack: this.handleStack(err.error?.stack || ""),
      });
    }
  };
  private unhandledrejection = (err: any): void => {
    console.error(err);
    if (err && typeof err.reason === "string") {
      this.callback({
        Type: "error",
        EventType: err.type,
        Path: window.location.href,
        Message: err.reason,
        Stack: "",
      });
    } else if (err && typeof err.reason === "object") {
      this.callback({
        Type: "error",
        EventType: err.type,
        Path: window.location.href,
        Message: err.reason?.message || "",
        Stack: this.handleStack(err.reason?.stack || ""),
      });
    }
  };


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
  private onLoad = () => {
    // 获取性能数据
    const entries = performance.getEntriesByType("navigation");
    if (entries.length > 0) {
      const performanceData = entries[0] as PerformanceNavigationTiming;

      console.log("performanceData", performanceData);

      // 计算页面onload时间
      let onloadTime =
        performanceData.loadEventStart - performanceData.startTime;
      if (
        this.props.onloadTimeOut &&
        onloadTime > this.props.onloadTimeOut
      ) {
        // 页面加载时间5s以上
        this.callback({
          Type: "onloadTime",
          Path: window.location.href,
          Time:onloadTime
        });
      }
    }

  };
  public listen(callback: ListenCallbackFn): void {
    this.callback = callback;
    window.addEventListener("unhandledrejection", this.unhandledrejection);
    window.addEventListener("error", this.listenError, true);
    // window.addEventListener("click", this.listenClick);
    // window.addEventListener("hashchange", this.listenHashRouterChange);
    // window.addEventListener("popstate", this.listenHistoryRouterChange);

    window.addEventListener("load", this.onLoad);
  }

  public unListen(): void {
    this.callback = () => {};
    window.removeEventListener("error", this.listenError);
    window.removeEventListener("unhandledrejection", this.unhandledrejection);
    // window.removeEventListener("click", this.listenClick);
    // window.removeEventListener("hashchange", this.listenHashRouterChange);
    // window.removeEventListener("popstate", this.listenHistoryRouterChange);
  }
}

export default Web;
