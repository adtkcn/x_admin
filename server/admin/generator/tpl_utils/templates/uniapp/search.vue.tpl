<template>
	<view class="page-content">
		<uv-form labelPosition="left" labelWidth="80" :model="form">
			{{{- range .Columns }}}
			{{{- if eq .IsQuery 1 }}}
			<uv-form-item label="{{{ .ColumnComment }}}" prop="{{{ (.GoField) }}}" borderBottom>
				{{{- if eq .HtmlType "datetime" }}}
					<x-date-range v-model:startTime="form.{{{ (.GoField) }}}_start"
							v-model:endTime="form.{{{ (.GoField) }}}_end"></x-date-range>
				{{{- else if or (eq .HtmlType "checkbox") (eq .HtmlType "radio") (eq .HtmlType "select") }}}
					{{{- if ne .DictType "" }}}
						<x-picker v-model="form.{{{ (.GoField) }}}" valueKey="value" labelKey="name" :columns="dictData.{{{ .DictType }}}"></x-picker>
					{{{- else if ne .ListAllApi "" }}}
							<x-picker v-model="form.{{{ (.GoField) }}}" valueKey="id" labelKey="id" :columns="listAllData.{{{pathToName .ListAllApi}}}"></x-picker>
					{{{- end }}}
				{{{- else if eq .HtmlType "input" }}}
					<uv-input v-model="form.{{{ (.GoField) }}}"> </uv-input>
				{{{- end }}}
			</uv-form-item>
			{{{- end }}}
			{{{- end }}}

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
	import type {type_{{{.ModuleName}}}_query} from "@/api/monitor_project";
{{{- if ge (len .DictFields) 1 }}}
{{{- $dictSize := sub (len .DictFields) 1 }}}
const { dictData } = useDictData<{
    {{{- range .DictFields }}}
    {{{ . }}}: any[]
    {{{- end }}}
}>([{{{- range .DictFields }}}'{{{ . }}}'{{{- if ne (index $.DictFields $dictSize) . }}},{{{- end }}}{{{- end }}}])
{{{- end }}}

{{{- if ge (len .ListAllFields) 1 }}}
{{{- $list_all_size := sub (len .ListAllFields) 1 }}}
const { listAllData } = useListAllData<{
    {{{- range .ListAllFields }}}
    {{{pathToName . }}}: any[]
    {{{- end }}}
}>({ 
	{{{- range .ListAllFields }}}
		{{{pathToName . }}}:'{{{deletePathPrefix . }}}',
	{{{- end }}}
})
{{{- end }}}

	let form = ref<type_{{{.ModuleName}}}_query>({
{{{- range .Columns }}}
{{{- if .IsQuery }}}
    {{{- if eq .HtmlType "datetime" }}}
    {{{ (.GoField) }}}_start: null,
    {{{ (.GoField) }}}_end: null,
    {{{- else }}}
    {{{ (.GoField) }}}: null,
    {{{- end }}}
{{{- end }}}
{{{- end }}}
	});

	function submit() {
		console.log("submit", form.value);

		const search = clearObjEmpty(form.value);

		if (Object.keys(search).length === 0) {
			return toast("请输入查询条件");
		}

		toPath("/pages/{{{nameToPath .ModuleName }}}/index", search);
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>