package eui

import (
	"github.com/twgh/xcgui/ani"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"strconv"
	"strings"
)

// Button is elementui-Button, 继承widget.Button.
type Button struct {
	opt             ButtonOption
	hFontAwesomeMap map[string]int // FontAwesome字体句柄
	dpi             int32          // 窗口dpi
	widget.Button
}

// SetLoading 启用或关闭加载中状态, 开启后会显示加载中图标, 同时按钮会禁止点击, 内部已自动重绘按钮.
//
// on: 是否启用.
//
// svgSize: 图标大小, 小于1时默认为20.
//
// text: 同时更改加载中按钮的文本, on参数为true时生效，加载状态结束后自动恢复原文本, 如果为空则不会更改按钮文本.
func (b *Button) SetLoading(on bool, svgSize int32, text string) *Button {
	hAni, _ := strconv.Atoi(b.GetProperty("element-hani"))
	if on {
		b.Enable(!on)
		// 记录按钮旧文本, 设置新文本
		b.SetProperty("element-old-text", b.GetText())
		if text != "" {
			b.SetText(text)
		}
		// 防止重复创建动画
		if hAni > 0 && xc.XC_GetObjectType(hAni) == xcc.XC_ANIMATION_SEQUENCE {
			xc.XAnima_Release(hAni, true)
		}
		// 设置加载中图标
		hSvg_loading := xc.XSvg_LoadStringW(svg_loading)
		if hSvg_loading > 0 && xc.XC_IsHXCGUI(hSvg_loading, xcc.XC_SVG) {
			// 设置svg大小
			if svgSize < 1 {
				svgSize = 20
			}
			xc.XSvg_SetSize(hSvg_loading, svgSize, svgSize)
			// 记录旧svg图标，设置加载中svg图标
			b.SetProperty("element-icon-hsvg-old", b.GetProperty("element-icon-hsvg"))
			b.SetProperty("element-icon-hsvg", strconv.Itoa(hSvg_loading))
			// 创建动画序列
			ani1 := ani.NewAnima(hSvg_loading, 0)
			ani1.Rotate(2000, 360, 0, 0, false)
			ani1.Run(b.Handle)
			b.SetProperty("element-hani", strconv.Itoa(ani1.Handle))
		}
	} else { // 销毁动画序列
		if hAni > 0 && xc.XC_GetObjectType(hAni) == xcc.XC_ANIMATION_SEQUENCE {
			if xc.XAnima_Release(hAni, true) {
				b.SetProperty("element-hani", "")
				hSvg_loading, _ := strconv.Atoi(b.GetProperty("element-icon-hsvg"))
				if hSvg_loading > 0 && xc.XC_IsHXCGUI(hSvg_loading, xcc.XC_SVG) {
					xc.XSvg_Destroy(hSvg_loading)
					// 还原svg图标
					b.SetProperty("element-icon-hsvg", b.GetProperty("element-icon-hsvg-old"))
				}
			}
		}
		b.SetText(b.GetProperty("element-old-text"))
		b.Enable(!on)
	}
	b.Redraw(false)
	return b
}

// SetSizeEle 设置 Button 的大小. 只能使用预设好的常量.
//
// size: 预设好的按钮大小, 可使用常量: ButtonSize_.
func (b *Button) SetSizeEle(size int) *Button {
	if size >= ButtonSize_Default && size <= ButtonSize_Mini {
		widths := []int32{98, 98, 80, 80}
		heights := []int32{40, 36, 32, 28}
		nWidth := widths[size-1]
		nHeight := heights[size-1]
		b.SetSize(nWidth, nHeight, false, xcc.AdjustLayout_All, 0)
	}
	return b
}

