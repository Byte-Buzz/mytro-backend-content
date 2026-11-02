package response

import "time"

type CreateUploadResponse struct {
	ContentID string    `json:"content_id"`
	UploadURL string    `json:"upload_url"`
	ExpiresAt time.Time `json:"expires_at"`
}
