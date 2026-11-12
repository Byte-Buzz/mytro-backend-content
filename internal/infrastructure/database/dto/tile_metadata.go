package dto

type TileMetadataDTO struct {
	TileID          string `gorm:"primaryKey;type:uuid;references:ID"`
	DurationSeconds *int
	DominantColor   *string
}

func (TileMetadataDTO) TableName() string {
	return "tile_metadata"
}
