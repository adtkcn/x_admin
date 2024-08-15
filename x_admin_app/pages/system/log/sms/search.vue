<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form">
			<uv-form-item label="场景编号" prop="scene" borderBottom>
			</uv-form-item>
			<uv-form-item label="手机号码" prop="mobile" borderBottom>
					<uv-input v-model="form.mobile"> </uv-input>
			</uv-form-item>
			<uv-form-item label="发送内容" prop="content" borderBottom>
			</uv-form-item>
			<uv-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="status" borderBottom>
						<x-picker v-model="form.status" valueKey="value" labelKey="name" :columns="dictData.flow_apply_status"></x-picker>
			</uv-form-item>
			<uv-form-item label="短信结果" prop="results" borderBottom>
			</uv-form-item>
			<uv-form-item label="发送时间" prop="send_time" borderBottom>
					<x-date-range v-model:startTime="form.send_time_start"
							v-model:endTime="form.send_time_end"></x-date-range>
			</uv-form-item>
			<uv-form-item label="创建时间" prop="create_time" borderBottom>
					<x-date-range v-model:startTime="form.create_time_start"
							v-model:endTime="form.create_time_end"></x-date-range>
			</uv-form-item>
			<uv-form-item label="更新时间" prop="update_time" borderBottom>
					<x-date-range v-model:startTime="form.update_time_start"
							v-model:endTime="form.update_time_end"></x-date-range>
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
const { dictData } = useDictData<{
    flow_apply_status: any[]
}>(['flow_apply_status'])

	let form = ref<type_system_log_sms_query>({
    scene: null,
    mobile: null,
    content: null,
    status: null,
    results: null,
    send_time_start: null,
    send_time_end: null,
    create_time_start: null,
    create_time_end: null,
    update_time_start: null,
    update_time_end: null,
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