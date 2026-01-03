package dto

type ContentFileDTO struct {
	ContentID string `gorm:"primaryKey;type:uuid;references:ID"`
	FileID    string `gorm:"primaryKey;type:uuid"`
	Role      string `gorm:"type:file_role;primaryKey"`
	Width     *int
	Height    *int
}
