
export type LogWithEnv = {
  Type: "env";
  ScreenHeight?: number;
  ScreenWidth?: number;
};
export type LogWithError = {
  Type: "error"|"event"|"resources";
  EventType: string;
  Path:string;
  Message: string;
  Stack: string;
};


export interface Platform {
  upload(url: string, data: Object);
  getEnvInfo(): LogWithEnv;
  listen(callback: ListenCallbackFn): void;
  getCache(key: string): any;
  setCache(key: string, info: any): void;
  delCache(key: string): void;
}
 
export type ListenCallbackFn = (params: LogWithError) => void;
