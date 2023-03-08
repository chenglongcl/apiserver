package user

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

// Refresh
// @Description:
// @param c
func Refresh(c *gin.Context) {
	if ctx, err, t, e, r := token.ParseRefreshRequest(c); err != nil {
		handler.SendResponseUnauthorized(c, errno.ErrTokenInvalid, nil)
	} else {
		handler.SendResponse(c, nil, CreateResponse{
			Username:         ctx.Username,
			Token:            t,
			ExpiredAt:        e,
			RefreshExpiredAt: r,
		})
	}
}
