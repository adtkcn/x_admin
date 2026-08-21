package queue_schema

// ImageWebpPayload image_webp 队列任务的消息契约。
// producer（UploadService.MaybeConvertWebp）与 consumer（task.ProcessImageWebp）共用，
// 字段 json tag 保持一致。
type ImageWebpPayload struct {
	FilePath   string `json:"file_path"`    // 原图磁盘存储 key（含扩展名）
	Ext        string `json:"ext"`          // 原图扩展名（不含点）
	FileHashId string `json:"file_hash_id"` // x_common_file_hash.id，转换完成后回写新扩展名
}
