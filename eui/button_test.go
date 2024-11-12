package eui

import (
	"fmt"
	"github.com/twgh/xcgui/xc"
	"strconv"
	"strings"
	"testing"
)

// 生成颜色
func Test_MakeColor(t *testing.T) {
	arr := []int{
		xc.RGBA(245, 108, 108, 255),
		xc.RGBA(245, 108, 108, 255),
		xc.RGBA(221, 97, 97, 255),
		-1,
		xc.RGBA(249, 167, 167, 255),
	}

	var sb strings.Builder
	for i, num := range arr {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(num))
	}
	fmt.Println(sb.String())
}
