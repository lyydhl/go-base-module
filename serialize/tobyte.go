package serialize

import (
	"bytes"
	"encoding/binary"
)

// 序列化成 二进制
func SerializeToBytes(date interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, date)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// 反序列化
func DeserializeToString(byte []byte) (string, error) {
	var str string
	buf := bytes.NewBuffer(byte)

	err := binary.Read(buf, binary.LittleEndian, &str)
	if err != nil {
		return "", err
	}
	return str,nil
}
