package response

type PageResult struct {
	List       interface{} `json:"list"`       //数据
	TotalCount int         `json:"totalCount"` //总条数
}
