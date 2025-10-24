package utils

import "time"

// FormatTimeString 格式化时间字符串
//
// 返回值:
//   - 格式化后的时间字符串，格式为 "2006-01-02T15:04:05Z07:00"
func FormatTimeString() string {
	return time.Now().Format(time.RFC3339)
}
