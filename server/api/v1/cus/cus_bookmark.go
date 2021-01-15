package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
	"github.com/siuvlqnm/bookmark/model/response"
	"github.com/siuvlqnm/bookmark/service"
	"github.com/siuvlqnm/bookmark/utils"
	"go.uber.org/zap"
)

func CreateBookmark(c *gin.Context) {
	var N request.NewBookmark
	_ = c.ShouldBindJSON(&N)
	if err := utils.Verify(N, utils.NewBookmarkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, P := utils.ParseUrl(N.Link)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	website := &model.CusWebsite{Protocol: P.Protocol, Domain: P.Domain, Port: int(P.Port)}
	service.CreateWebSite(website)
	bookmark := &model.CusBookmark{CusUserId: getUserID(c), Path: P.Path, Query: P.Query}
	if err, _ := service.CreateBookmark(*bookmark); err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}