// SetStyle 设置按钮样式.
//
// style: 按钮样式, 默认为 ButtonStyle_Default, 可使用常量: ButtonStyle_.
func (b *Button) SetStyle(style int) *Button {
	// 选择不同的绘制事件
	var funcDraw, bgColors, textColors, borderColors, textcolor string
	if b.IsPlain() && style != ButtonStyle_Text {
		if style == ButtonStyle_Default {
			funcDraw = "onDrawButton_Default"
		} else {
			funcDraw = "onDrawButton_Color_Plain"
			bgColors = bgColorsMap_Plain[style]
			textColors = textColorsMap_Plain[style]
			borderColors = borderColorsMap_Plain[style]
			b.SetProperty("element-text-colors", textColors)
			b.SetProperty("element-border-colors", borderColors)
		}
	} else if !b.IsPlain() && style != ButtonStyle_Text {
		if style == ButtonStyle_Default {
			funcDraw = "onDrawButton_Default"
		} else {
			funcDraw = "onDrawButton_Color"
			textcolor = "4294967295"
			b.SetProperty("element-text-color", textcolor)
		}
		bgColors = bgcolorsMap[style]
	} else { // 无边框无背景
		funcDraw = "onDrawButton_Text"
	}
	b.SetProperty("element-func-draw", funcDraw)
	b.SetProperty("element-bg-colors", bgColors)
	return b
}

// SetRound 设置按钮是否圆角.
//
// isRound: 是否圆角按钮.
func (b *Button) SetRound(isRound bool) *Button {
	if isRound {
		b.SetProperty("element-round", xc.Itoa(18*b.dpi/96))
	} else {
		b.SetProperty("element-round", xc.Itoa(4*b.dpi/96))
	}
	return b
}

// SetCircle 设置按钮是否圆形.
//
// isCircle: 是否圆形按钮.
func (b *Button) SetCircle(isCircle bool) *Button {
	if isCircle {
		b.SetProperty("element-circle", "true")
	} else {
		b.SetProperty("element-circle", "false")
	}
	return b
}

// SetPlain 设置按钮是否朴素.
//
// isPlain: 是否朴素按钮.
func (b *Button) SetPlain(isPlain bool) *Button {
	if isPlain {
		b.SetProperty("element-plain", "true")
	} else {
		b.SetProperty("element-plain", "false")
	}
	return b
}

// IsRound 是否为圆角按钮.
func (b *Button) IsRound() bool {
	return b.GetProperty("element-round") == "true"
}

// IsCircle 是否为圆形按钮.
func (b *Button) IsCircle() bool {
	return b.GetProperty("element-circle") == "true"
}

// IsPlain 是否为朴素按钮.
func (b *Button) IsPlain() bool {
	return b.GetProperty("element-plain") == "true"
}

// ClearIcon 清除掉已设置的Font Awesome 图标, hsvg和himage.
func (b *Button) ClearIcon() *Button {
	b.SetProperty("element-icon-hsvg", "")
	b.SetProperty("element-icon-himage", "")
	b.SetProperty("element-icon-fa", "")
	return b
}

// SetHSvg 设置炫彩svg句柄.
//
// hSvg: 炫彩svg句柄.
func (b *Button) SetHSvg(hSvg int) *Button {
	b.ClearIcon()
	if hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		b.SetProperty("element-icon-hsvg", strconv.Itoa(hSvg))
	}
	return b
}

// SetHImage 设置炫彩图片句柄.
//
// hImage: 炫彩图片句柄.
func (b *Button) SetHImage(hImage int) *Button {
	b.ClearIcon()
	if hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		b.SetProperty("element-icon-himage", strconv.Itoa(hImage))
	}
	return b
}

// setIconFa 设置iconfa的相关信息.
//
// iconFaStr: Font Awesome 图标字符串.
//
// fontType: 字体类型, 可为'fa-solid', 'fa-brands', 'fa-regular'.
func (b *Button) setIconFa(iconFaStr, fontType string) *Button {
	b.ClearIcon()
	b.SetProperty("element-icon-fa", iconFaStr)
	if iconFaStr != "" { // 确定字体句柄和字体显示大小
		hFontAwesome := b.hFontAwesomeMap[fontType]
		b.SetProperty("element-hfontawesome", strconv.Itoa(hFontAwesome))
		var hFontAwesomeShowSize xc.SIZE
		xc.XC_GetTextShowSize(iconFaStr, 1, hFontAwesome, &hFontAwesomeShowSize)
		b.SetProperty("element-hfontawesome-showsize-cx", xc.Itoa(hFontAwesomeShowSize.CX))
	}
	return b
}

