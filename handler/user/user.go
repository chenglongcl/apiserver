package user

type CreateRequest struct {
	Username string `json:"username" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" binding:"required" validate:"min=1,max=128"`
	Mobile   string `json:"mobile" validate:"numeric,max=11"`
}

type CreateResponse struct {
	Username         string `json:"username"`
	Token            string `json:"token"`
	ExpiredAt        string `json:"expiredAt"`
	RefreshExpiredAt string `json:"refreshExpiredAt"`
}

type GetRequest struct {
	ID uint64 `form:"id"`
}

type GetResponse struct {
	ID         uint64 `json:"id"`
	Username   string `json:"username"`
	Mobile     string `json:"mobile"`
	CreateTime string `json:"createTime"`
}

type UpdateRequest struct {
	ID       uint64 `json:"id" binding:"required"`
	Password string `json:"password" binding:"required" validate:"min=1,max=128"`
	Mobile   string `json:"mobile" validate:"numeric,max=11"`
}

type DeleteRequest struct {
	ID uint64 `form:"id" binding:"required"`
}

type ListRequest struct {
	UserName string `form:"username"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}
