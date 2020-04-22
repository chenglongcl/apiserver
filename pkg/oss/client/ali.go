package client

import (
	"apiserver/pkg/errno"
	"apiserver/util"
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/json-iterator/go"
	"github.com/spf13/viper"
	"hash"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time"
)

var ali *Ali

type Ali struct {
	Client *oss.Client
}
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

func InitAliClient() {
	var Client *oss.Client
	aliConfig := map[string]string{
		"aliyunEndPoint":        viper.GetString("aliyun_oss.end_point"),
		"aliyunAccessKeyId":     viper.GetString("aliyun_oss.access_key_id"),
		"aliyunAccessKeySecret": viper.GetString("aliyun_oss.access_key_secret"),
	}
	if aliConfig["aliyunEndPoint"] != "" &&
		aliConfig["aliyunAccessKeyId"] != "" &&
		aliConfig["aliyunAccessKeySecret"] != "" {
		Client, _ = oss.New(aliConfig["aliyunEndPoint"],
			aliConfig["aliyunAccessKeyId"],
			aliConfig["aliyunAccessKeySecret"])
	}
	ali = &Ali{
		Client: Client,
	}
}

func DefaultAliClient() *Ali {
	return ali
}

func (o *Ali) Upload(file multipart.File, header *multipart.FileHeader) (string, *errno.Errno) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", nil
	}
	bucketName := viper.GetString("aliyun_oss.bucket_name")
	if o.Client == nil {
		return "", errno.ErrAliYunOssUploadFail
	}
	// 获取存储空间。
	bucket, err := o.Client.Bucket(bucketName)
	if err != nil {
		return "", errno.ErrAliYunBucket
	}
	newFileName, _ := util.GenStr(16)
	objectKey := time.Now().Format("20060102") + "/" + newFileName +
		strings.ToLower(path.Ext(header.Filename))
	var fileUrl string
	finished := make(chan bool, 1)
	go func() {
		bucketInfo, _ := o.Client.GetBucketInfo(bucketName)
		fileUrl = "http://" + bucketInfo.BucketInfo.Name + "." +
			bucketInfo.BucketInfo.ExtranetEndpoint + "/" + objectKey
		close(finished)
	}()
	// 上传Byte数组。
	err = bucket.PutObject(objectKey, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return "", errno.ErrAliYunOssUploadFail
	}
	<-finished
	return fileUrl, nil
}

func (o *Ali) WebUploadSign() (*PolicyToken, *errno.Errno) {
	accessKeyId := o.Client.Config.AccessKeyID
	accessKeySecret := o.Client.Config.AccessKeySecret
	bucketName := viper.GetString("aliyun_oss.bucket_name")
	bucketInfo, _ := o.Client.GetBucketInfo(bucketName)
	host := "http://" + bucketInfo.BucketInfo.Name + "." +
		bucketInfo.BucketInfo.ExtranetEndpoint
	expireTime := int64(30)
	now := time.Now()
	dir := now.Format("20060102") + "/"

	nowTimestamp := now.Unix()
	expireEnd := nowTimestamp + expireTime
	tokenExpire := util.GetGmtIso8601(expireEnd)

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
		return nil, errno.ErrOssGenerateSignatureFail
	}
	deByte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash {
		return sha1.New()
	}, []byte(accessKeySecret))
	//io.WriteString(h, debyte)
	h.Write([]byte(deByte))
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	pt := &PolicyToken{
		AccessKeyId: accessKeyId,
		Host:        host,
		Expire:      expireEnd,
		Signature:   signedStr,
		Policy:      deByte,
		Directory:   dir,
	}
	return pt, nil
}
