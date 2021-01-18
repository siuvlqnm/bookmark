package request

import "github.com/siuvlqnm/bookmark/model"

type NewBookmark struct {
	MSeaEngineId uint32 `json:mSeaEngineId`
	Link         string `json:"link"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	TagStr       string `json:tagStr`
}

type GetBookmarkList struct {
	PageInfo
	model.CusBookmark
}
