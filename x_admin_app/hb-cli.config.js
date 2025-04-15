const {
	defineConfig
} = require("@adtkcn/hb-cli");

/**
 *
 * @param {object} param0
 * @param {string} param0.mode 打包模式,通过命令行参数--mode传入
 * @returns
 */
module.exports = ({
	mode
}) => {
	console.log("当前mode环境：", mode);

	return defineConfig({
		// 打包配置项，全部为uniapp官方配置(必须)
		packConfig() {
			return {
				//项目名字或项目绝对路径
				"project": "x_admin_app",
				//打包平台 默认值android  值有"android","ios" 如果要打多个逗号隔开打包平台
				"platform": "android",
				//是否使用自定义基座 默认值false  true自定义基座 false自定义证书
				"iscustom": false,
				//打包方式是否为安心打包默认值false,true安心打包,false传统打包
				"safemode": true,
				//android打包参数
				"android": {
					//安卓包名
					"packagename": "uni.UNIFB29F21",
					//安卓打包类型 默认值0 0 使用自有证书 1 使用公共证书 2 使用老版证书
					"androidpacktype": "3",
					//安卓使用自有证书自有打包证书参数
					//安卓打包证书别名,自有证书打包填写的参数
					"certalias": "",
					//安卓打包证书文件路径,自有证书打包填写的参数
					"certfile": "",
					//安卓打包证书密码,自有证书打包填写的参数
					"certpassword": "",
					//安卓平台要打的渠道包 取值有"google","yyb","360","huawei","xiaomi","oppo","vivo"，如果要打多个逗号隔开
					"channels": ""
				},
				//ios打包参数
				"ios": {
					//ios appid
					"bundle": "uni.UNIFB29F21",
					//ios打包支持的设备类型 默认值iPhone 值有"iPhone","iPad" 如果要打多个逗号隔开打包平台
					"supporteddevice": "iPhone,iPad",
					//iOS打包是否打越狱包,只有值为true时打越狱包,false打正式包
					"isprisonbreak": false,

					//iOS使用自定义证书打包的profile文件路径
					"profile": "",
					//iOS使用自定义证书打包的p12文件路径
					"certfile": "",
					//iOS使用自定义证书打包的证书密码
					"certpassword": ""
				},
				//是否混淆 true混淆 false关闭
				"isconfusion": false,
				//开屏广告 true打开 false关闭
				"splashads": false,
				//悬浮红包广告true打开 false关闭
				"rpads": false,
				//push广告 true打开 false关闭
				"pushads": false,
				//加入换量联盟 true加入 false不加入
				"exchange": false
			}
		},
		// 自定义manifest.json配置项: 合并到manifest.json中(可选)
		mergeManifestConfig() {
			return {
			};
		},
    /**
     * 创建APP内环境变量，生成HBuilderEnv.js文件（可选）
     * @returns {any} 环境变量
     */
    appConfig:{
      output:"./HBuilderEnv.js", // 默认./HBuilderEnv.js
      create(){
      },
    },
		/**
		 * 定义manifest.versionName的生成规则(可选，默认：auto-increment)
		 */
		version: {
			mode: "date", // 可选值："custom"、"date"、"auto-increment",
			/**
			 * 自定义版本号（可选）
			 * @param {string[]} VersionNameArr 版本号数组，如：[1, 0, 0]
			 * @returns {string} 版本号数组，如：1.0.0
			 */
			customVersion: (VersionNameArr) => {
				console.log(VersionNameArr);
				var lastIndex = VersionNameArr.length - 1;
				VersionNameArr[lastIndex] = String(
					parseInt(VersionNameArr[lastIndex]) + 1
				);
				return VersionNameArr.join(".");
			},
		},

    /**
     * 打包后回调(可选)
     * @param {string} filePath 文件路径,appResource打包后是目录
     * @param {"android"|"ios"| "appResource"| "wgt"} fileType 文件类型
     */
		async onPackEnd(filePath, fileType) {
			//上传回调
			console.log("上传回调开始",filePath, fileType);
		},
	});
};