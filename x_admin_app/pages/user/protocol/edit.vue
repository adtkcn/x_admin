<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form" :rules="formRules" ref="formRef">			
			<uv-form-item label="" prop="Id" borderBottom>
					<uv-number-box v-model="form.Id" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
			</uv-form-item>			
			<uv-form-item label="标题" prop="Title" borderBottom>
					<uv-input v-model="form.Title" border="surround"></uv-input>
			</uv-form-item>			
			<uv-form-item label="协议内容" prop="Content" borderBottom>
			</uv-form-item>			
			<uv-form-item label="排序" prop="Sort" borderBottom>
					<uv-number-box v-model="form.Sort" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
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
		user_protocol_detail,
		user_protocol_edit,
		user_protocol_add
	} from "@/api/user_protocol";
	import type { type_user_protocol_edit	} from "@/api/user/protocol";

	import {
		toast,
		alert
	} from "@/utils/utils";
	import {
		useDictData,useListAllData
	} from "@/hooks/useDictOptions";
	import type { type_dict } from '@/hooks/useDictOptions'

	let formRef = ref();
	let form = ref<type_user_protocol_edit>({
    Id: 0,
    Title: '',
    Content: '',
    Sort: 0,
	});
	const formRules = {
		Id: [
			{
				required: true,
				message: '请输入',
				trigger: ['blur']
			}
		],
		Title: [
			{
				required: true,
				message: '请输入标题',
				trigger: ['blur']
			}
		],
		Content: [
			{
				required: true,
				message: '请输入协议内容',
				trigger: ['blur']
			}
		],
		Sort: [
			{
				required: true,
				message: '请输入排序',
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
		user_protocol_detail(Id).then((res) => {
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
				user_protocol_edit(form.value).then((res) => {
					if (res.code == 200) {
						toast("编辑成功");				
						getDetails(form.value?.Id);
					} else {
						toast(res.message);
					}
				});
			}else{
				user_protocol_add(form.value).then((res) => {
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