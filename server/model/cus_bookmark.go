package model

import "github.com/siuvlqnm/bookmark/global"

type CusBookmark struct {
	global.GVA_MODEL
	CusWebId    uint   `json:"cusWebId"`
	CusUserId   uint   `json:"cusUserId"`
	Path        string `json:"path"`
	Query       string `json:"query"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CusGroupId  uint   `json:"cusGroupId"`
	CusTagStr   string `json:"cusTagStr"`
}
