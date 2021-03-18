package parameter

type Disabled struct {
	*Id
	IsDisabled string `json:"disabled" binding:"required"` // 禁用标志0=否，1=是
}

type Deleted struct {
	*Id
	IsDeleted string `json:"deleted" binding:"required"` // 删除标志0=否，1=是
}

type Page struct {
	Page int `json:"page"` // 页码
	Size int `json:"size"` // 每页条数
}
