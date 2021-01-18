package router

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/api/v1/cus"
	"github.com/siuvlqnm/bookmark/middleware"
)

func InitCusBookmarkGroupRouter(Router *gin.RouterGroup) {
	BookmarkGroupRouer := Router.Group("c").Use(middleware.OperationRecord())
	{
		BookmarkGroupRouer.POST("/folder/new", cus.CreateNewGroup)
		BookmarkGroupRouer.PUT("/folder/update")
		BookmarkGroupRouer.DELETE("/folder/delete")
	}
}
