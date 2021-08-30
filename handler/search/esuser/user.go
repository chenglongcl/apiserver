package esuser

type MGetRequest struct {
	IDS string `form:"ids" binding:"required"`
}
