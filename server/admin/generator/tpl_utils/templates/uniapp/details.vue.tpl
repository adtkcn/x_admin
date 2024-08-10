<template>
	<view class="page-content">
		<uv-form labelPosition="left" :model="form">
        {{{- range .Columns }}}
        {{{- if .IsList }}}
            <uv-form-item label="{{{.ColumnComment}}}" prop="{{{(.GoField)}}}" borderBottom>
                {{{- if and (ne .DictType "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
                    <dict-value :options="dictData.{{{ .DictType }}}" :value="row.{{{ (.GoField) }}}" />
				{{{- else if and (ne .ListAllApi "") (or (eq .HtmlType "select") (eq .HtmlType "radio") (eq .HtmlType "checkbox")) }}}
				 	<dict-value :options="listAllData.{{{pathToName .ListAllApi }}}" :value="row.{{{ (.GoField) }}}" labelKey='id' valueKey='id' />
                {{{- else if eq .HtmlType "imageUpload" }}}
                    <uv-image :src="$filePath(form.{{{(.GoField)}}})" width="100%"></uv-image>
                {{{- else }}}
                    {{form.{{{(.GoField)}}}}}
                {{{- end }}}
            </uv-form-item>
        {{{- end }}}
        {{{- end }}}
		</uv-form>
        <uv-button
            v-if="$perms('admin:{{{ .ModuleName }}}:edit')"
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
	import { {{{ .ModuleName }}}_detail } from "@/api/{{{ .ModuleName }}}";


	import {
		toast,
		alert,
		toPath
	} from "@/utils/utils";

	let form = ref({
	{{{- range .Columns }}}
    {{{- if or .IsList .IsPk }}}
		{{{ .GoField }}}: "",
	{{{- end }}}
    {{{- end }}}
	});
	
{{{- if ge (len .DictFields) 1 }}}
{{{- $dictSize := sub (len .DictFields) 1 }}}
const { dictData } = useDictData<
{
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
		{{{ .ModuleName }}}_detail(id).then((res) => {
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
		toPath("/pages/{{{nameToPath .ModuleName }}}/edit", { id: form.value.id });
	}
</script>

<style lang="scss" scoped>
	.page-content {
		padding: 10rpx 20rpx 300rpx;
	}
</style>