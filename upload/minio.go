package upload

import (
	"context"
	"fmt"
	"milliy/config"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioUploader struct {
	client *minio.Client
}

func NewMinioUploader() (*MinioUploader, error) {
	cfg := config.Load()
	client, err := minio.New(cfg.Minio.MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.MINIO_ACCESS_KEY_ID, cfg.Minio.MINIO_SECRET_ACCESS_KEY, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %v", err)
	}

	return &MinioUploader{client: client}, nil
}

func (m *MinioUploader) UploadFile(bucketName string, file multipart.File, header *multipart.FileHeader) (string, error) {
	ctx := context.Background()
	fmt.Println("\n\n\n1\n\n\n\n")

	// Generate unique filename
	fileExt := filepath.Ext(header.Filename)
	fmt.Println("\n\n\n2\n\n\n\n")
	newFileName := uuid.NewString() + fileExt
	fmt.Println("\n\n\n3\n\n\n\n")

	// Upload the file
	_, err := m.client.PutObject(ctx, bucketName, newFileName, file, header.Size, minio.PutObjectOptions{
		ContentType: getContentType(fileExt),
	})
	fmt.Println("\n\n\n4\n\n\n\n")
	if err != nil {
		return "", fmt.Errorf("failed to upload file: %v", err)
	}

	// Set bucket policy for public access
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"AWS": ["*"]
				},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/*"]
			}
		]
	}`, bucketName)

	err = m.client.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		return "", fmt.Errorf("failed to set bucket policy: %v", err)
	}

	// Generate URL
	url := fmt.Sprintf("http://%s/%s/%s", config.Load().Minio.MINIO_ENDPOINT, bucketName, newFileName)
	fmt.Println("miniodan chiqvoti")
	return url, nil
}

func getContentType(fileExt string) string {
	switch fileExt {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".mp4":
		return "video/mp4"
	case ".mp3":
		return "audio/mpeg"
	default:
		return "application/octet-stream"
	}
}