// SetIconName 设置Font Wesome 图标名.
//
// iconName: Font Awesome 图标名.
//   - 如'fa-solid fa-paw', 前面是风格, 后面是图标名, 用空格分开, 其中风格可省略, 没有风格时会自动根据'fa-solid', 'fa-brands', 'fa-regular'的顺序尝试添加风格.
//   - 图标大全: https://fa6.dashgame.com, 在网页里点导航栏图标, 然后点免费, 可筛选出2000+免费图标, 点击图标会复制完整风格+图标名到剪贴板, 可直接使用. 内置FontAwesome版本为6.6.0
func (b *Button) SetIconName(iconName string) *Button {
	// 删首尾空
	iconName = strings.TrimSpace(iconName)
	var iconFaStr, fontType string
	// 判断IconName是否存在, 如不存在就尝试加上所有风格前缀, 有就使用
	var iconUnicode int32
	var ok bool
	iconName2 := iconName
	styles := []string{"fa-solid ", "fa-brands ", "fa-regular "}
	for i := -1; i < len(styles); i++ {
		if i > -1 {
			iconName2 = styles[i] + iconName
		}
		if iconUnicode, ok = fontAwesomemMap[iconName2]; ok {
			iconName = iconName2
			break
		}
	}
	if iconUnicode > 0 {
		iconFaStr = string(iconUnicode)
		// 得到字体类型
		fontType = iconName
		index := strings.Index(fontType, " ")
		if index != -1 {
			fontType = fontType[:index]
		}
	}
	b.setIconFa(iconFaStr, fontType)
	return b
}

// SetIconHex 设置Font Awesome 图标对应的Unicode码点十六进制文本.
//
// iconHex: Font Wesome 图标对应的Unicode码点十六进制文本, 如'f1b0'相当于'fa-solid fa-paw'.
func (b *Button) SetIconHex(iconHex string) *Button {
	iconInt, _ := strconv.ParseInt(iconHex, 16, 32)
	iconUnicode := int32(iconInt)
	iconFaStr := string(iconUnicode)
	var fontType string
	// 遍历map找出图标name, 得到字体类型
	for name, value := range fontAwesomemMap {
		if value == iconUnicode {
			fontType = name
			index := strings.Index(fontType, " ")
			if index != -1 {
				fontType = fontType[:index]
			}
			break
		}
	}
	b.setIconFa(iconFaStr, fontType)
	return b
}

// SetIconUnicode 设置Font Awesome 图标对应的Unicode码点十进制数字.
//
// iconUnicode: Font Awesome 图标对应的Unicode码点十进制数字, 如61872相当于'fa-solid fa-paw'.
func (b *Button) SetIconUnicode(iconUnicode int32) *Button {
	iconFaStr := string(iconUnicode)
	var fontType string
	// 遍历map找出图标name, 得到字体类型
	for name, value := range fontAwesomemMap {
		if value == iconUnicode {
			fontType = name
			index := strings.Index(fontType, " ")
			if index != -1 {
				fontType = fontType[:index]
			}
			break
		}
	}
	b.setIconFa(iconFaStr, fontType)
	return b
}

