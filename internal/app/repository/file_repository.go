package repository

import (
	"context"
	"encoding/json"
	"mytro-backend-content/internal/domain/models"
	"mytro-backend-content/internal/infrastructure/grpc/pb"
)

type FileRepository interface {
	CreateUploadURL(ctx context.Context, data models.UploadURLData) (*models.UploadURL, error)
}

type fileRepository struct {
	grpcClient pb.ContentStorageClient
}

// TODO: implement file repository
func NewFileRepository(grpcClient pb.ContentStorageClient) FileRepository {
	return &fileRepository{
		grpcClient: grpcClient,
	}
}

func (r *fileRepository) CreateUploadURL(ctx context.Context, data models.UploadURLData) (*models.UploadURL, error) {
	settings := uint32(1)
	ttl := uint32(60 * 60)

	info := map[string]any{
		"content_id":    data.ContentID.String(),
		"preview_sizes": []int{480},
	}

	if data.IsVideo {
		settings |= 1 << 2
		info["media_type"] = "video"
	} else {
		settings |= 1 << 1
		info["media_type"] = "image"
		info["thumbnail_sizes"] = []int{300, 600}
	}

	infoJson, err := json.Marshal(info)

	if err != nil {
		return nil, err
	}

	response, err := r.grpcClient.CreateUpload(ctx, &pb.CreateUploadRequest{
		AppName:      "mytro",
		Filename:     data.Filename,
		ContentType:  data.ContentType,
		MaxSize:      data.MaxSize,
		Visibility:   pb.FileVisibility(getVisibility(data.Visibility)),
		Settings:     &settings,
		InfoJson:     string(infoJson),
		UploadUrlTtl: &ttl,
	})

	if err != nil {
		return nil, err
	}

	return &models.UploadURL{
		FileID:    response.FileId,
		UploadURL: response.UploadUrl,
		ExpiresAt: response.UploadExpiresAt,
	}, nil
}

func getVisibility(visibility models.ContentVisibility) pb.FileVisibility {
	switch visibility {
	case models.ContentVisibilityPublic:
		return pb.FileVisibility_PUBLIC
	default:
		return pb.FileVisibility_PRIVATE
	}
}
