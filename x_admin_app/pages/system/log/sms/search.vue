<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form">
			<uv-form-item label="场景编号" prop="Scene" borderBottom>
			</uv-form-item>
			<uv-form-item label="手机号码" prop="Mobile" borderBottom>
					<uv-input v-model="form.Mobile"> </uv-input>
			</uv-form-item>
			<uv-form-item label="发送内容" prop="Content" borderBottom>
			</uv-form-item>
			<uv-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="Status" borderBottom>
			</uv-form-item>
			<uv-form-item label="短信结果" prop="Results" borderBottom>
			</uv-form-item>
			<uv-form-item label="发送时间" prop="SendTime" borderBottom>
					<x-date-range v-model:startTime="form.SendTimeStart"
							v-model:endTime="form.SendTimeEnd"></x-date-range>
			</uv-form-item>
			<uv-form-item label="创建时间" prop="CreateTime" borderBottom>
					<x-date-range v-model:startTime="form.CreateTimeStart"
							v-model:endTime="form.CreateTimeEnd"></x-date-range>
			</uv-form-item>
			<uv-form-item label="更新时间" prop="UpdateTime" borderBottom>
					<x-date-range v-model:startTime="form.UpdateTimeStart"
							v-model:endTime="form.UpdateTimeEnd"></x-date-range>
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
		useDictData,useListAllData
	} from "@/hooks/useDictOptions";
	import xDateRange from "@/components/x-date-range/x-date-range.vue";
	import type {type_system_log_sms_query} from "@/api/system_log_sms";

	let form = ref<type_system_log_sms_query>({
    Scene: null,
    Mobile: null,
    Content: null,
    Status: null,
    Results: null,
    SendTimeStart: null,
    SendTimeEnd: null,
    CreateTimeStart: null,
    CreateTimeEnd: null,
    UpdateTimeStart: null,
    UpdateTimeEnd: null,
	});

	function submit() {
		console.log("submit", form.value);

		const search = clearObjEmpty(form.value);

		if (Object.keys(search).length === 0) {
			return toast("请输入查询条件");
		}

		toPath("/pages/system/log/sms/index", search);
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>