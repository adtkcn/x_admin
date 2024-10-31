import { v1 as uuid_v1 } from "uuid";
import type { LogWithError, IErrorEvent,LogWithEnv  } from "./types";

type Uid = string | number;

type Props = {
  Dns: string; // 域名
  Pid: string; // 项目标识
  Uid?: Uid; //用户标识
};

class XLog {
  private Dns: string = "";
  private client_id: string = "";
  private Pid: string = "";
  private Uid: Uid = "";

  private platform: IErrorEvent;

  private MessageList: any[] = [];
  private timer: number;

  constructor(props: Props, platform: IErrorEvent) {
    this.platform = platform;
    if (props) {
      props.Dns && (this.Dns = props.Dns);
      props.Pid && (this.Pid = props.Pid);
      props.Uid && (this.Uid = props.Uid);
    }
    this.SetClientID();
    this.getLocalMessage();
    this.initEnv();

 
    // 监听错误
    platform.listen((params: LogWithError) => {
      console.log('listenCallback',params);
      
      this.Push(params);
    });

    this.timer=setInterval(() => {
      this.upload();
    }, 1000 * 10);
  }
  private initEnv() {
    let envInfo:LogWithEnv  = this.platform.getEnvInfo();
    if (envInfo) {
      this.uploadInfo(envInfo);
    }
  }
  // 关联用户id
  public SetUid(uid: Uid) {
    this.Uid = uid;
    this.initEnv();
  }
  // 设置错误ID
  private SetClientID() {
    const client_id = this.platform.getCache("x_log_client_id");
    if (client_id) {
      this.client_id = client_id;
    } else {
      this.client_id = uuid_v1();
      this.platform.setCache("x_log_client_id", this.client_id);
    }
  }
  // 获取本地缓存消息
  private getLocalMessage() {
    let x_message_list = this.platform.getCache("x_message_list");
    if (x_message_list) {
      this.MessageList = x_message_list;
    } else {
      this.MessageList = [];
    }
  }
  // 记录错误，超出5条时立即上报，否则缓存到本地(等待定时器上报)
  public Push(data: LogWithError) {
    console.log(this);

    this.MessageList.push({
      ...data,
      ProjectKey:this.Pid,
      ClientId:this.client_id
    });

    if (this.MessageList.length > 5) {
      this.upload(); //上传
    } else {
      this.platform.setCache(
        "x_message_list",
        this.MessageList
      );
    }
  }
  public uploadInfo(envInfo:LogWithEnv ) {
    if (!this.Dns) return; //未设置Dns服务器不上传
 
    try {
      this.platform
        .upload(
          this.Dns +
            `/admin/monitor_client/add`,
          {
            ProjectKey:this.Pid,
            ClientId:this.client_id,
            UserId:this.Uid,
            Width: envInfo.ScreenWidth,
            Height:envInfo.ScreenHeight
          }
        )
        .catch((err: any) => {
          // 上传失败
        });
    } catch (error) {}
  }

  // 上传文件
  public upload() {
    if (!this.Dns) return; //未设置Dns服务器不上传
    if (!this.MessageList.length) {
      return;
    }
    try {
      this.platform
        .upload(
          this.Dns +
            `/admin/monitor_error/add`,
          this.MessageList
        )
        .catch((err: any) => {
          // 上传失败
        });
      this.MessageList = [];
      this.platform.delCache("x_message_list");
    } catch (error) {}
  }
  public unListen(){
    clearInterval(this.timer)
    this.platform.unListen()
  }
  // public Vue2 = (app: any) => {
  //   app.prototype.$xLog = this;
  // };
  // public Vue3 = (app: any) => {
  //   app.config.globalProperties.$xLog = this;
  // };
}
export default XLog;
