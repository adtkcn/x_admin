package online

import "x_admin/app/service/iot_service"

// 在线/离线子包只认运行态自身，不反向依赖告警等上层。上层（pipeline）在 init 里
// 注册下列钩子，把“离线收尾活动告警”“删设备清告警运行态”交回上层执行，从而打破
// online ↔ alert 的循环依赖。

var (
	offlineHooks     []func(deviceId string)
	clearDeviceHooks []func(deviceId string)
)

// OnDeviceOffline 注册“设备转为离线”钩子（每台设备每次离线转换调用一次）
func OnDeviceOffline(fn func(deviceId string)) { offlineHooks = append(offlineHooks, fn) }

// OnDeviceClear 注册“设备运行态被清除”钩子（删设备时调用）
func OnDeviceClear(fn func(deviceId string)) { clearDeviceHooks = append(clearDeviceHooks, fn) }

// fireOfflineHooks 触发所有离线钩子
func fireOfflineHooks(deviceId string) {
	for _, fn := range offlineHooks {
		fn(deviceId)
	}
}

// markOfflineDB 库侧置离线，返回是否发生状态转换
func markOfflineDB(deviceId string) (bool, error) {
	return iot_service.DeviceService.MarkStatus(deviceId, 0)
}
