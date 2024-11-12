package eui

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_initFontAwesomeJson(t *testing.T) {
	// 打印 fontAwesomemMap 所有的图标名和Unicode码点
	var sb strings.Builder
	for k, v := range fontAwesomemMap {
		sb.WriteString(fmt.Sprintf("%s, %d, %x\n", k, v, v))
	}
	os.WriteFile("fontAwesomemMap6.6.0.txt", []byte(sb.String()), 0666)
}
