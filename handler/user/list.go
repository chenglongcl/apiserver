package user

import (
	"apiserver/dal/apiserverdb/apiserverquery"
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service"
	"apiserver/service/userservice"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

// List
// @Description:
// @param c
func List(c *gin.Context) {
	var (
		r    ListRequest
		ps   util.PageSetting
		resp handler.RegularListResponse
	)
	if err := c.BindQuery(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	ps.Setting(r.Page, r.Limit)
	userService := userservice.NewUserService(c)
	infos, count, errNo := userService.InfoList(&service.ListParams{
		PS: ps,
		Fields: []field.Expr{
			apiserverquery.Q.TbUser.ALL,
		},
		Conditions: append(func() []gen.Condition {
			conditions := make([]gen.Condition, 0)
			if r.UserName != "" {
				conditions = append(conditions, apiserverquery.Q.TbUser.Username.Like(util.StringBuilder("%", r.UserName, "%")))
			}
			return conditions
		}(), []gen.Condition{}...),
		Orders: []field.Expr{
			apiserverquery.Q.TbUser.ID.Desc(),
		},
	})
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	resp = handler.RegularListResponse{Page: util.PageUtil(count, ps.Page, ps.Limit, infos)}
	handler.SendResponse(c, nil, resp)
}
