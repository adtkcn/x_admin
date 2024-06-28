<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form">
			<uv-form-item label="项目key" prop="projectKey" borderBottom>
			</uv-form-item>
			<uv-form-item label="sdk生成的客户端id" prop="clientId" borderBottom>
					<uv-input v-model="form.clientId"> </uv-input>
			</uv-form-item>
			<uv-form-item label="用户id" prop="userId" borderBottom>
					<uv-input v-model="form.userId"> </uv-input>
			</uv-form-item>
			<uv-form-item label="系统" prop="os" borderBottom>
					<uv-input v-model="form.os"> </uv-input>
			</uv-form-item>
			<uv-form-item label="浏览器" prop="browser" borderBottom>
					<uv-input v-model="form.browser"> </uv-input>
			</uv-form-item>
			<uv-form-item label="城市" prop="city" borderBottom>
					<uv-input v-model="form.city"> </uv-input>
			</uv-form-item>
			<uv-form-item label="屏幕" prop="width" borderBottom>
					<uv-input v-model="form.width"> </uv-input>
			</uv-form-item>
			<uv-form-item label="屏幕高度" prop="height" borderBottom>
					<uv-input v-model="form.height"> </uv-input>
			</uv-form-item>
			<uv-form-item label="ua记录" prop="ua" borderBottom>
					<uv-input v-model="form.ua"> </uv-input>
			</uv-form-item>
			<uv-form-item label="创建时间" prop="createTime" borderBottom>
					<x-date-range v-model:startTime="form.createTimeStart"
							v-model:endTime="form.createTimeEnd"></x-date-range>
			</uv-form-item>
			<uv-form-item label="更新时间" prop="clientTime" borderBottom>
					<x-date-range v-model:startTime="form.clientTimeStart"
							v-model:endTime="form.clientTimeEnd"></x-date-range>
			</uv-form-item>

			<uv-button type="primary" text="搜索" customStyle="margin-top: 20rpx" @click="submit"></uv-button>
		</uv-form>
	</view>
</template>

<script setup  lang="ts">
 
	import {
		onLoad
	} from "@dcloudio/uni-app";
	import {
		reactive,
		ref,
		computed
	} from "vue";
	import {
		toPath,
		toast,
		clearObjEmpty
	} from "@/utils/utils";
	import {
		useDictData
	} from "@/hooks/useDictOptions";
	import xDateRange from "@/components/x-date-range/x-date-range.vue";
	import type {type_monitor_client_query} from "@/api/monitor_project";


	let form = ref<type_monitor_client_query>({
    projectKey: '',
    clientId: '',
    userId: '',
    os: '',
    browser: '',
    city: '',
    width: '',
    height: '',
    ua: '',
    createTimeStart: '',
    createTimeEnd: '',
    clientTimeStart: '',
    clientTimeEnd: '',
	});

	function submit() {
		console.log("submit", form.value);

		const search = clearObjEmpty(form.value);

		if (Object.keys(search).length === 0) {
			return toast("请输入查询条件");
		}

		toPath("/pages/monitor_client/index", search);
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>