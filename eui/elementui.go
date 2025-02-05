package eui

import (
	_ "embed"
	"encoding/json"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/fa-solid-900.ttf
var fontAwesomeSolid []byte

//go:embed res/fa-brands-400.ttf
var fontAwesomeBrands []byte

//go:embed res/fa-regular-400.ttf
var fontAwesomeRegular []byte

//go:embed res/icons.min.json
var fontAwesomeJson []byte

// fontAwesomemMap 存放FontAwesome图标名称和Unicode码点
var fontAwesomemMap map[string]int32

func init() {
	initFontAwesomeJson(fontAwesomeJson)
}

// iconFa FontAwesome图标信息
type iconFa struct {
	Styles  []string `json:"s"`
	Unicode int32    `json:"u"`
}

// initFontAwesomeJson 把FontAwesome icons的json数据解析后存入map中.
//
// jsonData: FontAwesome icons的json数据.
func initFontAwesomeJson(jsonData []byte) {
	// 解析 JSON
	var iconsMap map[string]iconFa
	err := json.Unmarshal(jsonData, &iconsMap)
	if err != nil {
		panic("Error unmarshalling JSON: " + err.Error())
	}
	fontAwesomemMap = make(map[string]int32)
	// 把icon风格和名字组合后放入map
	for name, icon := range iconsMap {
		for _, style := range icon.Styles {
			fontAwesomemMap["fa-"+style+" fa-"+name] = icon.Unicode
		}
	}
}

type Elementui struct {
	hFontAwesomeMap map[string]int // FontAwesome字体句柄
	dpi             int32          // 窗口dpi
}

// NewElementui 创建 Elementui 对象.
//
// fontSize: 字体大小. 一般使用12.
//
// dpi: 窗口dpi. 使用窗口.GetDPI()获取.
func NewElementui(fontSize, dpi int32) *Elementui {
	p := &Elementui{}
	p.dpi = dpi
	p.hFontAwesomeMap = make(map[string]int)
	p.hFontAwesomeMap["fa-solid"] = xc.XFont_CreateFromMem(fontAwesomeSolid, fontSize, xcc.FontStyle_Regular)
	p.hFontAwesomeMap["fa-brands"] = xc.XFont_CreateFromMem(fontAwesomeBrands, fontSize, xcc.FontStyle_Regular)
	p.hFontAwesomeMap["fa-regular"] = xc.XFont_CreateFromMem(fontAwesomeRegular, fontSize, xcc.FontStyle_Regular)
	xc.XC_SetTextRenderingHint(TextRenderingHintAntiAliasGridFit)
	return p
}

// GetFont 返回FontAwesome炫彩字体句柄map.
//
// map的键分别为:
//   - fa-solid
//   - fa-brands
//   - fa-regular
func (e *Elementui) GetFont() map[string]int {
	return e.hFontAwesomeMap
}

// funcDrawEleMap 存放元素绘制事件
var funcDrawEleMap = map[string]widget.XE_PAINT1{
	"onDrawButton_Default":     onDrawButton_Default,
	"onDrawButton_Color":       onDrawButton_Color,
	"onDrawButton_Text":        onDrawButton_Text,
	"onDrawButton_Color_Plain": onDrawButton_Color_Plain,
	"onDrawEdit":               onDrawEdit,
}

// onDrawEle 元素绘制事件
func onDrawEle(hEle int, hDraw int, pbHandled *bool) int {
	funcDrawEle := xc.XC_GetProperty(hEle, "element-func-draw-ele")
	if f, ok := funcDrawEleMap[funcDrawEle]; ok {
		f(hEle, hDraw, pbHandled)
	}
	return 0
}

// todo 用go来写显示所有图标, 根据json分类, 可搜索, 可复制名字, 十六进制, 十进制
