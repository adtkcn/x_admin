<template>
	<view v-if="state.info != null && state.phase !== Phase.HIDDEN" class="x-upgrade">
		<!-- 遮罩：普通模式可点击跳过，强制模式/下载中无效 -->
		<view class="x-upgrade__mask" @tap="onMaskTap"></view>
		<view class="x-upgrade__card">
			<view class="x-upgrade__title">
				发现新版本
				<text class="x-upgrade__ver">V{{ state.info.version }}</text>
				<text v-if="state.info.force" class="x-upgrade__force">必须更新</text>
			</view>

			<!-- 普通模式：等待用户确认 -->
			<view v-if="state.phase === Phase.CONFIRM" class="x-upgrade__body">
				<text class="x-upgrade__tip">{{ state.hint !== '' ? state.hint : '建议升级到最新版本以获得更好体验' }}</text>
			</view>
			<!-- 下载中：进度条 -->
			<view v-else-if="state.phase === Phase.DOWNLOADING" class="x-upgrade__body">
				<view class="x-upgrade__bar">
					<view class="x-upgrade__bar-inner" :style="{ width: state.progress + '%' }"></view>
				</view>
				<text class="x-upgrade__tip">下载中 {{ state.progress }}%</text>
			</view>
			<!-- 下载完成：等待/重试安装 -->
			<view v-else-if="state.phase === Phase.READY" class="x-upgrade__body">
				<text class="x-upgrade__tip">下载完成，请在安装界面完成安装</text>
			</view>

			<view class="x-upgrade__footer">
				<!-- 确认阶段：普通模式可跳过 -->
				<template v-if="state.phase === Phase.CONFIRM">
					<button v-if="!state.info.force" class="x-upgrade__btn x-upgrade__btn--plain" size="mini"
						@tap="refuse()">
						以后再说
					</button>
					<button class="x-upgrade__btn" size="mini" @tap="accept()">{{ state.hint !== '' ? '重试升级' : '立即升级' }}</button>
				</template>
				<!-- 下载中：普通模式可取消 -->
				<template v-else-if="state.phase === Phase.DOWNLOADING">
					<button v-if="!state.info.force" class="x-upgrade__btn x-upgrade__btn--plain" size="mini"
						@tap="refuse()">
						取消下载
					</button>
				</template>
				<!-- 就绪：重新拉起安装 -->
				<template v-else-if="state.phase === Phase.READY">
					<button v-if="!state.info.force" class="x-upgrade__btn x-upgrade__btn--plain" size="mini"
						@tap="refuse()">
						稍后再说
					</button>
					<button class="x-upgrade__btn" size="mini" @tap="accept()">立即安装</button>
				</template>
			</view>
		</view>
	</view>
</template>

<script setup lang="ts">
	/**
	 * x-upgrade 应用升级弹窗组件（uni-app vue3 端）
	 * easycom 自动注册；放置于常驻页面（如首页），检查入口在 App.vue onLaunch 调用 check()
	 * 交互规则：
	 *  - 普通模式：提示升级，同意后台下载，可"以后再说"（同版本不再提示）
	 *  - 强制模式：必须等待下载安装，无关闭入口（遮罩/按钮均不可关闭）
	 *  - wgt 热更：不经过本组件，flow.ts 静默下载安装后仅提示重启
	 *  - iOS：无 apk 下载，确认后打开 App Store / itms-services OTA
	 */
	import { state, Phase } from "@/sdk/upgrade/vue3/config";
	import { accept, refuse, onMaskTap, setupCheck } from "@/sdk/upgrade/vue3/flow";

	// 注册全局检查入口，供 App.vue 的 check() 调用
	setupCheck();
</script>

<style scoped>
	.x-upgrade {
		position: fixed;
		left: 0;
		top: 0;
		right: 0;
		bottom: 0;
		z-index: 9999;
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.x-upgrade__mask {
		position: absolute;
		left: 0;
		top: 0;
		right: 0;
		bottom: 0;
		background-color: rgba(0, 0, 0, 0.5);
	}

	.x-upgrade__card {
		position: relative;
		width: 620rpx;
		background-color: #ffffff;
		border-radius: 24rpx;
		padding: 40rpx 36rpx 30rpx;
	}

	.x-upgrade__title {
		font-size: 34rpx;
		font-weight: bold;
		color: #303133;
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	.x-upgrade__ver {
		margin-left: 12rpx;
		font-size: 26rpx;
		font-weight: normal;
		color: #576b95;
	}

	.x-upgrade__force {
		margin-left: 16rpx;
		font-size: 22rpx;
		color: #ffffff;
		background-color: #fa5151;
		border-radius: 8rpx;
		padding: 2rpx 12rpx;
	}

	.x-upgrade__body {
		margin-top: 28rpx;
	}

	.x-upgrade__tip {
		font-size: 26rpx;
		color: #606266;
	}

	.x-upgrade__bar {
		height: 16rpx;
		background-color: #f0f1f3;
		border-radius: 8rpx;
		overflow: hidden;
	}

	.x-upgrade__bar-inner {
		height: 100%;
		background-color: #007aff;
		border-radius: 8rpx;
	}

	.x-upgrade__footer {
		margin-top: 36rpx;
		display: flex;
		flex-direction: row;
		justify-content: flex-end;
		align-items: center;
	}

	.x-upgrade__btn {
		margin-left: 20rpx;
		background-color: #007aff;
		color: #ffffff;
		font-size: 28rpx;
		border-radius: 12rpx;
		line-height: 2.4;
	}

	.x-upgrade__btn--plain {
		background-color: #ffffff;
		color: #909399;
		border: 1rpx solid #dcdfe6;
	}
</style>
