<!-- 产品，分组，IMEI，SIm卡1、2，安装位置，在线状态 ，启用状态-->
<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form"  ref="formRef">
			<uv-form-item label="在线状态" prop="status" borderBottom>
				<x-picker valueKey="value" labelKey="label" :columns="dictData.devices_status"
					v-model="form.status"></x-picker>
			</uv-form-item>

	 
			<uv-form-item label="位置" prop="placedName" borderBottom>
				<uv-input v-model="form.placedName"> </uv-input>
			</uv-form-item>
 
			<uv-form-item label="开始时间" prop="startTime" borderBottom>
				<uv-input v-model="form.startTime" :readonly="true" placeholder="请选择时间"
					@click="startTimePicker.open()"> </uv-input>
				<uv-datetime-picker ref="startTimePicker" :minDate="minDate" :maxDate="maxDate" :value="form.startTime"
					mode="datetime" @confirm="startTimeConfirm">
				</uv-datetime-picker>
			</uv-form-item>
			<uv-form-item label="结束时间" prop="endTime" borderBottom>

				<uv-input v-model="form.endTime" :readonly="true" placeholder="请选择时间"
					@click="endTimePicker.open()"> </uv-input>
				<uv-datetime-picker ref="endTimePicker" :minDate="minDate" :maxDate="maxDate" :value="form.endTime" mode="datetime"
					@confirm="endTimeConfirm">
				</uv-datetime-picker>
			</uv-form-item>

			<uv-button type="primary" text="搜索" customStyle="margin-top: 20rpx" @click="submit"></uv-button>
		</uv-form>
	</view>
</template>

<script setup>
	import dayjs from 'dayjs'
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
	} from "@/methods/useDictOptions";
	const startTimePicker = ref();
	const endTimePicker = ref();
	
	const {
		dictData
	} = useDictData(["devices_status"]);

	let formRef = ref();
	let form = ref({
	 
		location: "",

		placedName: "",
		status: "",
		startTime: '',
		endTime: '',
	});
	let minDate = dayjs('2022-01-01 00:00:00').valueOf()
	let maxDate = dayjs().endOf('month').valueOf()
	function startTimeConfirm(e) {		
		form.value.startTime = dayjs(e.value).format('YYYY-MM-DD HH:mm:ss')
	}

	function endTimeConfirm(e) { 
		form.value.endTime = dayjs(e.value).format('YYYY-MM-DD HH:mm:ss')
	}
 

	function submit() {
		console.log("submit", form.value);

		const search = clearObjEmpty(form.value);

		if (Object.keys(search).length === 0) {
			return toast("请输入查询条件");
		}

		toPath("/pages/equipment/equipment", search);
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>