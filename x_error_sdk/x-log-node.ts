// import http from "http";

// export function upload(url: string, data: Object) {
// 	let h = new Image();
// 	h.onerror = function () {};
// 	h.src = url + "?log=" + encodeURIComponent(data.toString());

// 	const req = http.request(url, { method: "POST" }, (res: any) => {
// 		res.setEncoding("utf8");
// 		res.on("data", (chunk) => {});
// 		// 监听接口请求完成, 除非error，最终都会执行end;
// 		res.on("end", () => {});
// 	});
// 	req.on("error", (e) => {});
// 	// 写入请求的数据
// 	data && req.write(data);
// 	req.end();
// }
