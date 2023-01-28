package user

type CreateRequest struct {
	Username string `json:"username" binding:"min=1,max=16"`
	Password string `json:"password" binding:"min=6,max=18"`
	Mobile   string `json:"mobile" binding:"omitempty,numeric,min=11"`
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
	Password string `json:"password" binding:"omitempty,min=6,max=18"`
	Mobile   string `json:"mobile" binding:"omitempty,numeric,min=11"`
}

type DeleteRequest struct {
	ID uint64 `form:"id" binding:"required"`
}

type ListRequest struct {
	UserName string `form:"username"`
	Page     uint64 `form:"page"`
	Limit    uint64 `form:"limit"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
