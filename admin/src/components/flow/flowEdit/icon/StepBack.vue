<template>
    <img class="icon" :src="iconUrl" :width="size" :height="size" alt="" />
</template>

<script setup lang="ts">
import iconUrl from './StepBack.svg?url'

defineProps<{
    size?: string
}>()
</script>
