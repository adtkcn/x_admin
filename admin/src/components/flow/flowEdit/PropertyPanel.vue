<template>
  <el-drawer
    v-model="drawerVisible"
    size="500px"
    :title="node?.text?.value"
    @close="close"
  >
    <div class="setting-block">
      <!-- 开始节点 -->
      {{ node }}

      <div v-if="node.type == 'bpmn:startEvent'">开始节点</div>
      <div v-if="node.type == 'bpmn:userTask'">
        审批节点 设置审批人（具体人员，角色，部门（负责人），岗位？）
      </div>

      <div v-if="node.type == 'bpmn:serviceTask'">系统任务</div>
      <div v-if="node.type == 'bpmn:exclusiveGateway'">
        <div>网关</div>
        <div>从form取值判断</div>
      </div>

      <div v-if="node.type == 'bpmn:endEvent'">结束</div>

      <!-- 网关和结束节点不需要表单权限设置 -->
      <el-table
        v-if="
          ['bpmn:startEvent', 'bpmn:userTask', 'bpmn:serviceTask'].includes(
            node.type
          )
        "
        fit
        size="small"
        :data="fieldList"
        style="width: 100%"
      >
        <el-table-column prop="label" label="表单"></el-table-column>
        <el-table-column label="权限">
          <template #default="{ row }">
            <el-radio-group v-model="row.auth">
              <el-radio :label="1">读写</el-radio>
              <el-radio :label="2">可读</el-radio>
              <el-radio :label="3">隐藏</el-radio>
            </el-radio-group>
          </template>
        </el-table-column>
      </el-table>
      <!-- {{ fieldList }} -->
      <!-- 用户审批节点 -->

      <!-- 系统任务节点 -->
    </div>
  </el-drawer>
</template>

<script>
export default {
  name: "PropertyPanel",
  props: {},
  data() {
    return {
      drawerVisible: false,

      node: {},
      /**
       * 表单列表
       * [{
       *  id: 1,
       *  label: '表单1',
       *  auth: 1,
       * }]
       */
      fieldList: [],
    };
  },

  methods: {
    open(node, fieldList) {
      this.node = node;

      this.fieldList = fieldList.map((item) => {
        let auth = 1;
        let formId = item?.field?.id;
        if (node?.properties?.fieldAuth?.[formId]) {
          auth = node.properties.fieldAuth[formId];
        }
        return {
          id: formId,
          label: item?.field?.options?.label,
          auth: auth,
        };
      });

      this.drawerVisible = true;
    },
    close() {
      var fieldAuth = {};
      this.fieldList.forEach((item) => {
        fieldAuth[item.id] = item.auth;
      });
      this.setProperties("fieldAuth", {
        ...fieldAuth,
      });
    },
    setProperties(key, val) {
      this.$emit("setProperties", this.node, {
        [key]: val,
      });
    },
  },
  components: {},
};
</script>

<style scoped></style>
