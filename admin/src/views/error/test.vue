<script setup lang="ts">
import XLog from '../../../../x_error_sdk/x-log'
import x_log_browser from '../../../../x_error_sdk/x-log-browser'
const xlog = new XLog(
    {
        Dns: 'http://localhost:5174/api',
        Pid: 'e19e3be20de94f49b68fafb4c30668bc',
        Uid: '10'
    },
    x_log_browser
)
defineProps<{ msg: string }>()
function error() {
    throw new Error('主动抛error')
}
function promiseError() {
    Promise.reject('主动抛promise1')
    Promise.reject(new Error('主动抛promiseError2'))
}
function yufaError() {
    // try {
    setTimeout(() => {
        a = 1
        // a = 2;
    }, 100)
    // } catch (error) {
    // 	console.log(error);
    // }
}
function loadError() {
    const img = new Image()
    console.log(img)
    img.src = 'http://a.cn'
    document.body.appendChild(img)
}
</script>

<template>
    <h1>{{ msg }}</h1>

    <div class="card">
        <button class="a b" type="button" aa="11" @click="error">主动抛error</button>
        <button class="a b" type="button" aa="11" @click="yufaError">语法错误</button>
        <button class="a b c" id="abv" type="button" @click="promiseError">promiseerror</button>
        <button type="button" @click="loadError">资源加载错误</button>

        <!-- <img src="http://a.cn/a/b" alt="" srcset="" />
		<link rel="stylesheet" href="http://a.cn/c/c/c" /> -->
    </div>
</template>

<style scoped>
button {
    color: #fff;
    background-color: rgb(150, 149, 149);
    border: 0;
    padding: 6px 10px;
    margin: 6px;
    font-size: 14px;
    line-height: 18px;
}
</style>
