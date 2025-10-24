package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// GenerateAddres 生成一个唯一address
//
// 返回值：
//   - string 地址
func GenerateAddress() string {
	data := time.Now().String()
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
