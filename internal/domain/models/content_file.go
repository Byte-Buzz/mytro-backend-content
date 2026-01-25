package models

import "github.com/google/uuid"

type ContentFile struct {
	ContentID uuid.UUID
	FileID    uuid.UUID
	Role      FileRole
	Width     *int
	Height    *int
}
