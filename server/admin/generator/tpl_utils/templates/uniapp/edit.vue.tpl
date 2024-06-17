<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form" :rules="formRules" ref="formRef">
			{{{- range .Columns }}}
            {{{- if .IsEdit }}}			
			<uv-form-item label="{{{ .ColumnComment }}}" prop="{{{ (toCamelCase .GoField) }}}" borderBottom>
                {{{- if eq .HtmlType "input" }}}
					<uv-input v-model="form.{{{ (toCamelCase .GoField) }}}" border="surround"></uv-input>
				{{{- else if eq .HtmlType "number" }}}
					<uv-number-box v-model="form.{{{ (toCamelCase .GoField) }}}" :min="-99999999" :max="99999999" :integer="true"></uv-number-box>
				{{{- else if eq .HtmlType "textarea" }}}
					<uv-textarea v-model="form.{{{ (toCamelCase .GoField) }}}" border="surround"></uv-textarea>
				{{{- else if or (eq .HtmlType "checkbox") (eq .HtmlType "radio") (eq .HtmlType "select")}}}
				{{{- if ne .DictType "" }}}
					<x-picker v-model="form.{{{ (toCamelCase .GoField) }}}" valueKey="value" labelKey="name" :columns="dictData.{{{ .DictType }}}"></x-picker>
				{{{- else }}}
					请选择字典生成代码
				{{{- end }}}
				{{{- end }}}
			</uv-form-item>
			{{{- end }}}
			{{{- end }}}

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
		{{{ .ModuleName }}}_detail,
		{{{ .ModuleName }}}_edit,
		{{{ .ModuleName }}}_add
	} from "@/api/{{{ .ModuleName }}}";
	import type { type_{{{ .ModuleName }}}_edit	} from "@/api/{{{ .ModuleName }}}";

	import {
		toast,
		alert
	} from "@/utils/utils";
	import {
		useDictData
	} from "@/hooks/useDictOptions";
	
	let formRef = ref();
	let form = ref<type_{{{ .ModuleName }}}_edit>({
	{{{- range .Columns }}}
    {{{- if eq (toCamelCase .GoField) $.PrimaryKey }}}
    {{{ $.PrimaryKey }}}: '',
    {{{- else if .IsEdit }}}
    {{{- if eq .HtmlType "checkbox" }}}
    {{{ (toCamelCase .GoField) }}}: [],
    {{{- else if eq .HtmlType "number" }}}
    {{{ (toCamelCase .GoField) }}}: 0,
    {{{- else }}}
    {{{ (toCamelCase .GoField) }}}: '',
    {{{- end }}}
    {{{- end }}}
    {{{- end }}}
	});
	const formRules = {
		{{{- range .Columns }}}
		{{{- if and .IsEdit .IsRequired }}}
		{{{ (toCamelCase .GoField) }}}: [
			{
				required: true,
				{{{- if or (eq .HtmlType "checkbox") (eq .HtmlType "datetime") (eq .HtmlType "radio") (eq .HtmlType "select") (eq .HtmlType "imageUpload") }}}
				message: '请选择{{{ .ColumnComment }}}',
				{{{- else }}}
				message: '请输入{{{ .ColumnComment }}}',
				{{{- end }}}
				trigger: ['blur']
			}
		],
		{{{- end }}}
		{{{- end }}}
	}
	onLoad((e) => {
		console.log("onLoad", e);
		if (e.id) {
			getDetails(e.id);
		}
	});


	{{{- if ge (len .DictFields) 1 }}}
	{{{- $dictSize := sub (len .DictFields) 1 }}}
	const { dictData } = useDictData<{
    {{{- range .DictFields }}}
    {{{ . }}}: any[]
    {{{- end }}}
}>([{{{- range .DictFields }}}'{{{ . }}}'{{{- if ne (index $.DictFields $dictSize) . }}},{{{- end }}}{{{- end }}}])
	{{{- end }}}

	function getDetails(id) {
		{{{ .ModuleName }}}_detail(id).then((res) => {
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
				{{{ .ModuleName }}}_edit(form.value).then((res) => {
					if (res.code == 200) {
						toast("编辑成功");				
						getDetails(form.value?.id);
					} else {
						toast(res.message);
					}
				});
			}else{
				{{{ .ModuleName }}}_add(form.value).then((res) => {
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