<template>
    <div class="setting-container">
        <el-form
            ref="elForm"
            :model="formData"
            :rules="rules"
            label-width="100px"
            label-position="top"
        >
            <el-form-item label="审批名称" prop="flowName">
                <el-input
                    v-model="formData.flowName"
                    placeholder="请输入审批名称"
                    clearable
                    :style="{ width: '100%' }"
                >
                </el-input>
            </el-form-item>
            <el-form-item label="选择分组" prop="flowGroup">
                <el-select
                    v-model="formData.flowGroup"
                    placeholder="请选择选择分组"
                    clearable
                    :style="{ width: '100%' }"
                >
                    <el-option
                        v-for="(item, index) in flowGroupOptions"
                        :key="index"
                        :label="item.label"
                        :value="item.value"
                        :disabled="item.disabled"
                    ></el-option>
                </el-select>
            </el-form-item>
            <!-- <el-form-item label="谁可以发起审批" prop="approver">
        <span style="font-size: 12px; color: #aaa">默认所有人</span>
      </el-form-item> -->

            <el-form-item label="审批说明" prop="flowRemark">
                <el-input
                    v-model="formData.flowRemark"
                    type="textarea"
                    placeholder="请输入审批说明"
                    :maxlength="100"
                    show-word-limit
                    :autosize="{ minRows: 4, maxRows: 4 }"
                    :style="{ width: '100%' }"
                ></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
export default {
    name: 'BasicSetting',
    components: {},
    props: ['tabName', 'conf'],
    data() {
        return {
            formData: {
                flowName: '',
                flowImg: '',
                flowGroup: undefined,
                flowRemark: undefined
            },
            rules: {
                flowName: [
                    {
                        required: true,
                        message: '请输入审批名称',
                        trigger: 'blur'
                    }
                ],
                flowGroup: [
                    {
                        required: true,
                        message: '请选择选择分组',
                        trigger: 'change'
                    }
                ]
            },
            // iconList,
            flowGroupOptions: [
                {
                    label: '假勤管理',
                    value: 1
                },
                {
                    label: '人事管理',
                    value: 2
                },
                {
                    label: '财务管理',
                    value: 3
                },
                {
                    label: '业务管理',
                    value: 4
                },
                {
                    label: '行政管理',
                    value: 5
                },
                {
                    label: '法务管理',
                    value: 6
                },
                {
                    label: '其他',
                    value: 7
                }
            ]
        }
    },
    computed: {
        // activeIconSrc() {
        //   const icon = this.iconList.find((t) => t.id === this.activeIcon);
        //   return icon ? icon.src : "";
        // },
    },
    created() {
        if (typeof this.conf === 'object' && this.conf !== null) {
            Object.assign(this.formData, this.conf)
        }
    },
    methods: {
        // 给父级页面提供得获取本页数据得方法
        getData() {
            return new Promise((resolve, reject) => {
                this.$refs['elForm'].validate((valid) => {
                    if (!valid) {
                        reject({ target: this.tabName })
                        return
                    }
                    // this.formData.flowImg = this.activeIcon
                    resolve({ formData: this.formData, target: this.tabName }) // TODO 提交表单
                })
            })
        }
    }
}
</script>

<style lang="scss" scoped>
.icon-item {
    width: 40px;
    height: 40px;
    margin: 6px;
    position: relative;
    cursor: pointer;
    opacity: 0.5;
    &.active {
        opacity: 1;
    }
    &:hover {
        opacity: 1;
    }
}
.setting-container {
    width: 600px;
    margin: auto;
    background: #fff;
    padding: 16px;
}
</style>