// ButtonOption 按钮选项.
type ButtonOption struct {
	// Font Wesome 图标名.
	//  - 如'fa-solid fa-paw', 前面是风格, 后面是图标名, 用空格分开, 其中风格可省略, 没有风格时会自动根据'fa-solid', 'fa-brands', 'fa-regular'的顺序尝试添加风格.
	//  - 图标大全: https://fa6.dashgame.com, 在网页里点导航栏图标, 然后点免费, 可筛选出2000+免费图标, 点击图标会复制完整风格+图标名到剪贴板, 可直接使用. 内置FontAwesome版本为6.6.0
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, IconName 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > IconName.
	IconName string
	// Font Wesome 图标对应的Unicode码点十六进制文本, 如'f1b0'相当于'fa-solid fa-paw'.
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, IconName 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > IconName.
	IconHex string
	// Font Wesome 图标对应的Unicode码点十进制数字, 如61872相当于'fa-solid fa-paw'.
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, IconName 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > IconName.
	IconUnicode int32

	// 自定义炫彩svg句柄, 优先级最高, 如果使用了这个那么 HImage, IconName, IconHex, IconUnicode 这几个参数就无效了.
	HSvg int
	// 自定义炫彩图片句柄, 优先级仅低于 HSvg 参数, 如果使用了这个那么 IconName, IconHex, IconUnicode 这几个参数就无效了.
	HImage int

	// 按钮尺寸, 默认为 ButtonSize_Default, 可使用常量: ButtonSize_
	//  - 使用预设的按钮尺寸效果会比较好.
	//  - 如果cx或cy参数 > 0那么本字段就无效.
	//  - 1 = default (98x40)
	//  - 2 = medium (98x36)
	//  - 3 = small (80x32)
	//  - 4 = mini (80x28)
	Size int

	// 按钮样式, 默认为 ButtonStyle_Default, 可使用常量: ButtonStyle_
	//  - 0 = default
	//  - 1 = primary
	//  - 2 = success
	//  - 3 = info
	//  - 4 = warning
	//  - 5 = danger
	//  - 6 = text
	Style int

	// 是否为朴素按钮, 默认为false.
	//  - 当 Style 字段 = ButtonStyle_Text 时本字段无效.
	IsPlain bool
	// 是否为圆角按钮, 默认为false.
	//  - 注意: IsRound 参数和 IsCircle 参数只能二选一, 要么是圆角, 要么是圆形, 圆形优先级高于圆角.
	//  - 当 Style 字段 = ButtonStyle_Text 时本字段无效.
	IsRound bool
	// 是否为圆形按钮, 默认为false.
	//  - 注意: IsRound 参数和 IsCircle 参数只能二选一, 要么是圆角, 要么是圆形, 圆形优先级高于圆角.
	//  - 当 Style 字段 = ButtonStyle_Text 时本字段无效.
	IsCircle bool
}

// 按钮尺寸. 已经预设好的.

const (
	ButtonSize_Default = iota + 1 // 98x40
	ButtonSize_Mdeium             // 98x36
	ButtonSize_Small              // 80x32
	ButtonSize_Mini               // 80x28
)

// 按钮样式. 已经预设好的.

const (
	ButtonStyle_Default = iota
	ButtonStyle_Primary
	ButtonStyle_Success
	ButtonStyle_Info
	ButtonStyle_Warning
	ButtonStyle_Danger
	ButtonStyle_Text
)

// funcDrawMap 存放元素绘制事件
var funcDrawMap = map[string]widget.XE_PAINT1{
	"onDrawButton_Default":     onDrawButton_Default,
	"onDrawButton_Color":       onDrawButton_Color,
	"onDrawButton_Text":        onDrawButton_Text,
	"onDrawButton_Color_Plain": onDrawButton_Color_Plain,
}

// bgcolorsMap 存放普通按钮的背景颜色字符串
var bgcolorsMap = map[int]string{
	ButtonStyle_Default: "4294967295, 4294964716, 4294964716, -1, 4294967295",
	ButtonStyle_Primary: "4294942272, 4294947174, 4293299770, -1, 4294954912",
	ButtonStyle_Success: "4282040935, 4284599941, 4281642845, -1, 4288537011",
	ButtonStyle_Info:    "4288254864, 4289571238, 4287267970, -1, 4291611080",
	ButtonStyle_Warning: "4282163942, 4284724715, 4281766607, -1, 4288598515",
	ButtonStyle_Danger:  "4285295861, 4287203831, 4284572125, -1, 4290164474",
}

// borderColorsMap_Plain 存放朴素按钮的边框颜色字符串, 不包含default样式的
var borderColorsMap_Plain = map[int]string{
	ButtonStyle_Primary: "4294957235, 4294942272, 4293299770, -1, 4294962393",
	ButtonStyle_Success: "4289783746, 4282040935, 4281642845, -1, 4292408289",
	ButtonStyle_Info:    "4292269267, 4288254864, 4287267970, -1, 4293650921",
	ButtonStyle_Warning: "4289846005, 4282163942, 4281766607, -1, 4292406522",
	ButtonStyle_Danger:  "4291085563, 4285295861, 4284572125, -1, 4293059325",
}

