package dto

type ContentMetadataDTO struct {
	ContentID       string `gorm:"primaryKey;type:uuid;references:ID"`
	DurationSeconds *int
	DominantColor   *string
}

func (ContentMetadataDTO) TableName() string {
	return "content_metadata"
}
