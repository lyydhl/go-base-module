package serialize

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
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

// GetBytes
func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