// bgColorsMap_Plain 存放朴素按钮的背景颜色字符串, 不包含default样式的
var bgColorsMap_Plain = map[int]string{
	ButtonStyle_Primary: "4294964716, 4294942272, 4293299770, -1, 4294964716",
	ButtonStyle_Success: "4293655024, 4282040935, 4281642845, -1, 4293655024",
	ButtonStyle_Info:    "4294309108, 4288254864, 4287267970, -1, 4294309108",
	ButtonStyle_Warning: "4293719805, 4282163942, 4281766607, -1, 4293719805",
	ButtonStyle_Danger:  "4293980414, 4285295861, 4284572125, -1, 4293980414",
}

// textColorsMap_Plain 存放朴素按钮的字体颜色字符串, 不包含default样式的
var textColorsMap_Plain = map[int]string{
	ButtonStyle_Primary: "4294942272, 4294967295, 4294967295, -1, 4294952332",
	ButtonStyle_Success: "4282040935, 4294967295, 4294967295, -1, 4287224484",
	ButtonStyle_Info:    "4288254864, 4294967295, 4294967295, -1, 4290952892",
	ButtonStyle_Warning: "4282163942, 4294967295, 4294967295, -1, 4287285232",
	ButtonStyle_Danger:  "4285295861, 4294967295, 4294967295, -1, 4289177593",
}

// onDrawEle 元素绘制事件
func onDrawEle(hEle int, hDraw int, pbHandled *bool) int {
	funcDraw := xc.XC_GetProperty(hEle, "element-func-draw")
	if f, ok := funcDrawMap[funcDraw]; ok {
		f(hEle, hDraw, pbHandled)
	}
	return 0
}

