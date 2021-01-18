package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
)

func CreateBookmarkGroup(group model.CusBookmarkGroup) (err error, g *model.CusBookmarkGroup) {
	err = global.GVA_DB.Create(&group).Error
	return err, &group
}

func UpateGroupGSeaEngineId(id int, GSeaEngineId uint32) (err error) {
	var group model.CusBookmarkGroup
	err = global.GVA_DB.Model(&group).Where("id = ?", id).Update("g_sea_engine_id", GSeaEngineId).Error
	return
}
