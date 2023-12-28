<template>
    <template v-if="!route.meta?.hidden">
        <app-link v-if="!hasShowChild" :to="`${routePath}?${queryStr}`">
            <el-menu-item :index="routePath">
                <Icon :size="16" v-if="routeMeta?.icon" :name="routeMeta?.icon" />
                <template #title>
                    <span>{{ routeMeta?.title }}</span>
                </template>
            </el-menu-item>
        </app-link>
        <el-sub-menu v-else :index="routePath" :popper-class="props.popperClass">
            <template #title>
                <Icon :size="16" v-if="routeMeta?.icon" :name="routeMeta?.icon" />
                <span>{{ routeMeta?.title }}</span>
            </template>
            <menu-item
                v-for="item in route?.children || []"
                :key="resolvePath(item.path)"
                :route="item"
                :route-path="resolvePath(item.path)"
                :popper-class="props.popperClass"
            />
        </el-sub-menu>
    </template>
</template>

<script lang="ts" setup>
import { getNormalPath, objectToQuery } from '@/utils/util'
import { isExternal } from '@/utils/validate'
import type { RouteRecordRaw } from 'vue-router'
interface Props {
    route: RouteRecordRaw
    routePath: string
    popperClass: string
}
defineOptions({
    name: 'MenuItem'
})
const props = defineProps<Props>()

const hasShowChild = computed(() => {
    const children: RouteRecordRaw[] = props.route.children ?? []
    return !!children.filter((item) => !item.meta?.hidden).length
})

const routeMeta = computed(() => {
    return props?.route?.meta
})

const resolvePath = (path: string) => {
    if (isExternal(path)) {
        return path
    }
    const newPath = getNormalPath(`${props.routePath}/${path}`)
    return newPath
}
const queryStr = computed<string>(() => {
    const query = props.route.meta?.query as string
    try {
        const queryObj = JSON.parse(query)
        return objectToQuery(queryObj)
    } catch (error) {
        // console.log(error)

        return query
    }
})
</script>
<style lang="scss" scoped></style>
