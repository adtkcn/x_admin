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

<script setup>
import { ref, onMounted, computed, watch } from "vue";

const emit = defineEmits(["update:modelValue"]);

const props = defineProps({
  modelValue: {
    type: [String, Number],
    default: null,
  },
  columns: {
    type: Array,
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
const model = computed({
  get() {
    return props.modelValue;
  },
  set(value) {
    emit("update:modelValue", value);
  },
});

const columns = computed(() => {
  return [props.columns];
});
const pickerRef = ref(null);

// const model = defineModel('modelValue');
const pickerIndex = ref([0]);

const selectItem = ref({});

function openPicker() {

  pickerRef.value.open();
}
function handleConfirm(e) {
  // debugger;
  if (e.value[0] !== undefined) {
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
  if (props.columns.length == 0) {
    pickerIndex.value = [0];
    selectItem.value = {};
    return;
  }
  let find = false;
  for (let index = 0; index < props.columns.length; index++) {
    const item = props.columns[index];
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
