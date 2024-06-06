<template>
  <view style="position: relative">
    <view class="verify-image-out" v-show="showImage">
      <view
        class="verify-image-panel"
        :style="{
          width: props.imgSize.width + 'px',
          height: props.imgSize.height + 'px',
          'margin-bottom': vSpace + 'px',
        }"
      >
        <view
          class="verify-refresh icon-refresh"
          style="z-index: 3"
          @click="refresh"
          v-show="showRefresh"
        >
        </view>
        <image
          :src="pointBackImgBase"
          id="image"
          ref="canvas"
          style="width: 100%; height: 100%; display: block"
          @click="bindingClick ? canvasClick($event) : undefined"
        ></image>
        <view
          v-for="(tempPoint, index) in tempPoints"
          :key="index"
          class="point-area"
          :style="{
            'background-color': '#1abd6c',
            color: '#fff',
            'z-index': 9999,
            width: '20px',
            height: '20px',
            'text-align': 'center',
            'line-height': '20px',
            'border-radius': '50%',
            position: 'absolute',
            top: tempPoint.y - 10 + 'px',
            left: tempPoint.x - 10 + 'px',
          }"
        >
          {{ index + 1 }}
        </view>
      </view>
    </view>

    <view
      class="verify-bar-area"
      :style="{
        width: props.imgSize.width + 'px',
        color: barAreaColor,
        'border-color': barAreaBorderColor,
        'line-height': '40px',
      }"
    >
      <text class="verify-msg">{{ text }}</text>
    </view>
  </view>
</template>
<script setup lang="ts">
import { ref, reactive, nextTick, onMounted, getCurrentInstance } from "vue";
/**
 * VerifyPoints
 * @description 点选
 * */
import { aesEncrypt } from "./../utils/ase.js";
import { myRequest } from "../utils/request.js";
defineOptions({ name: "VerifyPoints" });
let emit = defineEmits(["success", "error"]);
const props = defineProps({
  captchaType: String,
  vSpace: {
    type: Number,
    default: 5,
  },
  imgSize: {
    type: Object,
    default: () => ({
      width: 310,
      height: 155,
    }),
  },
});
const secretKey = ref(""); // 后端返回的加密秘钥 字段
const checkNum = ref(3); //
const fontPos = ref([]); // 选中的坐标信息
const checkPosArr = ref([]); // 用户点击的坐标
const num = ref(1); // 点击的记数
const pointBackImgBase = ref(""); // 后端获取到的背景图片
const pointTextList = ref([]); // 后端返回的点击字体顺序
const backToken = ref(""); // 后端返回的token值

const showImage = ref(true);
const tempPoints = ref([]);
const text = ref("");
const barAreaColor = ref("#fff");
const barAreaBorderColor = ref("#fff");
const showRefresh = ref(true);
const bindingClick = ref(true);
const imgLeft = ref(0);
const imgTop = ref(0);

function init() {
  // 加载页面
  fontPos.value.splice(0, fontPos.value.length);
  checkPosArr.value.splice(0, checkPosArr.value.length);
  num.value = 1;
  nextTick(() => {
    refresh();
  });
}

let canvas = ref(null);
const appInstance = getCurrentInstance().proxy;
function canvasClick(e) {
  const query = uni.createSelectorQuery().in(appInstance);
  query
    .select("#image")
    .boundingClientRect((rect) => {
      // @ts-ignore
      imgLeft.value = Math.ceil(rect.left);
      // @ts-ignore
      imgTop.value = Math.ceil(rect.top);
      checkPosArr.value.push(getMousePos(canvas.value, e));
      if (num.value == checkNum.value) {
        num.value = createPoint(getMousePos(canvas.value, e));
        //按比例转换坐标值
        checkPosArr.value = pointTransform(checkPosArr.value, props.imgSize);
        //等创建坐标执行完
        setTimeout(() => {
          //发送后端请求
          let info = {
            captchaType: props.captchaType,
            pointJson: secretKey.value
              ? aesEncrypt(JSON.stringify(checkPosArr.value), secretKey.value)
              : JSON.stringify(checkPosArr.value),
            token: backToken.value,
          };
          myRequest({
            url: `/captcha/check`,
            data: info,
            method: "POST",
          }).then((result) => {
            let res = result.data;
            if (res.repCode == "0000") {
              barAreaColor.value = "#4cae4c";
              barAreaBorderColor.value = "#5cb85c";
              text.value = "验证成功";
              bindingClick.value = false;
              setTimeout(() => {
                refresh();
              }, 1500);
              emit("success", {
                ...info,
              });
            } else {
              emit("error");
              barAreaColor.value = "#d9534f";
              barAreaBorderColor.value = "#d9534f";
              text.value = "验证失败";
              setTimeout(() => {
                refresh();
              }, 700);
            }
          });
        }, 400);
      }
      if (num.value < checkNum.value) {
        num.value = createPoint(getMousePos(canvas.value, e));
      }
    })
    .exec();
}

