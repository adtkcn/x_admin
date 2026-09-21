export type LogWithEnv = {
  height?: number;
  width?: number;
};
export type LogWithError =LogWithEnv& {
  type: "error" | "event" | "resources" | "click";
  event_type: string;
  path: string;
  message?: string;
  stack?: string;  
};

//  扩展必须实现的接口
export interface IErrorEvent {
  upload(url: string, data: object): Promise<void>;
  setCache(key: string, info: any): void;
  getCache(key: string): any;
  delCache(key: string): void;
  getEnvInfo(): LogWithEnv;
  listen(callback: ListenCallbackFn): void;
  unListen(): void;
}
export type ListenCallbackFn = (params: LogWithError) => void;
