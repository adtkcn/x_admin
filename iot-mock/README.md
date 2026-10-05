# IoT 设备模拟器 (iot-mock)

向本项目设备侧 HTTP 接入层模拟上报，用于联调**区域大屏 / 告警 / 物模型校验**。纯标准库实现，无第三方依赖。

## 接入契约（对应本项目）

- `POST {base}/api/iot/telemetry`  body：`{"values":{"temperature":26.3},"ts":<毫秒>}`
- `POST {base}/api/iot/heartbeat`  body：`{"ts":<毫秒>}`
- 鉴权请求头：`X-Device-Sn` + `X-Access-Token`（设备凭证见 `x_iot_device.access_token`）
- 设备所属**产品的接入协议须为 `http`**（或留空），否则接入层拒绝 HTTP 上报
- 服务默认端口 `18000`；单次上报属性条数上限 `IOT.MaxValues`（默认 50）

## 编译

```powershell
go build -o iot-mock.exe .
```

## 运行

```powershell
# 手动模式：输入数值回车即推送到 -key 标识符
.\iot-mock.exe -sn DEV001 -token <access_token> -key temperature

# 自动模式：每 2s 上报一次，值在 26 上下 ±3 随机抖动，并附带心跳
.\iot-mock.exe -sn DEV001 -token <access_token> -key temperature -interval 2s -jitter 3 -heartbeat

# 多属性上报：初始键值覆盖 -key
.\iot-mock.exe -sn DEV001 -token <access_token> -values temperature=26.3,humidity=55 -interval 5s
```

### 参数

| 参数 | 默认值 | 说明 |
| --- | --- | --- |
| `-base` | `http://127.0.0.1:18000` | 服务基地址（env `IOT_BASE`） |
| `-sn` | 必填 | 设备编号 `X-Device-Sn`（env `IOT_SN`） |
| `-token` | 必填 | 设备接入凭证 `X-Access-Token`（env `IOT_TOKEN`） |
| `-key` | `value` | 输入数值时写入的属性标识符 |
| `-values` | 空 | 初始属性键值，`k=v` 逗号分隔 |
| `-interval` | `0` | 自动上报间隔，`0`=仅手动推送 |
| `-jitter` | `0` | 自动模式下对 `-key` 数值叠加的随机幅度（±绝对值） |
| `-heartbeat` | `false` | 每次自动上报前额外发送一次心跳 |
| `-timeout` | `5s` | 单次 HTTP 请求超时 |

### 交互命令（手动模式，stdin）

- `<数值>`：写入 `-key` 并立即推送遥测
- `hb`：发送一次心跳
- `q` / `quit` / `exit`：退出

## 物模型说明

上报的 `values` 键须是产品物模型中已定义的 identifier，否则会被处理层的物模型校验过滤/拒绝。先用 `-values` 或 stdin 输入对应标识符与数值即可模拟设备数据流。
