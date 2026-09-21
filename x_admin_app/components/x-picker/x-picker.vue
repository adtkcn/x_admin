<template>
  <!-- 在微信小程序readonly后无法click,所以套一层view -->
  <view style="flex: 1" @click="openPicker">
    <uv-input
      :modelValue="selectItem?.[props.labelKey]"
      placeholder="请选择"
      readonly
    >
    </uv-input>
  </view>

  <uv-picker
    ref="pickerRef"
    :columns="columns"
    :keyName="props.labelKey"
    :defaultIndex="pickerIndex"
    @confirm="handleConfirm"
  ></uv-picker>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, watch, type PropType } from "vue";

const emit = defineEmits(["update:modelValue"]);

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: null,
  },
  columns: {
    type: Array as PropType<any[]>,
    default: () => [],
  },
  labelKey: {
    type: String,
    default: "label",
  },
  valueKey: {
    type: String,
    default: "id",
  },
});
const model = computed<any>({
  get() {
    return props.modelValue;
  },
  set(value: any) {
    emit("update:modelValue", value);
  },
});

// 归一 columns：调用方可能传入 null（如字典无配置时接口返回 null），
// 而 prop 默认值仅在 undefined 时生效，故这里兜底为空数组，避免 .length 崩溃
const safeColumns = computed<any[]>(() => props.columns ?? []);
const columns = computed(() => {
  return [safeColumns.value];
});
const pickerRef = ref<any>(null);

// const model = defineModel('modelValue');
const pickerIndex = ref<number[]>([0]);

const selectItem = ref<any>({});

function openPicker() {

  pickerRef.value.open();
}
function handleConfirm(e: any) {
  // debugger;
  if (e.value[0] != null) {
    model.value = e.value[0][props.valueKey];
    selectItem.value = e.value[0];
  } else {
    model.value = null;
    selectItem.value = {};
    console.log("handleConfirm没有数据", e);
  }
}
function updateSelectItem() {
 
  if (!model.value) {
    pickerIndex.value = [0];
    selectItem.value = {};
    return;
  }
  if (safeColumns.value.length == 0) {
    pickerIndex.value = [0];
    selectItem.value = {};
    return;
  }
  let find = false;
  for (let index = 0; index < safeColumns.value.length; index++) {
    const item = safeColumns.value[index];
    if (model.value == item[props.valueKey]) {
      selectItem.value = item;
      pickerIndex.value = [index];
      find = true;
      break;
    }
  }
  if (!find) {
    selectItem.value = {};
    pickerIndex.value = [0];
    model.value = null;
  }
}
onMounted(() => {
 
  updateSelectItem();
});
watch(
  () => [model.value, props.columns],
  (newVal) => {
 
    updateSelectItem();
  }
);
</script>

<style lang="scss" scoped></style>
