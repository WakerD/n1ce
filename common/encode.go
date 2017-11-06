package common

import (
	"crypto/md5"
	"encoding/base64"
)

func EncodeMessageMd5(msg string) string {
	h := md5.New()
	coding := base64.NewEncoding(beego.AppConfig.String("base64key"))
	h.Write([]byte(msg))
	key := []byte(beego.AppConfig.String("md5key"))
	cipherStr := h.Sum([]byte(key))

	return coding.EncodeToString(cipherStr)
}
