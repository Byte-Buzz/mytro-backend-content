package dto

import "time"

type ContentDTO struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	OwnerID     string    `gorm:"type:uuid;not null;index:idx_content_owner_id"`
	Type        string    `gorm:"type:content_type;not null"`
	Title       string    `gorm:"type:text"`
	Description string    `gorm:"type:text"`
	Visibility  string    `gorm:"type:content_visibility;not null;default:'private'"`
	Status      string    `gorm:"type:content_status;not null;default:'draft';index:idx_content_status"`
	CreatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"not null;default:current_timestamp"`
	Deleted     bool      `gorm:"not null;default:false;index:idx_content_deleted"`
	DeletedAt   *time.Time
	Metadata    *ContentMetadataDTO `gorm:"constraint:OnDelete:CASCADE"`
	Files       []ContentFileDTO    `gorm:"foreignKey:ContentID;references:ID;constraint:OnDelete:CASCADE"`
}
