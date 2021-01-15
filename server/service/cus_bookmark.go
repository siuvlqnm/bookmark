package service

import (
	"github.com/siuvlqnm/bookmark/global"
	"github.com/siuvlqnm/bookmark/model"
)

func CreateBookmark(b model.CusBookmark) (err error, bookmark *model.CusBookmark) {
	err = global.GVA_DB.Create(&b).Error
	return err, &b
}
