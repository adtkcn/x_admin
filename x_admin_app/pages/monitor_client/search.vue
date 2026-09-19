<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form">
			<uv-form-item label="项目key" prop="project_key" borderBottom>
					<uv-input v-model="form.project_key"> </uv-input>
			</uv-form-item>
			<uv-form-item label="sdk生成的客户端id" prop="client_id" borderBottom>
					<uv-input v-model="form.client_id"> </uv-input>
			</uv-form-item>
			<uv-form-item label="用户id" prop="user_id" borderBottom>
					<uv-input v-model="form.user_id"> </uv-input>
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
			<uv-form-item label="ua记录" prop="ua" borderBottom>
					<uv-input v-model="form.ua"> </uv-input>
			</uv-form-item>
			<uv-form-item label="创建时间" prop="create_time" borderBottom>
					<x-date-range v-model:startTime="form.create_time_start"
							v-model:endTime="form.create_time_end"></x-date-range>
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
	import type {type_monitor_client_query} from "@/api/monitor_client";


	let form = ref<type_monitor_client_query>({
    project_key: '',
    client_id: '',
    user_id: '',
    os: '',
    browser: '',
    city: '',
    ua: '',
    create_time_start: '',
    create_time_end: '',
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