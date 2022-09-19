package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
	"upload_file_handler/config"
)

type UploadFileInterface interface {
	ValidateFile() error
	UploadFile() error
}

type UploadFileService struct {
	c *gin.Context
}

func NewUploadFileService(c *gin.Context) UploadFileInterface {

	return &UploadFileService{
		c: c,
	}
}
func (u *UploadFileService) ValidateFile() error {
	file, err := u.c.FormFile("file")
	if err != nil {
		return err
	}
	fileContentType := file.Header.Get("Content-Type")
	isImage := strings.Contains(fileContentType, "image")
	if isImage == false {
		return errors.New("file must be image")
	}
	fmt.Println(fileContentType)
	return nil
}

func (u *UploadFileService) UploadFile() error {
	// Upload the file to specific dst.
	// single file

	file, err := u.c.FormFile("file")
	if err != nil {
		return err
	}
	filePath := filepath.Join(config.Config.Storage, file.Filename)
	if err := u.c.SaveUploadedFile(file, filePath); err != nil {
		return err
	}
	return nil
}
