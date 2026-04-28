package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile handles multipart file uploads and saves to local storage.
func UploadFile(r *http.Request, formKey, destDir, filename string) (string, error) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return "", fmt.Errorf("failed to parse multipart form: %w", err)
	}

	file, handler, err := r.FormFile(formKey)
	if err != nil {
		return "", fmt.Errorf("failed to get file from form: %w", err)
	}
	defer file.Close()

	if filename == "" {
		filename = handler.Filename
	}

	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fullPath := filepath.Join(destDir, filename)
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file on server: %w", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return filepath.ToSlash(fullPath), nil
}
