package dto

type TileFileDTO struct {
	TileID string `gorm:"primaryKey;type:uuid;references:ID"`
	FileID string `gorm:"type:uuid;not null"`
	Role   string `gorm:"primaryKey;type:file_role"`
	Width  *int
	Height *int
}

func (TileFileDTO) TableName() string {
	return "tile_files"
}
