<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form">
            <uv-form-item label="项目key" prop="projectKey" borderBottom>
                    {{form.projectKey}}
            </uv-form-item>
            <uv-form-item label="sdk生成的客户端id" prop="clientId" borderBottom>
                    {{form.clientId}}
            </uv-form-item>
            <uv-form-item label="用户id" prop="userId" borderBottom>
                    {{form.userId}}
            </uv-form-item>
            <uv-form-item label="系统" prop="os" borderBottom>
                    {{form.os}}
            </uv-form-item>
            <uv-form-item label="浏览器" prop="browser" borderBottom>
                    {{form.browser}}
            </uv-form-item>
            <uv-form-item label="城市" prop="city" borderBottom>
                    {{form.city}}
            </uv-form-item>
            <uv-form-item label="屏幕" prop="width" borderBottom>
                    {{form.width}}
            </uv-form-item>
            <uv-form-item label="屏幕高度" prop="height" borderBottom>
                    {{form.height}}
            </uv-form-item>
            <uv-form-item label="ua记录" prop="ua" borderBottom>
                    {{form.ua}}
            </uv-form-item>
            <uv-form-item label="创建时间" prop="createTime" borderBottom>
                    {{form.createTime}}
            </uv-form-item>
            <uv-form-item label="更新时间" prop="clientTime" borderBottom>
                    {{form.clientTime}}
            </uv-form-item>
		</uv-form>
        <uv-button
            v-if="$perms('admin:monitor_client:edit')"
            type="primary"
            text="编辑"
            customStyle="margin: 40rpx 0"
            @click="edit"
        ></uv-button>

	</view>
</template>

<script setup lang="ts">
	import {ref} from "vue";
	import { onLoad,onShow } from "@dcloudio/uni-app";
	import { useDictData } from "@/hooks/useDictOptions";
	import { monitor_client_detail } from "@/api/monitor_client";


	import {
		toast,
		alert,
		toPath
	} from "@/utils/utils";

	let form = ref({
		id: "",
		projectKey: "",
		clientId: "",
		userId: "",
		os: "",
		browser: "",
		city: "",
		width: "",
		height: "",
		ua: "",
		createTime: "",
		clientTime: "",
	});
	onLoad((e) => {
		console.log("onLoad", e);
		getDetails(e.id);
	});
	onShow((e) => {
		if (form.value?.id) {
			getDetails(form.value.id);
		}
	});
	onPullDownRefresh(() => {
		getDetails(form.value.id);
	});
	function getDetails(id: number | string) {
		monitor_client_detail(id).then((res) => {
			uni.stopPullDownRefresh();
            if (res.code == 200) {
                if (res?.data) {
                    form.value = res?.data
                }
            } else {
                toast(res.message);
            }
        })
        .catch((err) => {
			uni.stopPullDownRefresh();
            toast(err.message||"网络错误");
        });
	}

	function edit() {
		toPath("/pages/monitor_client/edit", { id: form.value.id });
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>