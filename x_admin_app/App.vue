<script>
	import {
		useUserStore
	} from "@/stores/user";

	export default {
		onLaunch: function() {
			const userStore = useUserStore();
			userStore
				.getInfo()
				.then((res) => {
					console.log("userInfo", res);
				})
				.catch((err) => {
					uni.redirectTo({
						url: "/pages/login/login",
					});
				});
		},
		onShow: function() {},
		onHide: function() {},
	};
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