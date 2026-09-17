package fabu_service

import (
	"bytes"
	"text/template"

	"x_admin/app/model/fabu_model"
	"x_admin/core/response"
)

var PlistService = NewFabuPlistService()

func NewFabuPlistService() *fabuPlistService {
	return &fabuPlistService{}
}

type fabuPlistService struct{}

type plistData struct {
	AppName     string
	BundleId    string
	Version     string
	DownloadUrl string
}

// Build 渲染 iOS 安装 manifest plist
func (s fabuPlistService) Build(app fabu_model.FabuApp, version fabu_model.FabuAppVersion) (string, error) {
	tpl := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>items</key>
	<array>
		<dict>
			<key>assets</key>
			<array>
				<dict>
					<key>kind</key>
					<string>software-package</string>
					<key>url</key>
					<string>{{.DownloadUrl}}</string>
				</dict>
			</array>
			<key>metadata</key>
			<dict>
				<key>bundle-identifier</key>
				<string>{{.BundleId}}</string>
				<key>bundle-version</key>
				<string>{{.Version}}</string>
				<key>kind</key>
				<string>software</string>
				<key>title</key>
				<string>{{.AppName}}</string>
			</dict>
		</dict>
	</array>
</dict>
</plist>`
	t, err := template.New("plist").Parse(tpl)
	if err != nil {
		return "", response.CheckErr(err, "plist模板解析失败")
	}
	var buf bytes.Buffer
	data := plistData{AppName: app.Name, BundleId: app.BundleId, Version: version.Version, DownloadUrl: version.DownloadUrl}
	if err := t.Execute(&buf, data); err != nil {
		return "", response.CheckErr(err, "plist渲染失败")
	}
	return buf.String(), nil
}
