package parameter

type Id struct {
	Id uint `json:"id" binding:"required"` // 用户id
}

type Password struct {
	*Id
	Password string `json:"password" binding:"required"` // 密码
}

type Disabled struct {
	*Id
	IsDisabled string `json:"disabled" binding:"required"` // 禁用标志0=否，1=是
}

type Deleted struct {
	*Id
	IsDeleted string `json:"deleted" binding:"required"` // 删除标志0=否，1=是
}

type User struct {
	Account    string `json:"account" binding:"required"` // 帐号
	Password   string `json:"password" binding:"required"`
	Nickname   string `json:"nickname" binding:"required"` // 昵称
	Phone      string `json:"phone"`
	Expired    uint   `json:"expired" default:"0"`  // 有效期0=永久，unix时间戳
	IsDisabled uint8  `json:"disabled" default:"0"` // 禁用状态1=是，0=否
	IsDeleted  uint8  `json:"deleted" default:"0"`  // 删除标志1=是，0=否
}
