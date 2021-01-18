package model

import "github.com/siuvlqnm/bookmark/global"

type CusBookmarkGroup struct {
	global.GVA_MODEL
	GSeaEngineId uint32 `json:"gSeaEngineId"`
	ParentId     int    `json:"parentId"`
	GroupName    string `json:"groupName"`
	GroupIcon    string `json:"groupIcon"`
	IsArchive    uint8  `json:"isArchive"`
}
