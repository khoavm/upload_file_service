package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"upload_file_handler/service"
	"upload_file_handler/util"
)

func UploadFile(c *gin.Context) {

	err := service.UploadFile(c)
	if err != nil {
		util.SendError(c, 400, err)
		return
	}
	util.Success(c, "Upload successfully!")
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

}
