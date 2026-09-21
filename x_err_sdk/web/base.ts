import { v7 as uuid_v7 } from "uuid";
import type { LogWithError, IErrorEvent, LogWithEnv } from "../types";

type Uid = string | number;

type Props = {
  Dns: string; // 域名
  Pid: string; // 项目标识
  Uid?: Uid; //用户标识
};

class Base {
  private Dns: string = "";
  private client_id: string = "";
  private Pid: string = "";
  private Uid: Uid = "";

  private platform: IErrorEvent | null = null;

  private MessageList: any[] = [];
  private timer: number = 0;
  private lastTimeStamp: number = Date.now(); //上次上报时间戳
  private diffTime: number = 1000 * 10; // 上报间隔

  constructor(props: Props, platform: IErrorEvent) {
    if (!props) {
      console.error("props is null");
      return;
    }
    if (!platform) {
      console.error("platform is null");
      return;
    }
    this.platform = platform;
    if (props.Dns && props.Pid) {
      this.Dns = props.Dns;
      this.Pid = props.Pid;
    } else {
      console.error("props.Dns and props.Pid cannot be null");
      return;
    }
    if (props.Uid) {
      this.Uid = String(props.Uid);
    }
    this.setClientID();
    this.SetUid();
    this.getLocalMessage();

    // 监听错误
    platform.listen((params: LogWithError) => {
      this.Push(params);
    });
    // 定时检查发送一次
    this.timer = setInterval(() => {
      if (Date.now() - this.lastTimeStamp >= this.diffTime) {
        this.upload();
      }
    }, 2000);
  }

  // 设置用户id
  public SetUid(uid?: Uid) {
    if (uid) {
      this.Uid = String(uid);
      this.platform?.setCache("x_err_uid", this.Uid);
    } else {
      const u_id = this.platform?.getCache("x_err_uid");
      if (u_id) {
        this.Uid = u_id;
      }
    }
    this.uploadInfo();
  }

  // 设置错误ID
  private setClientID() {
    const client_id = this.platform?.getCache("x_err_client_id");
    if (client_id) {
      this.client_id = client_id;
    } else {
      this.client_id = uuid_v7();
      this.platform?.setCache("x_err_client_id", this.client_id);
    }
  }
  // 获取本地缓存消息
  private getLocalMessage() {
    let x_err_message_list = this.platform?.getCache("x_err_message_list");
    if (x_err_message_list) {
      this.MessageList = x_err_message_list;
    } else {
      this.MessageList = [];
    }
  }
  // 记录错误，超出5条时立即上报，否则缓存到本地(等待定时器上报)
  public Push = (data: LogWithError) => {
    this.MessageList.push({
      ...data,
      project_key: this.Pid,
      client_id: this.client_id,
      user_id: this.Uid,
    });

    if (this.MessageList.length > 5) {
      this.upload(); //上传
    } else {
      this.platform?.setCache("x_err_message_list", this.MessageList);
    }
  };
  public uploadInfo = () => {
    if (!this.Dns) return; //未设置Dns服务器不上传

    try {
      this.platform
        ?.upload(this.Dns + `/admin/monitor_client/add`, {
          project_key: this.Pid,
          client_id: this.client_id,
          user_id: this.Uid,
          // width: envInfo.ScreenWidth,
          // height: envInfo.ScreenHeight,
        })
        .catch((err: any) => {
          // 上传失败
        });
    } catch (error) {}
  };

  // 上传文件
  public upload() {
    if (!this.Dns) return; //未设置Dns服务器不上传
    if (!this.MessageList.length) {
      return;
    }
    this.lastTimeStamp = Date.now();
    try {
      this.platform
        ?.upload(this.Dns + `/admin/monitor_error/add`, this.MessageList)
        .catch((err: any) => {
          // 上传失败
        })
        .then(() => {
          this.MessageList = [];
          this.platform?.delCache("x_err_message_list");
        });
    } catch (error) {}
  }
  public unListen() {
    clearInterval(this.timer);
    this.platform?.unListen();
  }
}
export default Base;
