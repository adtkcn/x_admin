import env from "@/utils/env";
import {
	getLocalStorage
} from "@/utils/storage.js";
let requestNumber = 0; // 请求计数
function openLoading() {
	if (requestNumber <= 0) {
		uni.showLoading({
			title: "请求中",
			mask: true,
		});
	}
	requestNumber++;
}

function closeLoading() {
	// requestNumber--;
	// if (requestNumber == 0) {
	//   uni.hideLoading();
	// } else {
	// }
}

export function request(options = {}) {
	options.url = `${env.baseUrl}${env.urlPrefix}${options.url}`;
	var token = getLocalStorage("token");
	var opts = {
		timeout: 10000,
		header: {
			"content-type": "application/json",
			Token: token,
		},
	};

	// uni.showLoading({
	// 	title: "请求中",
	// 	mask: true,
	// });
	return new Promise((resolve, reject) => {
		// 发送 HTTP 请求
		uni.request({
			...opts,
			...options,
			success: function(res) {
				// requestNumber--;
				// uni.hideLoading()
				console.log(options.url, {
					request: options,
					response: res.data
				});
				if (res.statusCode == 401) {
					// token过期
					// reject(new Error('登录失效，请重新登录'));
					uni.redirectTo({
						url: "/pages/login/login",
					});
				} else if (res.statusCode !== 200) {
					reject(new Error(res.data.message));
				} else {
					resolve(res.data);
				}

			},
			fail: function(err) {
				console.log(options.url, {
					request: options,
					response: err
				});
				// requestNumber--;
				// uni.hideLoading()
				reject(new Error("网络繁忙,请稍后再试"));
			},
		});
	});
}
export function upload(options = {}) {
	options.url = `${env.baseUrl}${options.url}`;

	var token = getLocalStorage("token");

	var header = {
		Token: token,
	};
	return new Promise((resolve, reject) => {
		// 发送 HTTP 请求
		uni.uploadFile({
			name: "file",
			...options,
			header: {
				...header,
			},
			success: function(res) {
				// requestNumber--;
				// uni.hideLoading()
				console.log("request", res);

				if (res.statusCode !== 200) {
					var data = JSON.parse(res.data);
					reject(data.message);
				} else {
					var data = JSON.parse(res.data);
					resolve(data);
				}
			},
			fail: function(err) {
				// requestNumber--;
				// uni.hideLoading()
				reject("网络繁忙,请稍后再试");
			},
		});
	});
}

export function download(options = {}) {
	options.url = `${env.baseUrl}${options.url}`;

	var token = getLocalStorage("token");

	var header = {
		Token: token,
	};
	return new Promise((resolve, reject) => {
		// 发送 HTTP 请求
		uni.downloadFile({
			...options,
			//   filePath:"",
			header: {
				...header,
			},
			success: function(res) {
				// console.log("Content-Disposition", res.header["Content-Disposition"]);
				// var fileName = decodeURIComponent(
				// 	res.header["Content-Disposition"].split("=")[1]
				// );
				if (res.statusCode !== 200) {
					// var data = JSON.parse(res.data);
					reject("失败");
				} else {}
			},
			fail: function(err) {
				// requestNumber--;
				// uni.hideLoading()
				reject("网络繁忙,请稍后再试");
			},
		});
	});
}


export default {
	request,
	upload,
	download
}