package models

type TileType string
type TileVisibility string
type TileStatus string
type FileRole string

const (
	TileTypeImage TileType = "image"
	TileTypeVideo TileType = "video"

	TileVisibilityPublic   TileVisibility = "public"
	TileVisibilityPrivate  TileVisibility = "private"
	TileVisibilityUnlisted TileVisibility = "unlisted"

	TileStatusDraft      TileStatus = "draft"
	TileStatusProcessing TileStatus = "processing"
	TileStatusReady      TileStatus = "ready"
	TileStatusFailed     TileStatus = "failed"

	FileRoleOriginal        FileRole = "original"
	FileRoleThumbnailSmall  FileRole = "thumbnail_small"
	FileRoleThumbnailMedium FileRole = "thumbnail_medium"
	FileRoleVideoPreview    FileRole = "video_preview"
)
