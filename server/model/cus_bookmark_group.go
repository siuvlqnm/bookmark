package model

import "github.com/siuvlqnm/bookmark/global"

type CusBookmarkGroup struct {
	global.GVA_MODEL
	GSeaEngineId  uint32        `json:"gSeaEngineId"`
	GroupParentId int           `json:"-"`
	GroupName     string        `json:"groupName"`
	GroupIcon     string        `json:"groupIcon"`
	IsArchive     bool          `json:"isArchive"`
	CusBookmark   []CusBookmark `json:"bookmark" gorm:"foreignKey:CusGroupId"`
}
