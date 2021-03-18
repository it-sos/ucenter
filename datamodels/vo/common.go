package vo

type Paging struct {
	Page      int `json:"page"`      // 当前页数
	Size      int `json:"size"`      // 每页条数
	TotalPage int `json:"totalPage"` // 总页数
}
