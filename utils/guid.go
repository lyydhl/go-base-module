package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"github.com/lyydhl/go-base-module/encrypt"
)

// 生成一个GUID
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return encrypt.Md5Fun02(base64.URLEncoding.EncodeToString(b))
}
