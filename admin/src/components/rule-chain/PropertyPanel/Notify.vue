<template>
    <el-form label-width="90px">
        <el-form-item label="服务类型" required>
            <el-radio-group v-model="model.service_type">
                <el-radio value="site">站内消息</el-radio>
                <el-radio value="email">邮件</el-radio>
                <el-radio value="webhook">Webhook</el-radio>
            </el-radio-group>
        </el-form-item>

        <!-- 站内消息：下拉多选管理员 -->
        <el-form-item v-if="model.service_type === 'site'" label="接收人">
            <el-select
                v-model="model.receiver_id"
                placeholder="留空则不指定接收人"
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
            <el-form-item label="收件邮箱">
                <el-input-tag
                    v-model="model.email_to"
                    placeholder="输入邮箱后回车添加；或从下方选择用户追加"
                    clearable
                    trigger="Enter"
                />
            </el-form-item>
            <el-form-item label="选择用户">
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

        <!-- Webhook：回调地址 -->
        <el-form-item v-else-if="model.service_type === 'webhook'" label="回调地址">
            <el-input v-model="model.webhook_url" placeholder="https://example.com/webhook" clearable />
        </el-form-item>

        <el-form-item label="消息内容">
            <el-input
                v-model="model.service_content"
                type="textarea"
                :rows="4"
                placeholder="请输入要发送的内容"
                maxlength="500"
                show-word-limit
            />
            <div class="rc-tip">
                可在内容中使用占位符 <code>${sn}</code> <code>${product}</code> <code>${time}</code>，执行期由后端替换
            </div>
        </el-form-item>
    </el-form>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { adminListAll } from '@/api/perms/admin'
import type { RcNotifyProps } from '../rule.type'

const model = defineModel<RcNotifyProps>({ required: true })

type LabelValue = { label: string; value: any; email: string }
const adminUserList = ref<LabelValue[]>([])

function getAdminList() {
    adminListAll({}).then((res: any) => {
        adminUserList.value = res.map((item: any) => ({
            value: item.id,
            label: item.nickname + ' (' + item.email + ')',
            email: item.email
        }))
    })
}

function appendUserEmail(userId: any) {
    if (!userId) return
    const user = adminUserList.value.find((u) => u.value === userId)
    if (!user || !user.email) return
    const list: string[] = model.value.email_to ?? []
    if (list.includes(user.email)) return
    model.value.email_to = [...list, user.email]
}

onMounted(() => {
    getAdminList()
})
</script>

<style scoped>
.rc-tip {
    font-size: 12px;
    color: var(--el-text-color-secondary);
    line-height: 1.6;
    margin-top: 6px;
}
code {
    background: var(--el-fill-color-light);
    padding: 0 4px;
    border-radius: 3px;
}
</style>