// 默认按钮和朴素默认按钮 style 0
func onDrawButton_Default(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	var rc xc.RECT
	rc.Right = xc.XEle_GetWidth(hEle)
	rc.Bottom = xc.XEle_GetHeight(hEle)
	xc.XDraw_EnableSmoothingMode(hDraw, true)

	var textColor, borderColor, bgColor int
	nState := xc.XBtn_GetStateEx(hEle)
	isPlain := xc.XC_GetProperty(hEle, "element-plain") == "true"
	if isPlain { // 朴素按钮
		bgColor = xcc.COLOR_WHITE
	} else {
		if bgColorsText := xc.XC_GetProperty(hEle, "element-bg-colors"); bgColorsText != "" {
			colors := strings.Split(bgColorsText, ", ")
			if int(nState) < len(colors) {
				bgColor, _ = strconv.Atoi(colors[nState])
			}
		}
	}

	switch nState {
	case xcc.Button_State_Leave:
		borderColor = xc.RGBA(220, 223, 230, 255)
		textColor = xc.RGBA(96, 98, 102, 255)
	case xcc.Button_State_Stay:
		borderColor = xc.RGBA(198, 226, 255, 255)
		textColor = xc.RGBA(64, 158, 255, 255)
	case xcc.Button_State_Down:
		borderColor = xc.RGBA(58, 142, 230, 255)
		textColor = xc.RGBA(58, 142, 230, 255)
	case xcc.Button_State_Disable:
		borderColor = xc.RGBA(235, 238, 245, 255)
		textColor = xc.RGBA(192, 196, 204, 255)
	}

	var rc2 xc.RECT
	round := xc.Atoi(xc.XC_GetProperty(hEle, "element-round"))
	if xc.XC_GetProperty(hEle, "element-circle") == "true" { // 圆形按钮
		xc.XDraw_SetBrushColor(hDraw, borderColor)
		xc.XDraw_DrawEllipse(hDraw, &rc)
		rc2.Top = 1
		rc2.Left = 1
		rc2.Right = rc.Right - 1
		rc2.Bottom = rc.Bottom - 1
		xc.XDraw_SetBrushColor(hDraw, bgColor)
		xc.XDraw_FillEllipse(hDraw, &rc2)
	} else { // 圆角按钮
		xc.XDraw_SetBrushColor(hDraw, borderColor)
		xc.XDraw_DrawRoundRect(hDraw, &rc, round, round)
		rc2.Top = 1
		rc2.Left = 1
		rc2.Right = rc.Right - 1
		rc2.Bottom = rc.Bottom - 1
		xc.XDraw_SetBrushColor(hDraw, bgColor)
		xc.XDraw_FillRoundRect(hDraw, &rc2, round, round)
	}
	xc.XDraw_SetBrushColor(hDraw, textColor)

	btnText := xc.XBtn_GetText(hEle)
	if hSvg, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-hsvg")); hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		xc.XSvg_SetUserFillColor(hSvg, textColor, true)
		if btnText == "" { // 只有图标
			xc.XDraw_DrawSvg(hDraw, hSvg, (rc.Right-xc.XSvg_GetWidth(hSvg))/2, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		svgWidth := xc.XSvg_GetWidth(hSvg)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = OffsetRect(rc3, svgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if hImage, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-himage")); hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		if btnText == "" { // 只有图标
			xc.XDraw_Image(hDraw, hImage, (rc.Right-xc.XImage_GetWidth(hImage))/2, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		imgWidth := xc.XImage_GetWidth(hImage)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = OffsetRect(rc3, imgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_SetBrushColor(hDraw, textColor)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if iconFa := xc.XC_GetProperty(hEle, "element-icon-fa"); iconFa != "" {
		hFontAwesome, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome"))
		xc.XDraw_SetFont(hDraw, hFontAwesome)
		if btnText == "" { // 只有图标
			xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
			xc.XDraw_DrawText(hDraw, iconFa, &rc)
			return 0
		}

		// 图标+文字
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		var defaultFontShowSize xc.SIZE
		defaultFont := xc.XC_GetDefaultFont()
		hFontAwesomeShowSizeCx := xc.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome-showsize-cx"))
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), defaultFont, &defaultFontShowSize)
		rc3 := OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
		xc.XDraw_SetFont(hDraw, defaultFont)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if btnText != "" { // 纯文本
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
		xc.XDraw_DrawText(hDraw, btnText, &rc)
	}
	return 0
}

// 彩色按钮 style 1-5
func onDrawButton_Color(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	var rc xc.RECT
	rc.Right = xc.XEle_GetWidth(hEle)
	rc.Bottom = xc.XEle_GetHeight(hEle)
	xc.XDraw_EnableSmoothingMode(hDraw, true)

	textColor, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-text-color"))
	var bgColor int // 背景颜色
	if bgColorsText := xc.XC_GetProperty(hEle, "element-bg-colors"); bgColorsText != "" {
		colors := strings.Split(bgColorsText, ", ")
		nState := int(xc.XBtn_GetStateEx(hEle))
		if nState < len(colors) {
			bgColor, _ = strconv.Atoi(colors[nState])
		}
	}

	round := xc.Atoi(xc.XC_GetProperty(hEle, "element-round"))
	if xc.XC_GetProperty(hEle, "element-circle") == "true" { // 圆形按钮
		xc.XDraw_SetBrushColor(hDraw, bgColor)
		xc.XDraw_FillEllipse(hDraw, &rc)
	} else { // 圆角按钮
		xc.XDraw_SetBrushColor(hDraw, bgColor)
		xc.XDraw_FillRoundRect(hDraw, &rc, round, round)
	}
	xc.XDraw_SetBrushColor(hDraw, textColor)

	btnText := xc.XBtn_GetText(hEle)
	if hSvg, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-hsvg")); hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		xc.XSvg_SetUserFillColor(hSvg, textColor, true)
		if btnText == "" { // 只有图标
			xc.XDraw_DrawSvg(hDraw, hSvg, (rc.Right-xc.XSvg_GetWidth(hSvg))/2, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		svgWidth := xc.XSvg_GetWidth(hSvg)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = OffsetRect(rc3, svgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if hImage, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-himage")); hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		if btnText == "" { // 只有图标
			xc.XDraw_Image(hDraw, hImage, (rc.Right-xc.XImage_GetWidth(hImage))/2, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		imgWidth := xc.XImage_GetWidth(hImage)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = OffsetRect(rc3, imgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_SetBrushColor(hDraw, textColor)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if iconFa := xc.XC_GetProperty(hEle, "element-icon-fa"); iconFa != "" {
		hFontAwesome, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome"))
		xc.XDraw_SetFont(hDraw, hFontAwesome)
		if btnText == "" { // 只有图标
			xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
			xc.XDraw_DrawText(hDraw, iconFa, &rc)
			return 0
		}

		// 图标+文字
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		var defaultFontShowSize xc.SIZE
		defaultFont := xc.XC_GetDefaultFont()
		hFontAwesomeShowSizeCx := xc.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome-showsize-cx"))
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), defaultFont, &defaultFontShowSize)
		rc3 := OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
		xc.XDraw_SetFont(hDraw, defaultFont)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if btnText != "" { // 纯文本
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
		xc.XDraw_DrawText(hDraw, btnText, &rc)
	}
	return 0
}

// 朴素彩色按钮 style 1-5
func onDrawButton_Color_Plain(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	var rc xc.RECT
	rc.Right = xc.XEle_GetWidth(hEle)
	rc.Bottom = xc.XEle_GetHeight(hEle)
	xc.XDraw_EnableSmoothingMode(hDraw, true)

	nState := xc.XBtn_GetStateEx(hEle)
	var bgColor, textColor, borderColor int
	if bgColorsText := xc.XC_GetProperty(hEle, "element-bg-colors"); bgColorsText != "" {
		colors := strings.Split(bgColorsText, ", ")
		if int(nState) < len(colors) {
			bgColor, _ = strconv.Atoi(colors[nState])
		}
	}
	if textColorsText := xc.XC_GetProperty(hEle, "element-text-colors"); textColorsText != "" {
		colors := strings.Split(textColorsText, ", ")
		if int(nState) < len(colors) {
			textColor, _ = strconv.Atoi(colors[nState])
		}
	}
	if borderColorsText := xc.XC_GetProperty(hEle, "element-border-colors"); borderColorsText != "" {
		colors := strings.Split(borderColorsText, ", ")
		if int(nState) < len(colors) {
			borderColor, _ = strconv.Atoi(colors[nState])
		}
	}

	var rc2 xc.RECT
	switch nState {
	case xcc.Button_State_Leave:
		rc2.Top = 1
		rc2.Left = 1
		rc2.Right = rc.Right - 1
		rc2.Bottom = rc.Bottom - 1
	case xcc.Button_State_Stay:
		rc2 = rc
	case xcc.Button_State_Down:
		rc2 = rc
	case xcc.Button_State_Disable:
		rc2.Top = 1
		rc2.Left = 1
		rc2.Right = rc.Right - 1
		rc2.Bottom = rc.Bottom - 1
	}

	round := xc.Atoi(xc.XC_GetProperty(hEle, "element-round"))
	if xc.XC_GetProperty(hEle, "element-circle") == "true" { // 圆形按钮
		xc.XDraw_SetBrushColor(hDraw, borderColor)
		xc.XDraw_DrawEllipse(hDraw, &rc)
		xc.XDraw_SetBrushColor(hDraw, bgColor)
		xc.XDraw_FillEllipse(hDraw, &rc)
	} else { // 圆角按钮
		xc.XDraw_SetBrushColor(hDraw, borderColor)
		xc.XDraw_DrawRoundRect(hDraw, &rc, round, round)
		xc.XDraw_SetBrushColor(hDraw, bgColor)
		xc.XDraw_FillRoundRect(hDraw, &rc2, round, round)
	}
	xc.XDraw_SetBrushColor(hDraw, textColor)

	btnText := xc.XBtn_GetText(hEle)
	if hSvg, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-hsvg")); hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		xc.XSvg_SetUserFillColor(hSvg, textColor, true)
		if btnText == "" { // 只有图标
			xc.XDraw_DrawSvg(hDraw, hSvg, (rc.Right-xc.XSvg_GetWidth(hSvg))/2, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		svgWidth := xc.XSvg_GetWidth(hSvg)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = OffsetRect(rc3, svgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if hImage, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-himage")); hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		if btnText == "" { // 只有图标
			xc.XDraw_Image(hDraw, hImage, (rc.Right-xc.XImage_GetWidth(hImage))/2, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		imgWidth := xc.XImage_GetWidth(hImage)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = OffsetRect(rc3, imgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_SetBrushColor(hDraw, textColor)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if iconFa := xc.XC_GetProperty(hEle, "element-icon-fa"); iconFa != "" {
		hFontAwesome, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome"))
		xc.XDraw_SetFont(hDraw, hFontAwesome)
		if btnText == "" { // 只有图标
			xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
			xc.XDraw_DrawText(hDraw, iconFa, &rc)
			return 0
		}

		// 图标+文字
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		var defaultFontShowSize xc.SIZE
		defaultFont := xc.XC_GetDefaultFont()
		hFontAwesomeShowSizeCx := xc.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome-showsize-cx"))
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), defaultFont, &defaultFontShowSize)
		rc3 := OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
		xc.XDraw_SetFont(hDraw, defaultFont)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if btnText != "" { // 纯文本
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
		xc.XDraw_DrawText(hDraw, btnText, &rc)
	}
	return 0
}

// 无边框无背景按钮 style 6
func onDrawButton_Text(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	var rc xc.RECT
	rc.Right = xc.XEle_GetWidth(hEle)
	rc.Bottom = xc.XEle_GetHeight(hEle)
	xc.XDraw_EnableSmoothingMode(hDraw, true)

	var textColor int // 文本颜色
	nState := xc.XBtn_GetStateEx(hEle)
	switch nState {
	case xcc.Button_State_Leave:
		textColor = xc.RGBA(64, 158, 255, 255)
	case xcc.Button_State_Stay:
		textColor = xc.RGBA(102, 177, 255, 255)
	case xcc.Button_State_Down:
		textColor = xc.RGBA(58, 142, 230, 255)
	case xcc.Button_State_Disable:
		textColor = xc.RGBA(192, 196, 204, 255)
	}
	xc.XDraw_SetBrushColor(hDraw, textColor)

	btnText := xc.XBtn_GetText(hEle)
	if hSvg, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-hsvg")); hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		xc.XSvg_SetUserFillColor(hSvg, textColor, true)
		if btnText == "" { // 只有图标
			xc.XDraw_DrawSvg(hDraw, hSvg, (rc.Right-xc.XSvg_GetWidth(hSvg))/2, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		svgWidth := xc.XSvg_GetWidth(hSvg)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = OffsetRect(rc3, svgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if hImage, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-himage")); hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		if btnText == "" { // 只有图标
			xc.XDraw_Image(hDraw, hImage, (rc.Right-xc.XImage_GetWidth(hImage))/2, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)
			return 0
		}

		// 图标+文字
		var defaultFontShowSize xc.SIZE
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), xc.XC_GetDefaultFont(), &defaultFontShowSize)
		imgWidth := xc.XImage_GetWidth(hImage)
		var space int32 = 4 // 图标和文字之间的间距
		rc3 := OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = OffsetRect(rc3, imgWidth+space, 0, 0, 0)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_SetBrushColor(hDraw, textColor)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if iconFa := xc.XC_GetProperty(hEle, "element-icon-fa"); iconFa != "" {
		hFontAwesome, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome"))
		xc.XDraw_SetFont(hDraw, hFontAwesome)
		if btnText == "" { // 只有图标
			xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
			xc.XDraw_DrawText(hDraw, iconFa, &rc)
			return 0
		}

		// 图标+文字
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		var defaultFontShowSize xc.SIZE
		defaultFont := xc.XC_GetDefaultFont()
		hFontAwesomeShowSizeCx := xc.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome-showsize-cx"))
		xc.XC_GetTextShowSize(btnText, int32(len(btnText)), defaultFont, &defaultFontShowSize)
		rc3 := OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
		xc.XDraw_SetFont(hDraw, defaultFont)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if btnText != "" { // 纯文本
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
		xc.XDraw_DrawText(hDraw, btnText, &rc)
	}
	return 0
}
