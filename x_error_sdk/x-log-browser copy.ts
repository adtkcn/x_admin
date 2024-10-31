import type { LogWithError, LogWithEnv, ListenCallbackFn } from "./types";
export function upload(url: string, data: object) {
  return new Promise(() => {
    try {
      let h = new Image();
      h.onerror = function (event) {
        var e = event as Event;
        e.preventDefault();
      };
      h.onload = function () {};
      h.src = url + "?data=" + encodeURIComponent(JSON.stringify(data));
    } catch (error) {}
  });
}
export function setCache(key: string, info: any) {
  localStorage.setItem(key, JSON.stringify(info));
}
export function getCache(key: string): any {
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
export function delCache(key: string) {
  localStorage.removeItem(key);
}

// 获取环境信息
export function getEnvInfo(): LogWithEnv {
  const env: LogWithEnv = {
    Type: "env",
    ScreenHeight: 0,
    ScreenWidth: 0,
  };

  // navigator对象数据
  if (window ) {
    env.ScreenHeight = window.innerHeight || 0; // 获取显示屏信息
    env.ScreenWidth = window.innerWidth || 0;
  }

  return env;
}
function errorCallback(err:Event) {
  
}
export function listenError(callback: ListenCallbackFn) {
  window.addEventListener(
    "error",
    (err: any) => {
      console.error(err);
      let target = err.target;
      if (target?.localName) {
        if (target?.localName == "img" || target?.localName == "script") {
          // errorResources(target?.localName, target.src);
          callback({
            Type: "resources",
            EventType: target?.localName,
            Path: target.src,
            Message:'',
            Stack: "",
          });
        } else if (target?.localName == "link") {
          // errorResources(target?.localName, target.href);
          callback({
            Type: "resources",
            EventType: target?.localName,
            Path: target.href,
            Message:'',
            Stack: "",
          });
        }
      } else {
        // errorLog(err.type, err.message, err.error?.stack || "");
        callback({
          Type: "error",
          EventType: err.type,
          Path:window.location.href,
          Message: err.message,
          Stack: err.error?.stack || "",
        });
      }
    },
    true
  );
  window.addEventListener("unhandledrejection", (err) => {
    if (typeof err.reason == "string") {
      // errorLog(err.type, err.reason, "");
      callback({
        Type: "error",
        EventType: err.type,
        Path:window.location.href,
        Message: err.reason,
        Stack: "",
      });
    } else if (typeof err.reason == "object") {
      // errorLog(err.type, err.reason?.message || "", err.reason?.stack || "");
      callback({
        Type: "error",
        EventType: err.type,
        Path:window.location.href,
        Message: err.reason?.message || "",
        Stack: err.reason?.stack || "",
      });
    }
  });
}

// export function listenBehavior(eventLog: EventLog) {
//   window.addEventListener("click", (e) => {
//     let target = e.target as HTMLElement;
//     let tagName = target?.localName;
//     let name = [tagName];
//     if (target.id) {
//       name.push("#" + target.id);
//     }
//     target.classList.forEach((item) => {
//       name.push("." + item);
//     });
//     if (target.innerText) {
//       name.push(":" + target.innerText);
//     }
//     eventLog("click", name.join(""));
//   });
//   window.addEventListener("hashchange", (e) => {
//     eventLog("pageChange", e.newURL);
//   });

//   eventLog("pageChange", window.location.href);
// }

export function listen(callback: ListenCallbackFn) {
  listenError((params: LogWithError) => {
  
      if (params.Stack) {
        let newStack: string[] = [];
        params.Stack.split("\n").map((item, index) => {
          if (index < 4) {
            newStack.push(item);
          }
        });
        params.Stack = newStack.join("\n");
      }
  
    callback(params);
  });
}
// 取消监听
export function unListen() {
  window.removeEventListener("error", () => {});
  window.removeEventListener("unhandledrejection", () => {});
}


export default {
  getEnvInfo,
  listenError,
  upload,
  getCache,
  setCache,
  delCache,
  listen,

};
