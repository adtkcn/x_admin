export function upload(url: string, data: Object) {
	if (uni) {
		uni.request({
			url: url, //仅为示例，并非真实接口地址。
			data: {
				log: encodeURIComponent(JSON.stringify(data)),
			},
			method: "POST",
			header: {
				// xlog: "version", //自定义请求头信息
			},
			success: (res: any) => {
				console.log(res.data);
			},
			fail: (err: any) => {
				console.log(err);
			},
		});
	}
}
