package cus

import (
	"github.com/gin-gonic/gin"
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/response"
	"github.com/siuvlqnm/bookmark/service"
	"github.com/siuvlqnm/bookmark/utils"
	"go.uber.org/zap"
)

func CreateNewGroup(c *gin.Context) {
	var group model.CusBookmarkGroup
	_ = c.ShouldBindJSON(&group)
	if err := utils.Verify(group, utils.NewBookmarkGroupVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, cbg := service.CreateBookmarkGroup(group); err != nil {
		global.GVA_LOG.Error("添加失败", zap.Any("err", err))
		response.FailWithMessage("添加失败", c)
		return
	} else {
		murmur32 := utils.GetMurmur32("group:", int(cbg.ID))
		service.UpateGroupGSeaEngineId(int(cbg.ID), murmur32)
		response.OkWithMessage("添加成功", c)
	}
}
