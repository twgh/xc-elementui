package eui

import (
	_ "embed"
	"encoding/json"
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

// todo 用go来写显示所有图标, 根据json分类, 可搜索, 可复制名字, 十六进制, 十进制

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

// CreateButton 创建按钮, 本函数会自己注册元素绘制事件进行绘制.
//
// x: 左边.
//
// y: 顶边.
//
// cx: 宽度. 大于0时 ButtonOption 按钮选项中的Size字段无效.
//
// cy: 高度. 大于0时 ButtonOption 按钮选项中的Size字段无效.
//
// text: 文本.
//
// hParent: 父元素或父窗口句柄.
//
// opts: ButtonOption 按钮选项, 可不填.
func (e *Elementui) CreateButton(x, y, cx, cy int32, text string, hParent int, opts ...ButtonOption) *Button {
	var opt ButtonOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Style < ButtonStyle_Default || opt.Style > ButtonStyle_Text {
		opt.Style = ButtonStyle_Default
	}
	if opt.Size < ButtonSize_Default || opt.Size > ButtonSize_Mini {
		opt.Size = ButtonSize_Default
	}

	// cx, cy < 1时根据Size选择预设宽高
	var nWidth, nHeight int32
	if cx < 1 && cy < 1 {
		widths := []int32{98, 98, 80, 80}
		heights := []int32{40, 36, 32, 28}
		nWidth = widths[opt.Size-1]
		nHeight = heights[opt.Size-1]
	}

	// 创建按钮
	btn := &Button{opt: opt, hFontAwesomeMap: e.hFontAwesomeMap, dpi: e.dpi}
	if nWidth > 0 && nHeight > 0 {
		btn.SetHandle(xc.XBtn_Create(x, y, nWidth, nHeight, text, hParent))
	} else {
		btn.SetHandle(xc.XBtn_Create(x, y, cx, cy, text, hParent))
	}

	// 启用背景透明
	btn.EnableBkTransparent(true)
	// 设置是否圆角按钮
	btn.SetRound(opt.IsRound)
	// 设置是否圆形按钮
	btn.SetCircle(opt.IsCircle)
	// 设置是否朴素按钮
	btn.SetPlain(opt.IsPlain)
	// 设置样式
	btn.SetStyle(opt.Style)

	// 自定义炫彩svg句柄优先级最高, 其次是炫彩图片句柄, 再然后是iconFa
	if opt.HSvg > 0 && xc.XC_IsHXCGUI(opt.HSvg, xcc.XC_SVG) {
		btn.SetHSvg(opt.HSvg)
	} else if opt.HImage > 0 && xc.XC_IsHXCGUI(opt.HImage, xcc.XC_IMAGE_FRAME) {
		btn.SetHImage(opt.HImage)
	} else { // 确定iconFa图标和字体类型
		if opt.IconUnicode > 0 {
			btn.SetIconUnicode(opt.IconUnicode)
		} else if opt.IconHex != "" {
			btn.SetIconHex(opt.IconHex)
		} else if opt.IconName != "" {
			btn.SetIconName(opt.IconName)
		}
	}
	// 注册元素绘制事件
	btn.Event_PAINT1(onDrawEle)
	return btn
}

// ChangeButton 改变现有的按钮. 可配合界面设计器来使用, 设计器里放按钮, 然后在代码里调用改变样式.
//
// hBtn: 按钮句柄. 如果不是按钮句柄, 函数会返回nil.
//
// opts: ButtonOption 按钮选项, 可不填. 只有填写了其中的Size字段, 才会改变现有按钮的宽高.
func (e *Elementui) ChangeButton(hBtn int, opts ...ButtonOption) *Button {
	if xc.XC_GetObjectType(hBtn) != xcc.XC_BUTTON {
		return nil
	}
	var opt ButtonOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Style < ButtonStyle_Default || opt.Style > ButtonStyle_Text {
		opt.Style = ButtonStyle_Default
	}

	// 创建element按钮对象
	btn := &Button{opt: opt, hFontAwesomeMap: e.hFontAwesomeMap, dpi: e.dpi}
	btn.SetHandle(hBtn)
	// 正确填写Size时才改变按钮的宽高
	btn.SetSizeEle(opt.Size)

	// 启用背景透明
	btn.EnableBkTransparent(true)
	// 设置是否圆角按钮
	btn.SetRound(opt.IsRound)
	// 设置是否圆形按钮
	btn.SetCircle(opt.IsCircle)
	// 设置是否朴素按钮
	btn.SetPlain(opt.IsPlain)
	// 设置样式
	btn.SetStyle(opt.Style)

	// 自定义炫彩svg句柄优先级最高, 其次是炫彩图片句柄, 再然后是iconFa
	if opt.HSvg > 0 && xc.XC_IsHXCGUI(opt.HSvg, xcc.XC_SVG) {
		btn.SetHSvg(opt.HSvg)
	} else if opt.HImage > 0 && xc.XC_IsHXCGUI(opt.HImage, xcc.XC_IMAGE_FRAME) {
		btn.SetHImage(opt.HImage)
	} else { // 确定iconFa图标和字体类型
		if opt.IconUnicode > 0 {
			btn.SetIconUnicode(opt.IconUnicode)
		} else if opt.IconHex != "" {
			btn.SetIconHex(opt.IconHex)
		} else if opt.IconName != "" {
			btn.SetIconName(opt.IconName)
		}
	}

	// 注册元素绘制事件
	btn.Event_PAINT1(onDrawEle)
	return btn
}
