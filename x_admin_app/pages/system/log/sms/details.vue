<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form">
            <uv-form-item label="场景编号" prop="Scene" borderBottom>
                    {{form.Scene}}
            </uv-form-item>
            <uv-form-item label="手机号码" prop="Mobile" borderBottom>
                    {{form.Mobile}}
            </uv-form-item>
            <uv-form-item label="发送内容" prop="Content" borderBottom>
                    {{form.Content}}
            </uv-form-item>
            <uv-form-item label="发送状态：[0=发送中, 1=发送成功, 2=发送失败]" prop="Status" borderBottom>
                    {{form.Status}}
            </uv-form-item>
            <uv-form-item label="短信结果" prop="Results" borderBottom>
                    {{form.Results}}
            </uv-form-item>
            <uv-form-item label="发送时间" prop="SendTime" borderBottom>
                    {{form.SendTime}}
            </uv-form-item>
            <uv-form-item label="创建时间" prop="CreateTime" borderBottom>
                    {{form.CreateTime}}
            </uv-form-item>
            <uv-form-item label="更新时间" prop="UpdateTime" borderBottom>
                    {{form.UpdateTime}}
            </uv-form-item>
		</uv-form>
        <uv-button
            v-if="$perms('admin:system_log_sms:edit')"
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
	import { useDictData,useListAllData } from "@/hooks/useDictOptions";
	import { system_log_sms_detail } from "@/api/system_log_sms";


	import {
		toast,
		alert,
		toPath
	} from "@/utils/utils";

	let form = ref({
		Id: "",
		Scene: "",
		Mobile: "",
		Content: "",
		Status: "",
		Results: "",
		SendTime: "",
		CreateTime: "",
		UpdateTime: "",
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
		system_log_sms_detail(id).then((res) => {
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
		toPath("/pages/system/log/sms/edit", { id: form.value.id });
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>