
// let baseUrl = "https://captcha.anji-plus.com/captcha-api"
import env from "@/utils/env";
let baseUrl =env.baseUrl+'/api/common';

export const myRequest = (option={})=>{
	return new Promise((reslove,reject)=>{
		uni.request({
			url: baseUrl + option.url, 
			data :option.data,
			method:option.method || "GET",
			success: (result) => {
				reslove(result)
			},
			fail:(error)=>{
				reject(error)
			}
		})
	})
}
