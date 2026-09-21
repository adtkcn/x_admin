<template>
    <div class="notify-task">
        <el-form label-width="90px">
            <el-form-item label="服务类型" required>
                <el-radio-group v-model="model.service_type">
                    <el-radio value="site">站内消息</el-radio>
                    <el-radio value="email">邮件</el-radio>
                    <el-radio value="webhook">Webhook</el-radio>
                </el-radio-group>
            </el-form-item>

            <!-- 站内消息：下拉多选用户 -->
            <el-form-item v-if="model.service_type === 'site'" label="接收人" prop="receiver_id">
                <el-select
                    v-model="model.receiver_id"
                    placeholder="留空则默认通知申请人"
                    style="width: 100%"
                    multiple
                    clearable
                    filterable
                    collapse-tags
                    collapse-tags-tooltip
                >
                    <el-option
                        v-for="item in adminUserList"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value"
                    />
                </el-select>
            </el-form-item>

            <!-- 邮件：标签输入邮箱 + 从用户下拉追加其邮箱 -->
            <template v-else-if="model.service_type === 'email'">
                <el-form-item label="收件邮箱" prop="email_to">
                    <el-input-tag
                        v-model="model.email_to"
                        placeholder="输入邮箱后回车添加，可粘贴多个；或从下方选择用户追加"
                        clearable
                        trigger="Enter"
                    />
                </el-form-item>
                <el-form-item label="选择用户" prop="email_user">
                    <el-select
                        :model-value="undefined"
                        placeholder="选择用户，自动将其邮箱追加到收件邮箱"
                        style="width: 100%"
                        clearable
                        filterable
                        @change="appendUserEmail"
                    >
                        <el-option
                            v-for="item in adminUserList"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </el-select>
                </el-form-item>
            </template>

            <!-- Webhook：输入回调地址 -->
            <el-form-item v-else-if="model.service_type === 'webhook'" label="回调地址" prop="webhook_url">
                <el-input
                    v-model="model.webhook_url"
                    placeholder="https://example.com/webhook"
                    clearable
                />
            </el-form-item>

            <el-form-item label="消息内容" prop="service_content">
                <el-input
                    ref="contentInputRef"
                    v-model="model.service_content"
                    type="textarea"
                    :rows="4"
                    placeholder="请输入要发送的内容，可点击下方变量插入占位符"
                    maxlength="500"
                    show-word-limit
                />
                <div class="var-tags">
                    <span class="var-tags__label">插入变量：</span>
                    <el-tag
                        v-for="v in contentVars"
                        :key="v.key"
                        class="var-tags__item"
                        type="primary"
                        effect="plain"
                        disable-transitions
                        @click="insertVar(v.key)"
                    >
                        {{ v.label }}
                    </el-tag>
                </div>
            </el-form-item>
        </el-form>
    </div>
</template>
<script setup lang="ts">
// defineModel 直接暴露父层 v-model="nodeProps" 绑定的 notify_task 私有属性，可读写。
// 父层 defaultProps 已保证字段存在，无需再做 watch 双向同步
import { ref, onMounted } from 'vue'
import { adminListAll } from '@/api/perms/admin'
import type { NotifyTaskProps } from './property.type'

const model = defineModel<NotifyTaskProps>({ required: true })

type LabelValue = {
    label: string
    value: any
    email: string
}
const adminUserList = ref<LabelValue[]>([])

function getAdminList() {
    adminListAll({}).then((res) => {
        adminUserList.value = res.map((item) => {
            return {
                value: item.id,
                label: item.nickname + ' (' + item.email + ')',
                email: item.email
            }
        })
    })
}
// 选择用户时将其邮箱追加进 email_to 数组（去重）
function appendUserEmail(userId: any) {
    if (!userId) return
    const user = adminUserList.value.find((u) => u.value === userId)
    if (!user || !user.email) return
    const list: string[] = model.value.email_to ?? []
    if (list.includes(user.email)) return
    model.value.email_to = [...list, user.email]
}

// 消息内容变量占位符：编辑期插入，执行期由后端替换
const contentVars = [
    { key: '${apply_user}', label: '申请人' },
    { key: '${flow_name}', label: '流程名称' },
    { key: '${apply_id}', label: '申请单号' },
    { key: '${apply_time}', label: '申请时间' }
]
const contentInputRef = ref<any>(null)
// 在文本框光标处插入变量占位符
function insertVar(key: string) {
    const input = contentInputRef.value
    const el = input?.textarea ?? input?.$el?.querySelector('textarea')
    const cur = model.value.service_content ?? ''
    let start = cur.length
    let end = cur.length
    if (el && typeof el.selectionStart === 'number') {
        start = el.selectionStart
        end = el.selectionEnd
    }
    model.value.service_content = cur.slice(0, start) + key + cur.slice(end)
    // 还原光标位置到插入内容之后
    const pos = start + key.length
    if (el) {
        requestAnimationFrame(() => {
            el.focus()
            el.setSelectionRange(pos, pos)
        })
    }
}
onMounted(() => {
    getAdminList()
})
</script>
<style scoped>
.var-tags {
    margin-top: 8px;
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 6px;
}
.var-tags__label {
    font-size: 12px;
    color: var(--el-text-color-secondary);
}
.var-tags__item {
    cursor: pointer;
}
</style>
