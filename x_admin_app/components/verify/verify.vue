<template>
  <view :class="props.mode == 'pop' ? 'mask' : ''" v-show="showBox">
    <view
      :class="props.mode == 'pop' ? 'verify-box' : ''"
      :style="{ 'max-width': parseInt(props.imgSize.width) + 30 + 'px' }"
    >
      <view class="verify-box-top" v-if="props.mode == 'pop'">
        请完成安全验证
        <text class="verify-box-close icon-close" @click="clickShow = false">
          
        </text>
      </view>
      <view
        class="verify-box-bottom"
        :style="{ padding: props.mode == 'pop' ? '15px' : '0' }"
      >
        <!-- 验证码容器 -->
        <!-- 滑动 -->
        <view v-if="props.captchaType == 'blockPuzzle'">
          <VerifySlide
            :captchaType="'blockPuzzle'"
            :vSpace="props.vSpace"
            :explain="props.explain"
            :imgSize="props.imgSize"
            :blockSize="props.blockSize"
            ref="VerifySlideInstance"
            @success="success"
            @error="error"
          ></VerifySlide>
        </view>
        <!-- 点选 -->
        <view v-else-if="props.captchaType == 'clickWord'">
          <VerifyPoint
            :captchaType="'clickWord'"
            :vSpace="props.vSpace"
            :explain="props.explain"
            :imgSize="props.imgSize"
            ref="VerifyPointsInstance"
            @success="success"
            @error="error"
          ></VerifyPoint>
        </view>
      </view>
    </view>
  </view>
</template>
<script setup lang="ts">
/**
 * Verify 验证码组件
 * @description 分发验证码使用
 * */

import { ref, computed, onMounted, shallowRef } from "vue";
import type { PropType } from "vue";
import VerifySlide from "./verifySlider/verifySlider.vue";
import VerifyPoint from "./verifyPoint/verifyPoint.vue";

const props = defineProps({
  captchaType: {
    type: String as PropType<"clickWord" | "blockPuzzle">, //clickWord,blockPuzzle
    required: true,
    default: "",
  },

  mode: {
    type: String as PropType<"pop" | "">,
    default: "pop",
  },
  vSpace: {
    type: Number,
    default: 5,
  },
  explain: {
    type: String,
    default: "向右滑动完成验证",
  },
  imgSize: {
    type: Object,
    default: () => ({
      width: 310,
      height: 155,
    }),
  },
  blockSize: {
    type: Object,
    default: () => ({
      width: 50,
      height: 50,
    }),
  },
});

const emit = defineEmits(["success", "error"]);

const clickShow = ref(false);
const showBox = computed(() => {
  return props.mode === "pop" ? clickShow.value : true;
});

// const VerifySlideInstance = ref(null);
// const VerifyPointsInstance = ref(null);
const VerifySlideInstance = shallowRef<InstanceType<typeof VerifySlide>>();
const VerifyPointsInstance = shallowRef<InstanceType<typeof VerifyPoint>>();
const refresh = () => {
  if (props.captchaType === "blockPuzzle") {
    VerifySlideInstance.value.refresh();
  } else {
    VerifyPointsInstance.value.refresh();
  }
};

const show = () => {
  if (props.mode === "pop") {
    clickShow.value = true;
    // refresh();
  }
};
const success = (e) => {
  emit("success", e);
  if (props.mode === "pop") {
    clickShow.value = false;
  }
};

const error = () => {
  emit("error");
};

defineExpose({
  refresh,
  show,
});
onMounted(() => {
  // 组件挂载后执行的逻辑
});
</script>
<style scoped>
.verify-box {
  position: relative;
  box-sizing: border-box;
  border-radius: 2px;
  border: 1px solid #e4e7eb;
  background-color: #fff;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}

.verify-box-top {
  padding: 0 15px;
  height: 50px;
  line-height: 50px;
  text-align: left;
  font-size: 16px;
  color: #45494c;
  border-bottom: 1px solid #e4e7eb;
  box-sizing: border-box;
}

.verify-box-bottom {
  /* padding: 15px; */
  box-sizing: border-box;
}



.mask {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1001;
  width: 100%;
  height: 100vh;
  background: rgba(0, 0, 0, 0.3);
  /* display: none; */
  transition: all 0.5s;
}

.verify-box-close {
  position: absolute;
  top: 13px;
  right: 9px;
  width: 24px;
  height: 24px;
  text-align: center;
  cursor: pointer;
}
.icon-close:before {
  content: " ";
  display: block;
  width: 16px;
  height: 16px;
  position: absolute;
  margin: auto;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  z-index: 9999;
  background-image: url("data:image/png;base64,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");
  background-size: contain;
}

 
</style>
