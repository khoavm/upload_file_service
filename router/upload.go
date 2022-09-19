package router

import (
	"github.com/gin-gonic/gin"
	"upload_file_handler/controller"
)

func AddUploadRouter(router *gin.Engine) {
	trackingRouter := router.Group("/upload")
	trackingRouter.POST("", controller.UploadFile)

}
