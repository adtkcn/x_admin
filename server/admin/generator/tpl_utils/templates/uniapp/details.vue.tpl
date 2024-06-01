<template>
	<view class="page-content">
        <uv-navbar leftText="" :placeholder="true" title="{{{.FunctionName}}}详情" autoBack>
            <template v-slot:right>
                <uv-icon name="edit-pen" size="20" @click="edit"></uv-icon>
            </template>
        </uv-navbar>

		<uv-form labelPosition="left" :model="form">
        {{{- range .Columns }}}
        {{{- if .IsList }}}
            <uv-form-item label="{{{.ColumnComment}}}" prop="{{{(toCamelCase .GoField)}}}" borderBottom>
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                    <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{ (toCamelCase .GoField) }}}" />
                {{{- else if eq .HtmlType "imageUpload" }}}
                    <uv-image :src="$filePath(form.{{{(toCamelCase .GoField)}}})" width="100%"></uv-image>
                {{{- else }}}
                    {{form.{{{(toCamelCase .GoField)}}}}}
                {{{- end }}}
            </uv-form-item>
        {{{- end }}}
        {{{- end }}}
		</uv-form>
        <uv-button
            v-if="$perms('admin:{{{ .ModuleName }}}:edit')"
            type="primary"
            text="编辑"
            customStyle="margin-top: 20rpx"
            @click="edit"
        ></uv-button>

	</view>
</template>

<script setup>
	import {
		reactive,
		ref,
		computed
	} from "vue";
	import {
		onLoad,
		onShow
	} from "@dcloudio/uni-app";
	import {
		useDictOptions
	} from "@/api/dict.js";
	import {
		{{{ .ModuleName }}}_detail 
	} from "@/api/{{{ .ModuleName }}}.js";


	import {
		toast,
		alert,
		toPath
	} from "@/utils/utils";

	let form = ref({});

	onLoad((e) => {
		console.log("onLoad", e);
		getDetails(e.id);
	});
	onShow((e) => {
		if (form.value?.id) {
			getDetails(form.value.id);
		}
	});

	function getDetails(id) {
		{{{ .ModuleName }}}_detail(id).then((res) => {
            if (res.code == 200) {
                if (res?.result) {
                    form.value = res?.result
                }
            } else {
                toast(res.message);
            }
        })
        .catch((err) => {
            toast(err.message||"网络错误");
        });
	}

	function edit() {
		toPath("/pages/{{{ .ModuleName }}}/edit", { id: form.value.id });
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>