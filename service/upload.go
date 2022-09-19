package service

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"upload_file_handler/config"
)

func UploadFile(c *gin.Context) error {
	// Upload the file to specific dst.
	// single file

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	filePath := filepath.Join(config.Config.Storage, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return err
	}
	return nil
}
