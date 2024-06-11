<template>
    <div style="position: relative">
        <div
            class="verify-img-out"
            :style="{ height: parseInt(props.imgSize.height) + vSpace + 'px' }"
        >
            <div
                class="verify-img-panel"
                :style="{ width: props.imgSize.width + 'px', height: props.imgSize.height + 'px' }"
            >
                <img
                    v-show="backImgBase"
                    :src="backImgBase"
                    style="width: 100%; height: 100%; display: block"
                />
                <div
                    class="verify-sub-block"
                    :style="{
                        left: moveBlockLeft
                    }"
                >
                    <img
                        v-show="blockBackImgBase"
                        :src="blockBackImgBase"
                        style="width: 100%; height: 100%; display: block; -webkit-user-drag: none"
                    />
                </div>

                <div class="verify-refresh" @click="refresh" v-show="showRefresh">
                    <i class="iconfont icon-refresh"></i>
                </div>
                <transition name="tips">
                    <span
                        class="verify-tips"
                        v-if="tipWords"
                        :class="passFlag ? 'suc-bg' : 'err-bg'"
                        >{{ tipWords }}</span
                    >
                </transition>
            </div>
        </div>
        <!-- 公共部分 -->
        <div
            class="verify-bar-area"
            :style="{
                width: props.imgSize.width + 'px',
                height: blockSize.height + 'px',
                'line-height': blockSize.height + 'px'
            }"
        >
            <span class="verify-msg" v-text="text"></span>
            <div
                class="verify-left-bar"
                :style="{
                    width: leftBarWidth !== undefined ? leftBarWidth : blockSize.height + 'px',
                    height: blockSize.height + 'px',
                    'border-color': leftBarBorderColor,
                    transition: transitionWidth
                }"
            >
                <span class="verify-msg" v-text="finishText"></span>
                <div
                    class="verify-move-block"
                    @touchstart="start"
                    @mousedown="start"
                    :style="{
                        width: blockSize.width + 'px',
                        height: blockSize.height + 'px',
                        'background-color': moveBlockBackgroundColor,
                        left: moveBlockLeft,
                        transition: transitionLeft
                    }"
                >
                    <i
                        :class="['verify-icon iconfont', iconClass]"
                        :style="{ color: iconColor }"
                    ></i>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts">
/**
 * VerifySlide
 * @description 滑块
 * */
import { aesEncrypt } from '../utils/ase'
// import { resetSize } from './../utils/util'
import { reqGet, reqCheck } from './../api/index'
import { computed, onMounted, onUnmounted, ref, nextTick, toRefs, getCurrentInstance } from 'vue'
defineOptions({
    name: 'VerifySlide'
})
const emit = defineEmits(['success', 'error'])
const props = defineProps({
    captchaType: {
        type: String
    },

    //弹出式pop，固定fixed
    mode: {
        type: String,
        default: 'fixed'
    },
    vSpace: {
        type: Number,
        default: 5
    },
    explain: {
        type: String,
        default: '向右滑动完成验证'
    },
    imgSize: {
        type: Object,
        default() {
            return {
                width: 310,
                height: 155
            }
        }
    },
    blockSize: {
        type: Object,
        default() {
            return {
                width: 47,
                height: 47
            }
        }
    }
})
const blockSize = computed(() => {
    return {
        width: (props.imgSize.width / 310) * props.blockSize.width,
        height: props.blockSize.height
    }
})
// const defaultImgSize = {
//     width: 310,
//     height: 155
// }

const { mode, explain } = toRefs(props)
const { proxy } = getCurrentInstance()
const secretKey = ref(''), //后端返回的ase加密秘钥
    passFlag = ref(false), //是否通过的标识
    backImgBase = ref(''), //验证码背景图片
    blockBackImgBase = ref(''), //验证滑块的背景图片
    backToken = ref(''), //后端返回的唯一token值
    startMoveTime = ref(), //移动开始的时间
    endMovetime = ref(), //移动结束的时间
    tipWords = ref(''),
    text = ref(''),
    finishText = ref(''),
    moveBlockLeft = ref(undefined),
    leftBarWidth = ref(undefined),
    // 移动中样式
    moveBlockBackgroundColor = ref(undefined),
    leftBarBorderColor = ref('#ddd'),
    iconColor = ref(undefined),
    iconClass = ref('icon-right'),
    status = ref(false), //鼠标状态
    isEnd = ref(false), //是够验证完成
    showRefresh = ref(true),
    transitionLeft = ref(''),
    transitionWidth = ref(''),
    startLeft = ref(0)

