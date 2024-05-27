package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	ErrFileTooLarge = NewStatusError(400, errors.New("file too large"))
)

func HandleUpload(w http.ResponseWriter, r *http.Request, formKey string, maxBytes int64) (string, error) {
	// 设置内存限制（这里仅作为示例，根据需要调整）
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	// 解析 multipart/form-data
	err := r.ParseMultipartForm(maxBytes)
	if err != nil {
		return "", ErrFileTooLarge
	}
	// 获取上传的文件
	file, _, err := r.FormFile(formKey)
	if err != nil {
		return "", fmt.Errorf("error retrieving file: %w", err)
	}
	defer file.Close()

	// 读取文件内容到内存（实际应用中可能需要直接写入到磁盘或进行流式处理以避免大文件问题）
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	// 示例：将文件保存到服务器上的临时目录
	tempFile, err := os.CreateTemp("", "svc_apis_upload-*.tmp")
	if err != nil {
		return "", fmt.Errorf("error creating temp file: %w", err)
	}
	defer tempFile.Close()

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return "", fmt.Errorf("error writing file to disk: %w", err)
	}
	return tempFile.Name(), nil
}
