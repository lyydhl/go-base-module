package encrypt

import "encoding/base64"

// Base64 加密
func Base64Encrypt(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

// Base64 解密
func Base64Decrypt(str string) (string, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(str)
	return string(decodeBytes), err
}
