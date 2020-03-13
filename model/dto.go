package model

type bookVo struct {
	Type   string `json:"type,string form:type,string"`
	Author string `json:"author,string from:author,string"`
}

type Updatebook struct {
	BookId    int64  `json:"bookId,string" form:"bookId,string"`
	UserId    int64  `json:"userId,string" form:"userId,string"`
	LibraryId int64  `json:"libraryId,string" form:"libraryId,string"`
	StartTime string `json:"startTime" form:"startTime"`
	EndTime   string `json:"endTime" form:"endTime"`
}
