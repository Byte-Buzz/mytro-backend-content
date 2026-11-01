package request

type CreateContentRequest struct {
	ContentType string `json:"content_type" binding:"required"`
}
