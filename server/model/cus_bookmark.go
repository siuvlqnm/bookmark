package model

import (
	"github.com/siuvlqnm/bookmark/global"
)

type CusBookmark struct {
	global.GVA_MODEL
	MSeaEngineId uint32     `json:"mSeaEngineId"`
	CusWebId     uint       `json:"-"`
	CusUserId    uint       `json:"-"`
	Path         string     `json:"path"`
	Query        string     `json:"query"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	CusGroupId   uint       `json:"cusGroupId"`
	CusTagStr    string     `json:"cusTagStr"`
	IsStar       uint8      `json:"isStar"`
	CusWebsite   CusWebsite `json:"cusWebsite" gorm:"foreignKey:CusWebId"`
}
