package request

import "github.com/siuvlqnm/bookmark/model"

type NewBookmark struct {
	Link        string `json:"link"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetBookmarkList struct {
	PageInfo
	model.CusBookmark
}
