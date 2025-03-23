package fileHandler

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

const (
	MaxFileSize = 5 * 1024 * 1024 // 5MB
)

var (
	AllowedTypes map[string]map[string]bool
	cld          *cloudinary.Cloudinary
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Initialize Cloudinary client
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		log.Fatal("CLOUDINARY_URL not set in environment variables")
	}

	cld, err = cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatal("Failed to initialize Cloudinary:", err)
	}

	// Define allowed file types
	AllowedTypes = map[string]map[string]bool{
		"resume": {
			".pdf":  true,
			".doc":  true,
			".docx": true,
		},
		"project": {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		},
		"bio": {
			".jpg":  true,
			".jpeg": true,
			".png":  true,
		},
	}
}

// ValidateFile checks file size and type
func ValidateFile(file *multipart.FileHeader, category string) error {
	if file.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds 5MB limit")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if allowed, ok := AllowedTypes[category][ext]; !ok || !allowed {
		return fmt.Errorf("invalid file type '%s' for category '%s'", ext, category)
	}
	return nil
}

// UploadFile uploads a file to Cloudinary and returns the secure URL and public ID
func UploadFile(file *multipart.FileHeader, category string, userID int) (string, string, string, error) {
	f, err := file.Open()
	if err != nil {
		return "", "", "", fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	// Get the file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// Create a unique public ID without the extension (Cloudinary will append it)
	publicID := fmt.Sprintf("%d_%s", userID, strings.TrimSuffix(file.Filename, ext))

	ctx := context.Background()
	resp, err := cld.Upload.Upload(ctx, f, uploader.UploadParams{
		PublicID: publicID,
		Folder:   fmt.Sprintf("portfolio/%s", category),
	})
	if err != nil {
		return "", "", "", fmt.Errorf("failed to upload to Cloudinary: %v", err)
	}

	return resp.SecureURL, publicID, ext, nil
}

// DeleteFile deletes a file from Cloudinary by public ID
func DeleteFile(publicID string) error {
	ctx := context.Background()
	_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from Cloudinary: %v", err)
	}
	return nil
}
