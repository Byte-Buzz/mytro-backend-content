package dto

type ContentFileDTO struct {
	ContentID string `gorm:"primaryKey;type:uuid;references:ID"`
	FileID    string `gorm:"type:uuid;not null"`
	Role      string `gorm:"primaryKey;type:file_role"`
	Width     *int
	Height    *int
}

func (ContentFileDTO) TableName() string {
	return "content_files"
}
