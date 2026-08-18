package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"x_admin/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// ossStorageEngine 阿里云 OSS 存储引擎（基于 S3 兼容协议，使用 aws-sdk-go-v2，免手写 HMAC 签名）
type ossStorageEngine struct {
	client *s3.Client
	bucket string
}

// NewOssStorage 根据配置返回存储引擎：配置了 OSS 则使用 OSS，否则回退本地存储
func NewOssStorage() StorageEngine {
	oss := config.OssConfig
	if oss.Bucket != "" && oss.AccessKeyId != "" && oss.AccessKeySecret != "" && oss.Endpoint != "" {
		cfg, err := awsConfig(oss.AccessKeyId, oss.AccessKeySecret, oss.Region)
		if err != nil {
			// 配置加载失败则回退本地存储，避免进程启动崩溃
			return newLocalStorageEngine()
		}
		client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.BaseEndpoint = aws.String("https://" + strings.TrimSuffix(oss.Endpoint, "/"))
			// OSS 采用虚拟主机风格寻址 https://{bucket}.{endpoint}，与 SDK 默认一致
			o.UsePathStyle = false
		})
		return &ossStorageEngine{client: client, bucket: oss.Bucket}
	}
	return newLocalStorageEngine()
}

// awsConfig 构造带 AK/SK 与 region 的 AWS 配置，指向 OSS 端点由 BaseEndpoint 决定
func awsConfig(accessKeyId, accessKeySecret, region string) (aws.Config, error) {
	if region == "" {
		region = "oss-cn-hangzhou" // OSS S3 兼容接口对 region 不敏感，提供默认值
	}
	return awsconfig.LoadDefaultConfig(context.Background(),
		awsconfig.WithRegion(region),
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, ""),
		),
	)
}

// withTimeout 统一为外部存储调用附加超时，避免挂起阻塞请求
func withTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Minute)
}

// PutObject 上传完整对象
func (e *ossStorageEngine) PutObject(key string, data io.Reader, size int64) (string, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	_, err := e.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(key),
		Body:   data,
	})
	if err != nil {
		return "", fmt.Errorf("上传失败: %w", err)
	}
	return key, nil
}

// GetObjectURL 获取对象访问 URL
func (e *ossStorageEngine) GetObjectURL(key string) (string, error) {
	oss := config.OssConfig
	if oss.PublicEndpoint != "" {
		return strings.TrimSuffix(oss.PublicEndpoint, "/") + "/" + key, nil
	}
	return fmt.Sprintf("https://%s.%s/%s", e.bucket, strings.TrimSuffix(oss.Endpoint, "/"), key), nil
}

// ObjectExists 通过 HEAD 判断对象是否存在（用于秒传）
func (e *ossStorageEngine) ObjectExists(key string) (bool, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	_, err := e.client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if _, ok := errors.AsType[*types.NoSuchKey](err); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// InitMultipartUpload 初始化分片上传，返回 uploadId
func (e *ossStorageEngine) InitMultipartUpload(key string) (string, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	out, err := e.client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
		Bucket: aws.String(e.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("初始化分片上传失败: %w", err)
	}
	return aws.ToString(out.UploadId), nil
}

// UploadPart 上传单个分片，返回 etag
func (e *ossStorageEngine) UploadPart(key string, uploadId string, partNumber int, data io.Reader, size int64) (string, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	out, err := e.client.UploadPart(ctx, &s3.UploadPartInput{
		Bucket:     aws.String(e.bucket),
		Key:        aws.String(key),
		UploadId:   aws.String(uploadId),
		PartNumber: aws.Int32(int32(partNumber)),
		Body:       data,
	})
	if err != nil {
		return "", fmt.Errorf("上传分片失败: %w", err)
	}
	return aws.ToString(out.ETag), nil
}

// CompleteMultipartUpload 完成分片合并
func (e *ossStorageEngine) CompleteMultipartUpload(key string, uploadId string, parts []MultipartPart) (string, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	completed := make([]types.CompletedPart, 0, len(parts))
	for _, p := range parts {
		completed = append(completed, types.CompletedPart{
			ETag:       aws.String(p.ETag),
			PartNumber: aws.Int32(int32(p.PartNumber)),
		})
	}
	_, err := e.client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
		Bucket:          aws.String(e.bucket),
		Key:             aws.String(key),
		UploadId:        aws.String(uploadId),
		MultipartUpload: &types.CompletedMultipartUpload{Parts: completed},
	})
	if err != nil {
		return "", fmt.Errorf("完成分片上传失败: %w", err)
	}
	return key, nil
}

// AbortMultipartUpload 中止分片上传，清理服务端分片数据
func (e *ossStorageEngine) AbortMultipartUpload(key string, uploadId string) error {
	ctx, cancel := withTimeout()
	defer cancel()
	_, err := e.client.AbortMultipartUpload(ctx, &s3.AbortMultipartUploadInput{
		Bucket:   aws.String(e.bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
	})
	if err != nil {
		return fmt.Errorf("中止分片上传失败: %w", err)
	}
	return nil
}

// ListParts 列出已上传的分片（用于断点续传）
func (e *ossStorageEngine) ListParts(key string, uploadId string) ([]MultipartPart, error) {
	ctx, cancel := withTimeout()
	defer cancel()
	out, err := e.client.ListParts(ctx, &s3.ListPartsInput{
		Bucket:   aws.String(e.bucket),
		Key:      aws.String(key),
		UploadId: aws.String(uploadId),
	})
	if err != nil {
		return nil, fmt.Errorf("列出分片失败: %w", err)
	}
	parts := make([]MultipartPart, 0, len(out.Parts))
	for _, p := range out.Parts {
		parts = append(parts, MultipartPart{
			PartNumber: int(aws.ToInt32(p.PartNumber)),
			ETag:       aws.ToString(p.ETag),
		})
	}
	return parts, nil
}
