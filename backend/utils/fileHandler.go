package fileHandler

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxFileSize = 5 * 1024 * 1024
)

var (
	UploadDir    map[string]string
	AllowedTypes map[string]map[string]bool
)

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	// Initialize upload directories map
	UploadDir = map[string]string{
		"resume":  filepath.Join(cwd, "upload", "resume"),
		"project": filepath.Join(cwd, "upload", "project"),
	}

	// Initialize allowed types for different file categories
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
	}

	// Create all upload directories
	for _, dir := range UploadDir {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal("Failed to create upload directory:", err)
		}
	}
}

// ValidateFile check file size and type based on category
func ValidateFile(file *multipart.FileHeader, category string) error {
	if file.Size > MaxFileSize {
		return fmt.Errorf("file size too large. Maximum size is 5MB")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !AllowedTypes[category][ext] {
		return fmt.Errorf("invalid file type for %s", category)
	}
	return nil
}

// DeleteFile removes a file from the specified category
func DeleteFile(filename string, category string) error {
	filePath := filepath.Join(UploadDir[category], filename)
	if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Failed to delete file: %v", err)
	}
	return nil
}

// GetFilePath returns the full path for a file
func GetFilePath(filename string, category string) string {
	return filepath.Join(UploadDir[category], filename)
}
