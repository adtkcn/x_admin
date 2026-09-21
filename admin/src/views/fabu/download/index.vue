<template>
    <div class="fabu-download">
        <div class="download-card">
            <!-- 加载中 -->
            <div v-if="loading" class="state">
                <el-icon class="is-loading" :size="28"><Loading /></el-icon>
                <span>正在加载应用信息…</span>
            </div>

            <!-- 加载失败 -->
            <div v-else-if="errMsg" class="state">
                <el-icon :size="28" color="#f56c6c"><WarningFilled /></el-icon>
                <span>{{ errMsg }}</span>
            </div>

            <!-- 应用信息 -->
            <template v-else-if="info">
                <el-image class="app-icon" :src="info.icon" fit="cover">
                    <template #error>
                        <div class="app-icon app-icon--empty">
                            <el-icon :size="34"><Picture /></el-icon>
                        </div>
                    </template>
                </el-image>
                <div class="app-name">{{ info.name }}</div>
                <div class="app-meta">
                    <el-tag :type="info.platform === 'ios' ? 'warning' : 'success'" size="small">
                        {{ info.platform === 'ios' ? 'iOS' : 'Android' }}
                    </el-tag>
                    <span class="dot">·</span>
                    <span>版本 {{ info.version }}（{{ info.version_code }}）</span>
                </div>
                <div v-if="info.size" class="app-sub">
                    包大小 {{ formatSize(info.size) }} · 已下载 {{ info.download_times }} 次
                </div>
                <div v-else class="app-sub">已下载 {{ info.download_times }} 次</div>

                <el-button
                    class="download-btn"
                    type="primary"
                    size="large"
                    :disabled="!info.has_version"
                    @click="handleDownload"
                >
                    {{ info.has_version ? '立即下载' : '暂无可用版本' }}
                </el-button>
                <div class="tips">
                    {{
                        info.platform === 'ios'
                            ? '点击后按提示安装，需在 iPhone/iPad 上访问'
                            : '点击下载安装包并安装'
                    }}
                </div>
            </template>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { Loading, WarningFilled, Picture } from '@element-plus/icons-vue'
import { fabuAppDownloadInfo, type type_fabu_download_resp } from '@/api/fabu_download'
import { formatSize } from '@/utils/file'

defineOptions({ name: 'fabuDownload' })
const route = useRoute()
const loading = ref(true)
const errMsg = ref('')
const info = ref<type_fabu_download_resp | null>(null)

const getInfo = async () => {
    const shortUrl = String(route.params.short_url || '')
    if (!shortUrl) {
        loading.value = false
        errMsg.value = '缺少应用短链'
        return
    }
    loading.value = true
    try {
        info.value = await fabuAppDownloadInfo(shortUrl)
    } catch (e: any) {
        errMsg.value = e?.message || '应用不存在或已下线'
    } finally {
        loading.value = false
    }
}

const handleDownload = () => {
    const data = info.value
    if (!data || !data.has_version) return
    const origin = window.location.origin
    // iOS 走 itms-services 安装协议（指向 plist）；Android 走计数下载直装包
    const url =
        data.platform === 'ios'
            ? `itms-services://?action=download-manifest&url=${encodeURIComponent(
                  origin + data.install_url
              )}`
            : `${origin}/api/web/fabu/count/${data.app_id}/${data.version_id}`
    window.location.href = url
}

onMounted(getInfo)
</script>
<style lang="scss" scoped>
.fabu-download {
    min-height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 24px;
    background: linear-gradient(160deg, #f5f7fa 0%, #e9eef5 100%);
}
.download-card {
    width: 100%;
    max-width: 360px;
    padding: 36px 28px;
    background: #fff;
    border-radius: 16px;
    box-shadow: 0 12px 40px rgba(0, 0, 0, 0.08);
    text-align: center;
}
.state {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 12px;
    color: #909399;
    font-size: 14px;
    padding: 30px 0;
}
.app-icon {
    width: 88px;
    height: 88px;
    border-radius: 20px;
    box-shadow: 0 6px 18px rgba(0, 0, 0, 0.12);
}
.app-icon--empty {
    display: flex;
    align-items: center;
    justify-content: center;
    background: #f2f3f5;
    color: #c0c4cc;
}
.app-name {
    margin-top: 18px;
    font-size: 20px;
    font-weight: 600;
    color: #303133;
    word-break: break-all;
}
.app-meta {
    margin-top: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 6px;
    font-size: 13px;
    color: #606266;
}
.dot {
    color: #c0c4cc;
}
.app-sub {
    margin-top: 8px;
    font-size: 12px;
    color: #909399;
}
.download-btn {
    width: 100%;
    margin-top: 26px;
    border-radius: 10px;
    font-size: 16px;
}
.tips {
    margin-top: 14px;
    font-size: 12px;
    color: #a8abb2;
}
</style>
