<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form" :rules="rules" ref="formRef">
			<uv-form-item label="IMEI" prop="imei" borderBottom>
				<uv-input v-model="form.imei" border="surround">
				</uv-input>
			</uv-form-item>
			<uv-form-item label="位置" prop="location" borderBottom>
				<uv-input v-model="form.location" border="surround">
				</uv-input>
			</uv-form-item>

			<uv-form-item label="关联表号" prop="meters" borderBottom>
				<uv-textarea v-model="form.meters" border="surround"></uv-textarea>
			</uv-form-item>


			<uv-form-item label="卡1" prop="sim0" borderBottom>
				<uv-input v-model="form.sim0" border="surround">
				</uv-input>
			</uv-form-item>
			<uv-form-item label="卡2" prop="sim1" borderBottom>
				<uv-input v-model="form.sim1" border="surround">
				</uv-input>
			</uv-form-item>
			<uv-form-item label="备注" prop="remark" borderBottom>
				<uv-textarea v-model="form.remark" border="surround"></uv-textarea>
			</uv-form-item>

			<uv-button   type="primary" text="提交" customStyle="margin-top: 20rpx"
				@click="submit"></uv-button>



		</uv-form>


	</view>
</template>

<script setup>
	import {
		reactive,
		ref,
		computed
	} from "vue";
	import {
		onLoad
	} from "@dcloudio/uni-app";
	import {
		equipment_list,
		equipment_edit,

	} from "@/api/equipment.js";
 
	import {
		toast,
		alert
	} from "@/utils/utils";

	let formRef = ref();
	let form = ref({});
	let rules = {};




	onLoad((e) => {
		console.log("onLoad", e);
		//   form.value = e;
		getDetails(e.id);
	});


	function getDetails(id) {
		equipment_list({
				id
			})
			.then((res) => {
				if (res.code == 200) {
					// form.value = e;
					if (res?.result?.records?.length == 1) {
						form.value = res?.result?.records[0];
					}
				}
			})
			.catch((err) => {});
	}

	function submit(item) {
		console.log("submit", form);
		equipment_edit(form.value).then((res) => {
			if (res.code == 200) {
				toast("编辑成功");
		 
				getDetails(form.value?.id);
			} else {
				toast(res.message);
			}
		});
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>