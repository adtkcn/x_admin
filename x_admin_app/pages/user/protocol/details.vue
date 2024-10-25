<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form">
            <uv-form-item label="标题" prop="Title" borderBottom>
                    {{form.Title}}
            </uv-form-item>
            <uv-form-item label="协议内容" prop="Content" borderBottom>
                    {{form.Content}}
            </uv-form-item>
            <uv-form-item label="排序" prop="Sort" borderBottom>
                    {{form.Sort}}
            </uv-form-item>
            <uv-form-item label="创建时间" prop="CreateTime" borderBottom>
                    {{form.CreateTime}}
            </uv-form-item>
            <uv-form-item label="更新时间" prop="UpdateTime" borderBottom>
                    {{form.UpdateTime}}
            </uv-form-item>
		</uv-form>
        <uv-button
            v-if="$perms('admin:user_protocol:edit')"
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
	import { user_protocol_detail } from "@/api/user/protocol";


	import {
		toast,
		alert,
		toPath
	} from "@/utils/utils";

	let form = ref({
		Id: "",
		Title: "",
		Content: "",
		Sort: "",
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
		user_protocol_detail(id).then((res) => {
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
		toPath("/pages/user/protocol/edit", { id: form.value.id });
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>