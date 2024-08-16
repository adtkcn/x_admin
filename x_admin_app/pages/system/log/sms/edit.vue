<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form" :rules="formRules" ref="formRef">			
			<uv-form-item label="id" prop="Id" borderBottom>
					<uv-number-box v-model="form.Id" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
			</uv-form-item>			
			<uv-form-item label="场景编号" prop="Scene" borderBottom>
					<uv-number-box v-model="form.Scene" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
			</uv-form-item>			
			<uv-form-item label="手机号码" prop="Mobile" borderBottom>
					<uv-input v-model="form.Mobile" border="surround"></uv-input>
			</uv-form-item>			
			<uv-form-item label="发送内容" prop="Content" borderBottom>
			</uv-form-item>			
			<uv-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="Status" borderBottom>
						请选择字典生成代码
			</uv-form-item>			
			<uv-form-item label="短信结果" prop="Results" borderBottom>
					<uv-textarea v-model="form.Results" border="surround"></uv-textarea>
			</uv-form-item>			
			<uv-form-item label="发送时间" prop="SendTime" borderBottom>
					<x-date v-model:time="form.SendTime"></x-date>
			</uv-form-item>

			<uv-button   type="primary" text="提交" customStyle="margin: 40rpx 0"
				@click="submit"></uv-button>
		</uv-form>
	</view>
</template>

<script setup lang="ts">
	import { ref } from "vue";
	import {
		onLoad
	} from "@dcloudio/uni-app";
 	import {
		system_log_sms_detail,
		system_log_sms_edit,
		system_log_sms_add
	} from "@/api/system_log_sms";
	import type { type_system_log_sms_edit	} from "@/api/system_log_sms";

	import {
		toast,
		alert
	} from "@/utils/utils";
	import {
		useDictData,useListAllData
	} from "@/hooks/useDictOptions";
	import type { type_dict } from '@/hooks/useDictOptions'

	let formRef = ref();
	let form = ref<type_system_log_sms_edit>({
    Id: 0,
    Scene: 0,
    Mobile: '',
    Content: '',
    Status: '',
    Results: '',
    SendTime: '',
	});
	const formRules = {
		Id: [
			{
				required: true,
				message: '请输入id',
				trigger: ['blur']
			}
		],
		Scene: [
			{
				required: true,
				message: '请输入场景编号',
				trigger: ['blur']
			}
		],
		Mobile: [
			{
				required: true,
				message: '请输入手机号码',
				trigger: ['blur']
			}
		],
		Content: [
			{
				required: true,
				message: '请输入发送内容',
				trigger: ['blur']
			}
		],
		Status: [
			{
				required: true,
				message: '请选择发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
				trigger: ['blur']
			}
		],
		Results: [
			{
				required: true,
				message: '请输入短信结果',
				trigger: ['blur']
			}
		],
		SendTime: [
			{
				required: true,
				message: '请选择发送时间',
				trigger: ['blur']
			}
		],
	}
	onLoad((e) => {
		console.log("onLoad", e);
		if (e.Id) {
			getDetails(e.Id);
		}
	});

	function getDetails(Id) {
		system_log_sms_detail(Id).then((res) => {
            if (res.code == 200) {
                if (res?.data) {
                    form.value = res?.data
                }
            } else {
                toast(res.message);
            }
        })
        .catch((err) => {
            toast(err.message||"网络错误");
        });
	}

	function submit() {
		console.log("submit", form.value);
		formRef.value.validate().then(() => {
			if (form.value.Id) {
				system_log_sms_edit(form.value).then((res) => {
					if (res.code == 200) {
						toast("编辑成功");				
						getDetails(form.value?.Id);
					} else {
						toast(res.message);
					}
				});
			}else{
				system_log_sms_add(form.value).then((res) => {
					if (res.code == 200) {
						toast("添加成功");
						uni.navigateBack();
					} else {
						toast(res.message);
					}
				});
			}

		})
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>