<template>
  <uv-input
    :modelValue="title"
    :readonly="true"
    placeholder="请选择时间"
    @click="calendarsRef.open()"
  >
  </uv-input>
  <uv-calendars
    ref="calendarsRef"
    mode="range"
    :date="[props.startTime, props.endTime]"
    @confirm="Confirm"
  />
</template>

<script setup>
import { ref, computed } from "vue";

const emit = defineEmits(["update:startTime", "update:endTime"]);
const props = defineProps({
  startTime: {
    type: String,
    default: "",
  },
  endTime: {
    type: String,
    default: "",
  },
});
let calendarsRef = ref(null);
function Confirm(e) {
  console.log(e);
  emit("update:startTime", e.range.before);
  emit("update:endTime", e.range.after);
}

const title = computed(() => {
  if (!props.startTime || !props.endTime) {
    return "";
  }
  return props.startTime + "-" + props.endTime;
});
</script>

<style lang="scss" scoped></style>