function getMousePos(obj, e) {
  let position = {
    x: Math.ceil(e.detail.x) - imgLeft.value,
    y: Math.ceil(e.detail.y) - imgTop.value,
  };
  return position;
}

function createPoint(pos) {
  tempPoints.value.push(Object.assign({}, pos));
  return ++num.value;
}

function refresh() {
  tempPoints.value.splice(0, tempPoints.value.length);
  barAreaColor.value = "#000";
  barAreaBorderColor.value = "#ddd";
  bindingClick.value = true;

  fontPos.value.splice(0, fontPos.value.length);
  checkPosArr.value.splice(0, checkPosArr.value.length);
  num.value = 1;

  getPicture();

  // text.value = '验证失败'
  showRefresh.value = true;
}

function getPicture() {
  let info = {
    captchaType: props.captchaType,
    ts: Date.now(), // 现在的时间戳
  };
  myRequest({
    url: "/captcha/get", //仅为示例，并非真实接口地址。
    data: info,
    method: "POST",
  }).then((result) => {
    let res = result.data;
    if (res.repCode == "0000") {
      pointBackImgBase.value =
        "data:image/png;base64," + res.repData.originalImageBase64;
      backToken.value = res.repData.token;
      secretKey.value = res.repData.secretKey;
      pointTextList.value = res.repData.wordList;
      text.value = "请依次点击【" + pointTextList.value.join(",") + "】";
    }
    // 判断接口请求次数是否失效
    if (res.repCode == "6201") {
      pointBackImgBase.value = null;
    }
  });
}
function pointTransform(pointArr, imgSize) {
  var newPointArr = pointArr.map((p) => {
    let x = Math.round((310 * p.x) / parseInt(imgSize.width));
    let y = Math.round((155 * p.y) / parseInt(imgSize.height));
    return {
      x,
      y,
    };
  });
  // console.log(newPointArr,"newPointArr");
  return newPointArr;
}
defineExpose({
  refresh,
});
onMounted(init);
</script>

<style scoped>
/*滑动验证码*/
.verify-bar-area {
  position: relative;
  background: #ffffff;
  text-align: center;
  -webkit-box-sizing: content-box;
  -moz-box-sizing: content-box;
  box-sizing: content-box;
  border: 1px solid #ddd;
}

.verify-bar-area .verify-move-block {
  position: absolute;
  top: 0px;
  left: 0;
  background: #fff;
  cursor: pointer;
  -webkit-box-sizing: content-box;
  -moz-box-sizing: content-box;
  box-sizing: content-box;
  box-shadow: 0 0 2px #888888;
}

.verify-bar-area .verify-move-block:hover {
  background-color: #337ab7;
  color: #ffffff;
}

.verify-bar-area .verify-left-bar {
  position: absolute;
  top: -1px;
  left: -1px;
  background: #f0fff0;
  cursor: pointer;
  -webkit-box-sizing: content-box;
  -moz-box-sizing: content-box;
  box-sizing: content-box;
  border: 1px solid #ddd;
}

.verify-image-panel {
  margin: 0;
  -webkit-box-sizing: content-box;
  -moz-box-sizing: content-box;
  box-sizing: content-box;
  border-top: 1px solid #ddd;
  border-bottom: 1px solid #ddd;
  border-radius: 3px;
  position: relative;
}

.verify-image-panel .verify-refresh {
  width: 25px;
  height: 25px;
  text-align: center;
  padding: 5px;
  cursor: pointer;
  position: absolute;
  top: 0;
  right: 0;
  z-index: 2;
}

.verify-image-panel .icon-refresh {
  font-size: 20px;
  color: #fff;
}

.verify-image-panel .verify-gap {
  background-color: #fff;
  position: relative;
  z-index: 2;
  border: 1px solid #fff;
}

.verify-bar-area .verify-move-block .verify-sub-block {
  position: absolute;
  text-align: center;
  z-index: 3;
  /* border: 1px solid #fff; */
}

.verify-bar-area .verify-move-block .verify-icon {
  font-size: 18px;
}

.verify-bar-area .verify-msg {
  z-index: 3;
}

.icon-refresh:before {
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
