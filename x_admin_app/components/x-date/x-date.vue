<template>
  <view style="flex: 1" @click="calendarsRef.open()">
    <uv-input :modelValue="title" :readonly="true" placeholder="请选择时间">
    </uv-input>
  </view>
  <uv-calendars ref="calendarsRef" :date="props.time" @confirm="Confirm" />
</template>

<script setup>
import { ref, computed } from "vue";

const emit = defineEmits(["update:time"]);
const props = defineProps({
  time: {
    type: String,
    default: "",
  },
});
let calendarsRef = ref(null);
function Confirm(e) {
  console.log(e);
  emit("update:time", e.fulldate);
  // emit("update:endTime", e.range.after);
}

const title = computed(() => {
  return props.time;
});
</script>

<style lang="scss" scoped></style>
