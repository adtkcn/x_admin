export default class WebSocket {
	// 
	config = {
		url: "",
		reConnectTime: 3000, // 断线重连检测时间间隔
		heartTime: 1 * 20 * 1000, // 心跳时间间隔

		success: function() {},
		fail: function() {},

		onOpen: function() {},
		onClose: function() {},
		onMessage: function(data) {

		}
	}
	task = null
	isConnect = false //是否连接中
	reConnect = false; //是否在重连中，就是有定时器在尝试重连
	connectCount = 0; //链接次数

	IsNetWork = true //是否有网络
	constructor(config) {
		this.config = {
			...this.config,
			...config
		}
		this.initSocket()
		// this._startHeart()
		uni.onNetworkStatusChange((res) => {
			if (res.isConnected) {
				this.IsNetWork = true
			} else {
				this.IsNetWork = false
			}
		})
	}
	initSocket() {

		let task = uni.connectSocket({
			url: this.config.url,
			timeout: 1000 * 60 * 100, //100分钟
			success(res) {
				// this.config.success()
				console.log('websocket接口调用成功', res);
			},
			fail(err) {
				console.error('websocket接口调用失败', err);
				// this._reConnectSocket()
			}
		});
		task.onOpen(() => {
			this.connectCount++
			this.isConnect = true
			this.config.onOpen()
		})
		task.onMessage((message) => {
			this.config.onMessage(message)
		})
		task.onClose((res) => {
			console.error('websocket onClose', res);
			this.isConnect = false
			this._clearHeart()
			this.config.onClose(res)
		})
		task.onError((err) => {
			console.error('websocket onError', err);
			this.isConnect = false

			this.config.fail(err)

			// this._reConnectSocket()
		})
		this.task = task
		return task
	}
	send(val) {
		this.task && this.task.send({
			data: val,
			success() {
				console.log('发送消息成功');
			},
			fail() {
				console.error('发送消息失败');
			}
		})
	}
	close() {
		this._clearHeart()
		this._clearReConnect()
		this.task && this.task.close()
	}

	/**
	 * 心跳定时器
	 */
	heartTimer = null
	/**
	 * 开始心跳
	 * */
	_startHeart() {
		this.heartTimer = setInterval(() => {
			this.send('ping')
		}, this.config.heartTime)
	}
	// 清理心跳
	_clearHeart() {
		clearInterval(this.heartTimer)
	}
	/**
	 * 重连
	 */
	_reConnectSocket() {

		if (this.reConnect == false) {
			this.reConnect = true
			this.initSocket()
		}
	}
	/**
	 * 重连计时器
	 * */
	reConnectTimer = null
	_startReConnect() {
		this.reConnectTimer = setInterval(() => {
			if (this.isConnect == true) {
				this._clearReConnect()
				return
			}
			this._reConnectSocket()
		}, this.config.reConnectTime)
	}
	// 清理重连计时器
	_clearReConnect() {
		clearInterval(this.reConnectTimer)
	}
}