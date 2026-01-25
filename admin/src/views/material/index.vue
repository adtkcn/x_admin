<template>
    <div class="material-index">
        <MaterialComponent mode="page" :limit="-1" :page-size="10" />
    </div>
</template>

<script lang="ts" setup>
// import MaterialComponent from '@/components/material/index.vue'
import { defineAsyncComponent } from 'vue'
const MaterialComponent = defineAsyncComponent(() => import('@/components/material/index.vue'))

defineOptions({
    name: 'materialCenter'
})
</script>

<style lang="scss" scoped>
.material-index {
    height: calc(100vh - 135px);
    overflow: hidden;
    background-color: white;
    padding: 0 5px 5px;
}
</style>
