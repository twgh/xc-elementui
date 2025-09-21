package eui

import (
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

type Elementui struct {
	hFontAwesomeMap map[string]int // FontAwesome 字体句柄
	dpi             int32          // 窗口 dpi
}

// NewElementui 创建 Elementui 对象.
//
// fontSize: 字体大小. 一般使用 12.
//
// dpi: 窗口 dpi. 使用窗口.GetDPI()获取.
func NewElementui(fontSize, dpi int32) *Elementui {
	p := &Elementui{}
	p.dpi = dpi
	p.hFontAwesomeMap = make(map[string]int)
	p.hFontAwesomeMap["fa-solid"] = xc.XFont_CreateFromMem(fontAwesomeSolid, fontSize, xcc.FontStyle_Regular)
	p.hFontAwesomeMap["fa-brands"] = xc.XFont_CreateFromMem(fontAwesomeBrands, fontSize, xcc.FontStyle_Regular)
	p.hFontAwesomeMap["fa-regular"] = xc.XFont_CreateFromMem(fontAwesomeRegular, fontSize, xcc.FontStyle_Regular)
	xc.XC_SetTextRenderingHint(xcc.TextRenderingHintAntiAliasGridFit)
	return p
}

// GetFont 返回 FontAwesome 炫彩字体句柄 map.
//
// map 的键如下:
//   - fa-solid
//   - fa-brands
//   - fa-regular
func (e *Elementui) GetFont() map[string]int {
	return e.hFontAwesomeMap
}

// todo 用go来写显示所有图标, 根据json分类, 可搜索, 可复制名字, 十六进制, 十进制
