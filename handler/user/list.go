package user

import (
	"apiserver/dal/apiserverdb/apiserverquery"
	"apiserver/handler"
	"apiserver/pkg/errno"
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
	fields := make([]field.Expr, 0)
	conditions := make([]gen.Condition, 0)
	orders := make([]field.Expr, 0)
	fields = append(fields, apiserverquery.Q.TbUser.ALL)
	if r.UserName != "" {
		conditions = append(conditions, apiserverquery.Q.TbUser.Username.Like(util.StringBuilder("%", r.UserName, "%")))
	}
	orders = append(orders, apiserverquery.Q.TbUser.ID.Desc())
	info, count, errNo := userService.GetUserList(ps, fields, conditions, orders)
	if errNo != nil {
		handler.SendResponse(c, errNo, nil)
		return
	}
	resp = handler.RegularListResponse{Page: util.PageUtil(count, ps.Page, ps.Limit, info)}
	handler.SendResponse(c, nil, resp)
}
