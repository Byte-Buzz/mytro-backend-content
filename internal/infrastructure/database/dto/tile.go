package dto

import "time"

type TileDTO struct {
	ID                 string  `gorm:"primaryKey;type:uuid"`
	OwnerID            string  `gorm:"type:uuid;not null"`
	Type               string  `gorm:"type:tile_type;not null"`
	Title              *string `gorm:"type:text"`
	Description        *string `gorm:"type:text"`
	Visibility         string  `gorm:"type:tile_visibility;not null;default:'private'"`
	Status             string  `gorm:"type:tile_status;not null;default:'draft'"`
	PublishRequestedAt *time.Time
	PublishedAt        *time.Time
	CreatedAt          time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt          time.Time `gorm:"not null;default:current_timestamp"`
	Deleted            bool      `gorm:"not null;default:false"`
	DeletedAt          *time.Time
	Metadata           *TileMetadataDTO `gorm:"foreignKey:TileID;constraint:OnDelete:CASCADE"`
	Files              []TileFileDTO    `gorm:"foreignKey:TileID;references:ID;constraint:OnDelete:CASCADE"`
}

func (TileDTO) TableName() string {
	return "tiles"
}
