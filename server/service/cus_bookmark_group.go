package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
)

func GetBookMarkGroupList(isArchive string) (err error, list interface{}) {
	var groupList []model.CusBookmarkGroup
	err = global.GVA_DB.Model(&model.CusBookmarkGroup{}).Preload("CusBookmark").Order("id desc").Find(&groupList).Error
	return err, groupList
}

func CreateBookmarkGroup(group model.CusBookmarkGroup) (err error, g *model.CusBookmarkGroup) {
	err = global.GVA_DB.Create(&group).Error
	return err, &group
}

func UpateGroupGSeaEngineId(id int, GSeaEngineId uint32) (err error) {
	var group model.CusBookmarkGroup
	err = global.GVA_DB.Model(&group).Where("id = ?", id).Update("g_sea_engine_id", GSeaEngineId).Error
	return
}

func UpdateBookmarkGroup(u *model.CusBookmarkGroup) (err error) {
	var g *model.CusBookmarkGroup
	upDateMap := make(map[string]interface{})
	upDateMap["group_parent_id"] = u.GroupParentId
	upDateMap["group_name"] = u.GroupName
	upDateMap["group_icon"] = u.GroupIcon
	upDateMap["is_archive"] = u.IsArchive
	err = global.GVA_DB.Model(&g).Where("g_sea_engine_id = ?", u.GSeaEngineId).Updates(upDateMap).Error
	return
}

func DeleteBookmarkGroup(GSeaEngineId uint32) (err error) {
	var g model.CusBookmarkGroup
	return global.GVA_DB.Where("g_sea_engine_id = ?", GSeaEngineId).Delete(&g).Error
}
