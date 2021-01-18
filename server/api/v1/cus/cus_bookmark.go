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
	_, w := service.CreateWebSite(website)
	bookmark := &model.CusBookmark{CusWebId: w.ID, CusUserId: getUserID(c), Path: P.Path, Query: P.Query}
	if err, cbm := service.CreateBookmark(*bookmark); err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
	} else {
		murmur32 := utils.GetMurmur32("bookmark:", int(cbm.ID))
		service.UpateBookmarkMSeaEngineId(int(cbm.ID), murmur32)
		response.OkWithMessage("添加成功", c)
	}
}

func GetBookmarkList(c *gin.Context) {
	var G request.GetBookmarkList
	_ = c.ShouldBindJSON(&G)
	if err := utils.Verify(G.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err, list, total := service.GetBookmarkList(getUserID(c), G.CusBookmark, G.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     G.Page,
			PageSize: G.PageSize,
		}, "获取成功", c)
	}
}

func UpdateBookmark(c *gin.Context) {
	var U request.NewBookmark
	_ = c.ShouldBindJSON(&U)
	if err := utils.Verify(U, utils.UpdateBookmarkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, P := utils.ParseUrl(U.Link)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	website := &model.CusWebsite{Protocol: P.Protocol, Domain: P.Domain, Port: int(P.Port)}
	_, w := service.CreateWebSite(website)
	bookmark := &model.CusBookmark{CusWebId: w.ID, Path: P.Path, Query: P.Query, Title: U.Title, Description: U.Description, CusTagStr: U.TagStr}
	if err = service.UpdateBookmar(U.MSeaEngineId, bookmark); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
		return
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func DeleteBookmark(c *gin.Context) {
	var D request.NewBookmark
	_ = c.ShouldBindJSON(&D)
	if err := utils.Verify(D, utils.DeleteBookmarkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := service.DeleteBookmark(D.MSeaEngineId); err != nil {
		global.GVA_LOG.Error("删除失败", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
