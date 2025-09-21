package eui

import (
	"strconv"
	"strings"

	"github.com/twgh/xcgui/common"
)

// Xchar 传入 Unicode 码点转换到字符. 如 20013 是'中'.
func Xchar(UnicodePoint int32) string {
	return string(UnicodePoint)
}

// Xchar2 传入 Unicode 码点十六进制文本转换到字符. 如 4E2D 是'中'.
func Xchar2(UnicodePointHex string) string {
	decimal, _ := strconv.ParseInt(UnicodePointHex, 16, 32)
	return string(int32(decimal))
}

// JoinColorString 将颜色数组拼接成字符串, 以英文逗号分割.
func JoinColorString(colors ...uint32) string {
	var sb strings.Builder
	for i, color := range colors {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(common.Uint32ToA(color))
	}
	return sb.String()
}

const (
	// 加载
	svg_loading = `<svg t="1731132887070" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4306" width="16" height="16"><path d="M512 97c-11.4 0-20.8 9.3-20.8 20.8v166c0 11.4 9.3 20.8 20.8 20.8s20.8-9.3 20.8-20.8v-166c0-11.5-9.4-20.8-20.8-20.8zM247.9 218.6c-8.1-8.1-21.3-8.1-29.3 0s-8.1 21.3 0 29.3L336 365.3c8.1 8.1 21.3 8.1 29.3 0s8.1-21.3 0-29.3L247.9 218.6zM304.5 512c0-11.4-9.3-20.8-20.8-20.8h-166c-11.4 0-20.8 9.3-20.8 20.8s9.3 20.8 20.8 20.8h166c11.5 0 20.8-9.4 20.8-20.8zM335.9 658.7L218.6 776.1c-8.1 8.1-8.1 21.3 0 29.3 8.1 8.1 21.3 8.1 29.3 0L365.3 688c8.1-8.1 8.1-21.3 0-29.3s-21.3-8-29.4 0zM512 719.5c-11.4 0-20.8 9.3-20.8 20.8v166c0 11.4 9.3 20.8 20.8 20.8s20.8-9.3 20.8-20.8v-166c0-11.5-9.4-20.8-20.8-20.8zM688.1 658.7c-8.1-8.1-21.3-8.1-29.3 0s-8.1 21.3 0 29.3l117.4 117.4c8.1 8.1 21.3 8.1 29.3 0 8.1-8.1 8.1-21.3 0-29.3L688.1 658.7zM906.3 491.3h-166c-11.4 0-20.8 9.3-20.8 20.8s9.3 20.8 20.8 20.8h166c11.4 0 20.8-9.3 20.8-20.8s-9.4-20.8-20.8-20.8zM688.1 365.3l117.4-117.4c8.1-8.1 8.1-21.3 0-29.3s-21.3-8.1-29.3 0L658.7 335.9c-8.1 8.1-8.1 21.3 0 29.3s21.3 8.1 29.4 0.1z" p-id="4307"></path></svg>`
)
