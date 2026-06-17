/**
 * S3 客户端（纯 S3 标准协议，不做任何非标准注入）
 *
 * 后端实现 S3 REST API（文件存本地），前端用此客户端直接调用。
 * 未来切换到 MinIO 只需修改 endpoint/credentials，前端零改动。
 */
import { S3Client } from "@aws-sdk/client-s3"

function getToken(): string {
    return localStorage.getItem("token") || ""
}

const s3Client = new S3Client({
    region: "us-east-1",
    endpoint: `${location.protocol}//${location.host}/api/admin/s3`,
    forcePathStyle: true,
    credentials: { accessKeyId: "dummy", secretAccessKey: "dummy" },
})

// 注入 token，覆盖 AWS SigV4 签名（干净：不改请求体/不改 header 业务语义）
s3Client.middlewareStack.add(
    (next: any) => async (args: any) => {
        const req = args.request as any
        const token = getToken()
        if (token) req.headers["token"] = token
        delete req.headers["authorization"]
        delete req.headers["x-amz-date"]
        if (req.headers["x-amz-content-sha256"] === "UNSIGNED-PAYLOAD") {
            delete req.headers["x-amz-content-sha256"]
        }
        return next(args)
    },
    { step: "finalizeRequest", name: "injectToken", priority: "high" },
)

export default s3Client
