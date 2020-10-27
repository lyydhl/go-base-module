package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func Md5Fun01(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Md5Fun02(str string) string {
	m5 := md5.New()
	_, err := m5.Write([]byte(str))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(m5.Sum(nil))
}

func Md5Fun03(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", md5.Sum(nil))
}
