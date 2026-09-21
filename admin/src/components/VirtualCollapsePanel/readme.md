# 使用示例

```vue
<VirtualCollapsePanel
  ref="scrollbarRef"
  :data="filterMarkers"
  key-field="id"
  name-field="name"
  :min-item-size="40"
>
  <template #header="{ item }">
  </template>
  <template #content="{ item }"> </template>
</VirtualCollapsePanel>
```
