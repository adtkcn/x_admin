<template>
    <div class="material-index">
        <el-card class="!border-none" shadow="never">
            <el-tabs v-model="activeTab">
                <el-tab-pane
                    v-for="item in tabsMap"
                    :label="item.name"
                    :name="item.name"
                    :index="item.name"
                    :key="item.name"
                    lazy
                >
                    <material
                        :ext="item.ext"
                        mode="page"
                        file-size="120px"
                        :limit="-1"
                        :page-size="10"
                    />
                </el-tab-pane>
            </el-tabs>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import material from '@/components/material/index.vue'
defineOptions({
    name: 'materialCenter'
})
const tabsMap = [
    {
        // type: '',
        name: '全部',
        ext: []
    },
    {
        // type: 'image',
        name: '图片',
        ext: ['jpg', 'jpeg', 'png', 'gif', 'bmp']
    },
    {
        // type: 'video',
        name: '视频',
        ext: ['mp4', 'avi', 'mov']
    },
    {
        // type: 'pdf',
        name: 'pdf',
        ext: ['pdf']
    }
]
const activeTab = ref('image')
</script>

<style lang="scss" scoped>
.material-index {
    min-width: 700px;
    :deep(.el-tabs) {
        height: calc(100vh - 180px);

        .el-tabs__content,
        .el-tab-pane {
            min-height: 0;
            flex: 1;
            display: flex;
            flex-direction: column;
        }
    }
}
</style>
