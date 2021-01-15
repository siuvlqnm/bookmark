package service

import (
	"errors"

	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
	"gorm.io/gorm"
)

func CreateWebSite(w *model.CusWebsite) (err error, webInfo *model.CusWebsite) {
	var website model.CusWebsite
	if !errors.Is(global.GVA_DB.Where("domain = ?", w.Domain).First(&website).Error, gorm.ErrRecordNotFound) {
		return err, &website
	}
	err = global.GVA_DB.Create(&w).Error
	return err, w
}

func GetWebSite(domain string, port int) (err error, website *model.CusWebsite) {
	err = global.GVA_DB.Where("domain = ? AND port = ?", domain, port).First(&website).Error
	return
}
