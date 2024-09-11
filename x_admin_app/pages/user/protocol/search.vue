<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form">
			<uv-form-item label="标题" prop="Title" borderBottom>
					<uv-input v-model="form.Title"> </uv-input>
			</uv-form-item>
			<uv-form-item label="协议内容" prop="Content" borderBottom>
			</uv-form-item>
			<uv-form-item label="排序" prop="Sort" borderBottom>
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
	import type {type_user_protocol_query} from "@/api/user/protocol";

	let form = ref<type_user_protocol_query>({
    Title: null,
    Content: null,
    Sort: null,
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

		toPath("/pages/user/protocol/index", search);
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>