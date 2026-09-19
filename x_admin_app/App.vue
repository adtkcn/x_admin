<script setup lang="ts">
	import {
		onLaunch,
		onShow,
		onHide
	} from "@dcloudio/uni-app";
	import {
		useUserStore
	} from "@/stores/user";
	import { check as checkAppUpdate } from "@/sdk/upgrade/vue3/config";

	onLaunch(() => {
		// 启动时检查应用更新（普通/强制/wgt 策略见 sdk/upgrade/README.md）
		checkAppUpdate();
		const userStore = useUserStore();
		userStore
			.getInfo()
			.then((res: any) => {
				console.log("userInfo", res);
			})
			.catch(() => {
				uni.redirectTo({
					url: "/pages/login/login",
				});
			});
	});

	onShow(() => {});
	onHide(() => {});
</script>

<style>
	#app {
		/* 解决app上uv-sticky浮动的问题 ，估计又修复了先注释掉*/
		/* height: auto; */
	}
	button:after {
		border: 0 !important;
	}
	/*每个页面公共css */
	.h100 {
		height: 100%;
	}

	.row {
		display: flex;
		flex-direction: row;
	}

	.flex_1 {
		flex: 1;
	}

	.center {
		justify-content: center;
		align-items: center;
	}

	.right {
		display: flex;
		flex-direction: row;
		justify-content: flex-end;
	}

	.search {
		padding: 5rpx;
		background-color: #fff;
	}
</style>