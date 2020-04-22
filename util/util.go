package util

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"time"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}

func GenStr(length int) (string, error) {
	b := make([]byte, length)
	n, err := rand.Read(b)
	if n != len(b) || err != nil {
		return "", errors.New("Could not successfully read from the system CSPRNG")
	}
	return hex.EncodeToString(b), nil
}

func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-ID")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}

func GetGmtIso8601(expireEnd int64) string {
	tokenExpire := time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
