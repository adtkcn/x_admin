<template>
  <uv-input
    :modelValue="selectItem?.[props.labelKey]"
    placeholder="请选择"
    readonly
    @click="openPicker"
  >
  </uv-input>

  <uv-picker
    ref="picker"
    :columns="columns"
    :keyName="props.labelKey"
    :defaultIndex="pickerIndex"
    @confirm="handleConfirm"
  ></uv-picker>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
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
const columns = computed(() => {
  return [props.columns];
});
const picker = ref(null);

const model = defineModel('modelValue');
const pickerIndex = ref([0]);

const selectItem = ref({});

function openPicker() {
  picker.value.open();
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
  console.log("updateSelectItem", model.value, props.columns);
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
  console.log("onMounted");
  updateSelectItem();
});
watch(
  () => [model.value, props.columns],
  (newVal) => {
    console.log("watch", newVal);
    updateSelectItem();
  }
);
</script>

<style lang="scss" scoped></style>
