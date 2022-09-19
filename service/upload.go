package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
	"time"
	"upload_file_handler/config"
	"upload_file_handler/db"
	"upload_file_handler/model"
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

	file, err := u.c.FormFile("file")
	if err != nil {
		return err
	}
	filePath := filepath.Join(config.Config.Storage, file.Filename)

	if err := u.c.SaveUploadedFile(file, filePath); err != nil {
		return err
	}

	fileMeta := model.FileMeta{
		Name:      file.Filename,
		Type:      file.Header.Get("Content-Type"),
		Length:    file.Size,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := db.DB.Create(&fileMeta)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
