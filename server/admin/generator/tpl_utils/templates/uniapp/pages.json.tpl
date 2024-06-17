// 请将pages里的数据手动合并到根目录下pages.json中
{
	"pages": [
        {
			"path": "pages/{{{ .ModuleName }}}/index",
			"style": {
				"navigationBarTitleText": "{{{.FunctionName}}}",
				"enablePullDownRefresh": true,
				"onReachBottomDistance": 100
			}
		},
		{
			"path": "pages/{{{ .ModuleName }}}/details",
			"style": {
				"navigationBarTitleText": "{{{.FunctionName}}}详情",
				"enablePullDownRefresh": true
			}
		},
		{
			"path": "pages/{{{ .ModuleName }}}/edit",
			"style": {
				"navigationBarTitleText": "编辑{{{.FunctionName}}}"
			}
		},
		{
			"path": "pages/{{{ .ModuleName }}}/search",
			"style": {
				"navigationBarTitleText": "搜索{{{.FunctionName}}}"
			}
		}
    ]
}
