package sd

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"github.com/spf13/viper"
	"hash"
	"net/http"
	"time"
)

type PolicyConfig struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
}

func GenerateSignature(c *gin.Context) {
	accessKeyId := viper.GetString("aliyun_oss.access_key_id")
	accessKeySecret := viper.GetString("aliyun_oss.access_key_secret")
	host := viper.GetString("aliyun_oss.host")
	expireTime := int64(30)
	now := time.Now()
	dir := now.Format("20060102") + "/"

	nowTimestamp := now.Unix()
	expireEnd := nowTimestamp + expireTime
	tokenExpire := getGmtIso8601(expireEnd)

	//create post policy json
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, dir)
	pc := PolicyConfig{}
	pc.Expiration = tokenExpire
	pc.Conditions = append(pc.Conditions, condition)
	//calucate signature
	result, err := jsoniter.Marshal(pc)
	if err != nil {
		SendResponse(c, errno.ErrOssGenerateSignatureFail, nil)
		return
	}
	deByte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash {
		return sha1.New()
	}, []byte(accessKeySecret))
	//io.WriteString(h, debyte)
	h.Write([]byte(deByte))
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	pt := PolicyToken{
		AccessKeyId: accessKeyId,
		Host:        host,
		Expire:      expireEnd,
		Signature:   signedStr,
		Policy:      deByte,
		Directory:   dir,
	}
	c.JSON(http.StatusOK, pt)
}

func getGmtIso8601(expireEnd int64) string {
	tokenExpire := time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
