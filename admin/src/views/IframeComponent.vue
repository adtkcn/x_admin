vue
<template>
    <div class="h-full">
        <!-- {{ url }} -->
        <iframe :src="url" frameborder="0" width="100%" height="100%"></iframe>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
const route = useRoute()
const url = computed(() => {
    return route.query.url as string
})

// const props = defineProps({
//     url: {
//         type: String,
//         default: ''
//     }
// })
</script>
