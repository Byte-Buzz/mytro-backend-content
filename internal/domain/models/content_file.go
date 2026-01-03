package models

type ContentFile struct {
	ContentID string
	FileID    string
	Role      FileRole
	Width     *int
	Height    *int
}
