package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/siuvlqnm/bookmark/api/v1"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
	}
	return BaseRouter
}
