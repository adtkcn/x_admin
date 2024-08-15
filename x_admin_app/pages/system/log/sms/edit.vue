<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form" :rules="formRules" ref="formRef">			
			<uv-form-item label="id" prop="id" borderBottom>
					<uv-number-box v-model="form.id" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
			</uv-form-item>			
			<uv-form-item label="场景编号" prop="scene" borderBottom>
					<uv-number-box v-model="form.scene" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
			</uv-form-item>			
			<uv-form-item label="手机号码" prop="mobile" borderBottom>
					<uv-input v-model="form.mobile" border="surround"></uv-input>
			</uv-form-item>			
			<uv-form-item label="发送内容" prop="content" borderBottom>
			</uv-form-item>			
			<uv-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="status" borderBottom>
						<x-picker v-model="form.status" valueKey="value" labelKey="name" :columns="dictData.flow_apply_status"></x-picker>
			</uv-form-item>			
			<uv-form-item label="短信结果" prop="results" borderBottom>
					<uv-textarea v-model="form.results" border="surround"></uv-textarea>
			</uv-form-item>			
			<uv-form-item label="发送时间" prop="send_time" borderBottom>
					<x-date v-model:time="form.send_time"></x-date>
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
    id: '',
    scene: 0,
    mobile: '',
    content: '',
    status: '',
    results: '',
    send_time: '',
	});
	const formRules = {
		id: [
			{
				required: true,
				message: '请输入id',
				trigger: ['blur']
			}
		],
		scene: [
			{
				required: true,
				message: '请输入场景编号',
				trigger: ['blur']
			}
		],
		mobile: [
			{
				required: true,
				message: '请输入手机号码',
				trigger: ['blur']
			}
		],
		content: [
			{
				required: true,
				message: '请输入发送内容',
				trigger: ['blur']
			}
		],
		status: [
			{
				required: true,
				message: '请选择发送状态：[0=发送中, 1=发送成功, 2=发送失败]',
				trigger: ['blur']
			}
		],
		results: [
			{
				required: true,
				message: '请输入短信结果',
				trigger: ['blur']
			}
		],
		send_time: [
			{
				required: true,
				message: '请选择发送时间',
				trigger: ['blur']
			}
		],
	}
	onLoad((e) => {
		console.log("onLoad", e);
		if (e.id) {
			getDetails(e.id);
		}
	});
	const { dictData } = useDictData<{
    flow_apply_status: type_dict[]
}>(['flow_apply_status'])

	function getDetails(id) {
		system_log_sms_detail(id).then((res) => {
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
			if (form.value.id) {
				system_log_sms_edit(form.value).then((res) => {
					if (res.code == 200) {
						toast("编辑成功");				
						getDetails(form.value?.id);
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