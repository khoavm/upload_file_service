package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"upload_file_handler/service"
	"upload_file_handler/util"
)

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	uploadFileService := service.NewUploadFileService(c)
	if err := uploadFileService.ValidateFile(); err != nil {
		util.SendError(c, 403, err)
		return
	}
	err := uploadFileService.UploadFile()

	if err != nil {
		util.SendError(c, 400, err)
		return
	}
	util.Success(c, "Upload successfully!")

	log.Println(file.Filename)

}
