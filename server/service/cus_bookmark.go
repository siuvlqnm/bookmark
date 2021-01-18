package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"github.com/siuvlqnm/bookmark/model/request"
)

func CreateBookmark(b model.CusBookmark) (err error, bookmark *model.CusBookmark) {
	err = global.GVA_DB.Create(&b).Error
	return err, &b
}

func GetBookmarkList(userId uint, where model.CusBookmark, info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.CusBookmark{})
	var bookmarkList []model.CusBookmark

	db = db.Where("cus_user_id = ?", userId)

	if where.CusGroupId != 0 {
		db = db.Where("cus_group_id = ?", where.CusGroupId)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, bookmarkList, total
	} else {
		err = db.Preload("CusWebsite").Order("id desc").Limit(limit).Offset(offset).Find(&bookmarkList).Error
	}
	return err, bookmarkList, total
}

func UpateBookmarkMSeaEngineId(id int, MSeaEngineId uint32) (err error) {
	var bookmark model.CusBookmark
	err = global.GVA_DB.Model(&bookmark).Where("id = ?", id).Update("m_sea_engine_id", MSeaEngineId).Error
	return
}

func UpdateBookmar(MSeaEngineId uint32, b *model.CusBookmark) (err error) {
	var bookmark model.CusBookmark
	err = global.GVA_DB.Model(&bookmark).Where("m_sea_engine_id = ?", MSeaEngineId).Update("CusWebId", b.CusWebId).Update("Path", b.Path).Update("Query", b.Query).Update("Title", b.Title).Update("Description", b.Description).Update("CusTagStr", b.CusTagStr).Error
	return
}
