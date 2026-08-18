<template>
    <div :style="{ width: width, height: height }">
        <video ref="videoRef" class="video-js vjs-bootstrap5"></video>
    </div>
</template>

<script setup lang="ts">
import { onBeforeUnmount, onMounted, useTemplateRef, watch } from 'vue'
import videojs from 'video.js'
import type Player from 'video.js/dist/types/player'
import 'video.js/dist/video-js.css'

const props = defineProps({
    src: {
        type: String,
        required: true
    },
    width: String,
    height: String,
    poster: String
})

const videoRef = useTemplateRef<HTMLVideoElement>('videoRef')
let player: Player | null = null

const initPlayer = () => {
    if (!videoRef.value) return
    player = videojs(videoRef.value, {
        controls: true,
        autoplay: false,
        muted: false,
        loop: false,
        preload: 'auto',
        fill: true,
        playbackRates: [0.75, 1.0, 1.25, 1.5, 2.0],
        poster: props.poster,
        sources: [
            {
                src: props.src,
                type: 'video/mp4'
            }
        ]
    })

    player.on('play', (event: any) => {
        console.log(event, '播放')
    })
    player.on('pause', (event: any) => {
        console.log(event, '暂停')
    })
    player.on('timeupdate', (event: any) => {
        // console.log(event, '时间更新')
    })
    player.on('canplay', (event: any) => {
        console.log(event, '可以播放')
    })
}

watch(
    () => props.src,
    (src) => {
        player?.src({ src, type: 'video/mp4' })
    }
)

onMounted(() => {
    initPlayer()
})

onBeforeUnmount(() => {
    player?.dispose()
    player = null
})

const play = () => {
    player?.play()
}

const pause = () => {
    player?.pause()
}

defineExpose({
    play,
    pause
})
</script>
