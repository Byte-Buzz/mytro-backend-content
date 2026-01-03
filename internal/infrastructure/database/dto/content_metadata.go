package dto

type ContentMetadataDTO struct {
	ContentID       string `gorm:"primaryKey;type:uuid;references:ID"`
	DurationSeconds *int
	DominantColor   *string
}
