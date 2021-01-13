package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/model/response"
	"github.com/siuvlqnm/bookmark/utils"
	"go.uber.org/zap"
)

func NewBookmark(c *gin.Context) {
	var N request.NewBookmark
	_ = c.ShouldBindJSON(&N)
	if err := utils.Verify(N, utils.NewBookmarkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, P := utils.ParseUrl(N.Link)
	global.GVA_LOG.Error("test", zap.Any("err", P))
}
