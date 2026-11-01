package models

type ContentType string
type ContentVisibility string
type ContentStatus string
type FileRole string

const (
	ContentTypeImage ContentType = "image"
	ContentTypeVideo ContentType = "video"

	ContentVisibilityPublic   ContentVisibility = "public"
	ContentVisibilityPrivate  ContentVisibility = "private"
	ContentVisibilityUnlisted ContentVisibility = "unlisted"

	ContentStatusDraft      ContentStatus = "draft"
	ContentStatusProcessing ContentStatus = "processing"
	ContentStatusReady      ContentStatus = "ready"
	ContentStatusFailed     ContentStatus = "failed"

	FileRoleOriginal        FileRole = "original"
	FileRoleThumbnailSmall  FileRole = "thumbnail_small"
	FileRoleThumbnailMedium FileRole = "thumbnail_medium"
	FileRoleVideoPreview    FileRole = "video_preview"
)
