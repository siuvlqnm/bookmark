package request

type CusBookmarkGroup struct {
	GSeaEngineId uint32 `json:"gSeaEngineId"`
	ParentId     int    `json:"parentId"`
	GroupName    string `json"groupName"`
	GroupIcon    string `json:"groupIcon"`
	IsArchive    uint8  `json:"isArchive"`
}
