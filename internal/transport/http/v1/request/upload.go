package request

type CreateUploadRequest struct {
	ContentType string `json:"content_type" binding:"required"`
}