const barArea = computed(() => {
    return proxy.$el.querySelector('.verify-bar-area')
})
function init() {
    text.value = explain.value
    getPicture()
    nextTick(() => {})

    window.addEventListener('touchmove', move)
    window.addEventListener('mousemove', move)

    //鼠标松开
    window.addEventListener('touchend', end)
    window.addEventListener('mouseup', end)
}
onUnmounted(() => {
    window.removeEventListener('touchmove', move)
    window.removeEventListener('mousemove', move)

    //鼠标松开
    window.removeEventListener('touchend', end)
    window.removeEventListener('mouseup', end)
})
onMounted(() => {
    // 禁止拖拽
    init()
    proxy.$el.onselectstart = function () {
        return false
    }
})
//鼠标按下
function start(e) {
    e = e || window.event
    let x = 0
    if (!e.touches) {
        //兼容PC端
        x = e.clientX
    } else {
        //兼容移动端
        x = e.touches[0].pageX
    }
    console.log(barArea)
    startLeft.value = Math.floor(x - barArea.value.getBoundingClientRect().left)
    console.log('startLeft', startLeft.value)

    startMoveTime.value = +new Date() //开始滑动的时间
    if (isEnd.value == false) {
        text.value = ''
        moveBlockBackgroundColor.value = '#337ab7'
        leftBarBorderColor.value = '#337AB7'
        iconColor.value = '#fff'
        e.stopPropagation()
        status.value = true
    }
}
//鼠标移动
function move(e) {
    e = e || window.event
    if (status.value && isEnd.value == false) {
        let x
        if (!e.touches) {
            //兼容PC端
            x = e.clientX
        } else {
            //兼容移动端
            x = e.touches[0].pageX
        }
        const bar_area_left = barArea.value.getBoundingClientRect().left
        // console.log('bar_area_left', x, bar_area_left)

        let left = x - bar_area_left - startLeft.value //小方块相对于父元素的left值

        // @ts-ignore

        if (left >= props.imgSize.width - blockSize.value.width) {
            left = props.imgSize.width - blockSize.value.width
        }

        if (left < 0) {
            left = 0
        }
        //拖动后小方块的left值
        moveBlockLeft.value = left + 'px'
        leftBarWidth.value = left + 'px'
        // console.log(moveBlockLeft.value)
    }
}

//鼠标松开
function end() {
    endMovetime.value = +new Date()
    //判断是否重合
    if (status.value && isEnd.value == false) {
        let moveLeftDistance = parseInt((moveBlockLeft.value || '').replace('px', ''))

        moveLeftDistance = (moveLeftDistance * 310) / parseInt(props.imgSize.width)
        const data = {
            captchaType: props.captchaType,
            pointJson: secretKey.value
                ? aesEncrypt(JSON.stringify({ x: moveLeftDistance, y: 5.0 }), secretKey.value)
                : JSON.stringify({ x: moveLeftDistance, y: 5.0 }),
            token: backToken.value
        }
        reqCheck(data).then((res) => {
            if (res && res.repCode == '0000') {
                moveBlockBackgroundColor.value = '#5cb85c'
                leftBarBorderColor.value = '#5cb85c'
                iconColor.value = '#fff'
                iconClass.value = 'icon-check'
                showRefresh.value = false
                isEnd.value = true
                if (mode.value == 'pop') {
                    setTimeout(() => {
                        refresh()
                    }, 1500)
                }
                passFlag.value = true
                tipWords.value = `${((endMovetime.value - startMoveTime.value) / 1000).toFixed(
                    2
                )}s验证成功`
                setTimeout(() => {
                    tipWords.value = ''
                    emit('success', { ...data })
                }, 1000)
            } else {
                moveBlockBackgroundColor.value = '#d9534f'
                leftBarBorderColor.value = '#d9534f'
                iconColor.value = '#fff'
                iconClass.value = 'icon-close'
                passFlag.value = false
                setTimeout(function () {
                    refresh()
                }, 1000)
                emit('error')
                tipWords.value = '验证失败'
                setTimeout(() => {
                    tipWords.value = ''
                }, 1000)
            }
        })
        status.value = false
    }
}

const refresh = () => {
    showRefresh.value = true
    finishText.value = ''
    tipWords.value = ''

    transitionLeft.value = 'left .3s'
    moveBlockLeft.value = 0

    leftBarWidth.value = undefined
    transitionWidth.value = 'width .3s'

    leftBarBorderColor.value = '#ddd'
    moveBlockBackgroundColor.value = '#fff'
    iconColor.value = '#000'
    iconClass.value = 'icon-right'
    isEnd.value = false

    getPicture()
    setTimeout(() => {
        transitionWidth.value = ''
        transitionLeft.value = ''
        text.value = explain.value
    }, 300)
}

// 请求背景图片和验证图片
function getPicture() {
    const data = {
        captchaType: props.captchaType
    }
    reqGet(data).then((res) => {
        if (!res) {
            tipWords.value = '网络错误'
            return
        }
        if (res.repCode == '0000') {
            backImgBase.value = 'data:image/png;base64,' + res.repData.originalImageBase64
            blockBackImgBase.value = 'data:image/png;base64,' + res.repData.jigsawImageBase64
            backToken.value = res.repData.token
            secretKey.value = res.repData.secretKey
        } else {
            tipWords.value = res.repMsg
        }
    })
}
</script>
