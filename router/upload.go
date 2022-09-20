package router

import (
	"github.com/gin-gonic/gin"
	"upload_file_handler/controller"
	"upload_file_handler/middleware"
)

func AddUploadRouter(router *gin.Engine) {
	trackingRouter := router.Group("/upload")
	trackingRouter.POST("", middleware.AuthorizeUser(), controller.UploadFile)

}
