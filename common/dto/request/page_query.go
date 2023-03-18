package request

type PageQuery struct {
	PageNum  int `json:"pageNum" form:"pageNum"`   //页码
	PageSize int `json:"pageSize" form:"pageSize"` //每页数量
}
