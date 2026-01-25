export type LogWithEnv = {
  Height?: number;
  Width?: number;
};
export type LogWithError =LogWithEnv& {
  Type: "error" | "event" | "resources" | "click";
  EventType: string;
  Path: string;
  Message?: string;
  Stack?: string;  
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
