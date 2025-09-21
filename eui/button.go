package eui

import (
	"strconv"
	"strings"

	"github.com/twgh/xcgui/ani"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Button 是 Elementui 风格的按钮, 继承 widget.Button.
type Button struct {
	widget.Button
	objBase
}

// CreateButton 创建按钮.
//   - 内存注册了元素绘制事件.
//
// text: 文本.
//
// hParent: 父元素或父窗口句柄.
//
// opts: ButtonOption 按钮选项, 可不填.
func (e *Elementui) CreateButton(text string, hParent int, opts ...ButtonOption) *Button {
	return updateButton(e, false, text, hParent, 0, opts...)
}

// ChangeButton 改变现有的按钮.
//   - 可配合界面设计器来使用, 设计器里放按钮, 然后在代码里调用改变样式.
//
// hBtn: 按钮句柄. 如果不是按钮句柄, 函数会返回 nil.
//
// opts: ButtonOption 按钮选项, 可不填. 只有填写了其中的 Size 字段, 才会改变现有按钮的宽高.
func (e *Elementui) ChangeButton(hBtn int, opts ...ButtonOption) *Button {
	return updateButton(e, true, "", 0, hBtn, opts...)
}

// 修改按钮.
//
// isChange: true 是改变模式, false 是创建模式.
//
// text: 文本. [创建模式]
//
// hParent: 父元素或父窗口句柄. [创建模式]
//
// hBtn: 按钮句柄. 如果不是按钮句柄, 函数会返回 nil. [改变模式]
//
// opts: ButtonOption 按钮选项, 可不填.
func updateButton(e *Elementui, isChange bool, text string, hParent, hBtn int, opts ...ButtonOption) *Button {
	if isChange && xc.XC_GetObjectType(hBtn) != xcc.XC_BUTTON {
		return nil
	}
	var opt ButtonOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if opt.Style < ButtonStyle_Default || opt.Style > ButtonStyle_Text {
		opt.Style = ButtonStyle_Default
	}
	if !isChange && opt.Size < ButtonSize_Default || opt.Size > ButtonSize_Mini {
		opt.Size = ButtonSize_Default
	}

	// 创建按钮对象
	btn := &Button{}
	btn.hFontAwesomeMap = e.hFontAwesomeMap
	btn.dpi = e.dpi
	if !isChange {
		hBtn = xc.XBtn_Create(opt.X, opt.Y, opt.Width, opt.Height, text, hParent)
	}
	btn.SetHandle(hBtn)
	btn.H = btn.Handle
	if isChange {
		// 正确填写 Size 时才改变宽高
		btn.SetSizeEle(opt.Size)
	} else {
		// 设置大小
		if opt.Width < 1 && opt.Height < 1 {
			btn.SetSizeEle(opt.Size)
		}
	}

	// 启用背景透明
	btn.EnableBkTransparent(true)
	// 设置圆角大小
	btn.SetRound(4)
	// 设置是否圆形按钮
	btn.EnableCircle(opt.IsCircle)
	// 设置是否朴素按钮
	btn.EnablePlain(opt.IsPlain)
	// 设置样式
	btn.SetStyle(opt.Style)

	// 自定义炫彩 svg 句柄优先级最高, 其次是炫彩图片句柄, 再然后是 iconFa
	if opt.HSvg > 0 && xc.XC_IsHXCGUI(opt.HSvg, xcc.XC_SVG) {
		btn.SetHSvg(opt.HSvg)
	} else if opt.HImage > 0 && xc.XC_IsHXCGUI(opt.HImage, xcc.XC_IMAGE_FRAME) {
		btn.SetHImage(opt.HImage)
	} else { // 确定 iconFa 图标和字体类型
		if opt.IconUnicode > 0 {
			btn.SetIconUnicode(opt.IconUnicode)
		} else if opt.IconHex != "" {
			btn.SetIconHex(opt.IconHex)
		} else if opt.Icon != "" {
			btn.SetIconName(opt.Icon)
		}
	}

	// 注册元素绘制事件
	btn.Event_PAINT1(onDrawEle)
	return btn
}

// SetLoading 启用或关闭加载中状态, 开启后会显示加载中图标, 同时按钮会禁止点击, 内部已自动重绘按钮.
//
// on: 是否启用.
//
// svgSize: 图标大小, 小于 1 时默认为 20.
//
// text: 同时更改加载中按钮的文本, on 参数为 true 时生效，加载状态结束后自动恢复原文本, 如果为空则不会更改按钮文本.
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
			// 设置 svg 大小
			if svgSize < 1 {
				svgSize = 20
			}
			xc.XSvg_SetSize(hSvg_loading, svgSize, svgSize)
			// 记录旧 svg 图标，设置加载中 svg 图标
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
					// 还原 svg 图标
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
// size: 预设好的大小, 可使用常量: ButtonSize_.
//   - 1 = default (98x40)
//   - 2 = medium (98x36)
//   - 3 = small (80x32)
//   - 4 = mini (80x28)
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
//   - 0 = default
//   - 1 = primary
//   - 2 = success
//   - 3 = info
//   - 4 = warning
//   - 5 = danger
//   - 6 = text
func (b *Button) SetStyle(style int) *Button {
	// 选择不同的绘制事件
	var funcDrawEle, bgColors, textColors, borderColors, textcolor string
	if b.IsPlain() && style != ButtonStyle_Text {
		if style == ButtonStyle_Default {
			funcDrawEle = "onDrawButton_Default"
		} else {
			funcDrawEle = "onDrawButton_Color_Plain"
			bgColors = ButtonBgColors_Plain[style]
			textColors = ButtonTextColors_Plain[style]
			borderColors = ButtonBorderColors_Plain[style]
			b.SetProperty("element-text-colors", textColors)
			b.SetProperty("element-border-colors", borderColors)
		}
	} else if !b.IsPlain() && style != ButtonStyle_Text {
		if style == ButtonStyle_Default {
			funcDrawEle = "onDrawButton_Default"
		} else {
			funcDrawEle = "onDrawButton_Color"
			textcolor = "4294967295"
			b.SetProperty("element-text-color", textcolor)
		}
		bgColors = ButtonBgColors[style]
	} else { // 无边框无背景
		funcDrawEle = "onDrawButton_Text"
	}
	b.SetProperty("element-func-draw-ele", funcDrawEle)
	b.SetProperty("element-bg-colors", bgColors)
	return b
}

// SetRound 设置按钮的圆角大小, 没有设置时的默认圆角是 4.
//
// round: 圆角大小, 小于 1 时为直角.
func (b *Button) SetRound(round int32) *Button {
	if round < 0 {
		round = 0
	}
	b.SetProperty("element-round", xc.Itoa(round*b.dpi/96))
	return b
}

// GetRound 获取按钮的圆角大小.
func (b *Button) GetRound() int32 {
	return xc.Atoi(b.GetProperty("element-round")) * 96 / b.dpi
}

// EnableCircle 设置按钮是否圆形.
//
// isCircle: 是否圆形按钮.
func (b *Button) EnableCircle(isCircle bool) *Button {
	b.SetProperty("element-circle", common.BoolToString(isCircle))
	return b
}

// EnablePlain 设置按钮是否朴素.
//
// isPlain: 是否朴素按钮.
func (b *Button) EnablePlain(isPlain bool) *Button {
	b.SetProperty("element-plain", common.BoolToString(isPlain))
	return b
}

// IsCircle 是否为圆形按钮.
func (b *Button) IsCircle() bool {
	return b.GetProperty("element-circle") == "true"
}

// IsPlain 是否为朴素按钮.
func (b *Button) IsPlain() bool {
	return b.GetProperty("element-plain") == "true"
}

// ButtonOption 按钮选项.
type ButtonOption struct {
	// 自定义炫彩 svg 句柄.
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, Icon 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > Icon.
	HSvg int
	// 自定义炫彩图片句柄.
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, Icon 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > Icon.
	HImage int

	// Font Wesome 图标对应的 Unicode 码点十进制数字, 如 61872 相当于'fa-solid fa-paw'.
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, Icon 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > Icon.
	IconUnicode int32
	// Font Wesome 图标对应的 Unicode 码点十六进制文本, 如'f1b0'相当于'fa-solid fa-paw'.
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, Icon 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > Icon.
	IconHex string
	// Font Wesome 图标名.
	//  - 如'fa-solid fa-paw', 前面是风格, 后面是图标名, 用空格分开, 其中风格可省略, 没有风格时会自动根据'fa-solid', 'fa-brands', 'fa-regular'的顺序尝试添加风格.
	//  - 图标大全: https://fa6.dashgame.com, 在网页里点导航栏图标, 然后点免费, 可筛选出 2000+ 免费图标, 点击图标会复制完整风格+图标名到剪贴板, 可直接使用. 内置 FontAwesome 版本为6.6.0
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, Icon 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > Icon.
	Icon string

	X, Y, Width, Height int32

	// 按钮尺寸, 默认为 ButtonSize_Default, 可使用常量: ButtonSize_
	//  - 使用预设的尺寸效果会比较好.
	//  - 如果 Width 或 Height 字段 > 0 那么本字段就无效.
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

	// 是否为朴素按钮, 默认为 false.
	//  - 当 Style 字段 = ButtonStyle_Text 时本字段无效.
	IsPlain bool
	// 是否为圆形按钮, 默认为 false.
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

// ButtonBgColors 存放按钮不同样式的背景颜色字符串, 不包含朴素按钮的.
//   - 顺序: Leave, Stay, Down, Check, Disable
var ButtonBgColors = map[int]string{
	ButtonStyle_Default: JoinColorString(xc.RGBA(255, 255, 255, 255), xc.RGBA(236, 245, 255, 255), xc.RGBA(236, 245, 255, 255), 0, xc.RGBA(255, 255, 255, 255)),
	ButtonStyle_Primary: JoinColorString(xc.RGBA(64, 158, 255, 255), xc.RGBA(102, 177, 255, 255), xc.RGBA(58, 142, 230, 255), 0, xc.RGBA(160, 207, 255, 255)),
	ButtonStyle_Success: JoinColorString(xc.RGBA(103, 194, 58, 255), xc.RGBA(133, 206, 97, 255), xc.RGBA(93, 175, 52, 255), 0, xc.RGBA(179, 225, 157, 255)),
	ButtonStyle_Info:    JoinColorString(xc.RGBA(144, 147, 153, 255), xc.RGBA(166, 169, 173, 255), xc.RGBA(130, 132, 138, 255), 0, xc.RGBA(200, 201, 204, 255)),
	ButtonStyle_Warning: JoinColorString(xc.RGBA(230, 162, 60, 255), xc.RGBA(235, 181, 99, 255), xc.RGBA(207, 146, 54, 255), 0, xc.RGBA(243, 209, 158, 255)),
	ButtonStyle_Danger:  JoinColorString(xc.RGBA(245, 108, 108, 255), xc.RGBA(247, 137, 137, 255), xc.RGBA(221, 97, 97, 255), 0, xc.RGBA(250, 182, 182, 255)),
}

// ButtonBorderColors_Plain 存放朴素按钮不同样式的边框颜色字符串, 不包含 default 样式的
//   - 顺序: Leave, Stay, Down, Check, Disable
var ButtonBorderColors_Plain = map[int]string{
	ButtonStyle_Primary: JoinColorString(xc.RGBA(179, 216, 255, 255), xc.RGBA(64, 158, 255, 255), xc.RGBA(58, 142, 230, 255), 0, xc.RGBA(217, 236, 255, 255)),
	ButtonStyle_Success: JoinColorString(xc.RGBA(194, 231, 176, 255), xc.RGBA(103, 194, 58, 255), xc.RGBA(93, 175, 52, 255), 0, xc.RGBA(225, 243, 216, 255)),
	ButtonStyle_Info:    JoinColorString(xc.RGBA(211, 212, 214, 255), xc.RGBA(144, 147, 153, 255), xc.RGBA(130, 132, 138, 255), 0, xc.RGBA(233, 233, 235, 255)),
	ButtonStyle_Warning: JoinColorString(xc.RGBA(245, 218, 177, 255), xc.RGBA(230, 162, 60, 255), xc.RGBA(207, 146, 54, 255), 0, xc.RGBA(250, 236, 216, 255)),
	ButtonStyle_Danger:  JoinColorString(xc.RGBA(251, 196, 196, 255), xc.RGBA(245, 108, 108, 255), xc.RGBA(221, 97, 97, 255), 0, xc.RGBA(253, 226, 226, 255)),
}

// ButtonBgColors_Plain 存放朴素按钮不同样式的背景颜色字符串, 不包含 default 样式的
//   - 顺序: Leave, Stay, Down, Check, Disable
var ButtonBgColors_Plain = map[int]string{
	ButtonStyle_Primary: JoinColorString(xc.RGBA(236, 245, 255, 255), xc.RGBA(64, 158, 255, 255), xc.RGBA(58, 142, 230, 255), 0, xc.RGBA(236, 245, 255, 255)),
	ButtonStyle_Success: JoinColorString(xc.RGBA(240, 249, 235, 255), xc.RGBA(103, 194, 58, 255), xc.RGBA(93, 175, 52, 255), 0, xc.RGBA(240, 249, 235, 255)),
	ButtonStyle_Info:    JoinColorString(xc.RGBA(244, 244, 245, 255), xc.RGBA(144, 147, 153, 255), xc.RGBA(130, 132, 138, 255), 0, xc.RGBA(244, 244, 245, 255)),
	ButtonStyle_Warning: JoinColorString(xc.RGBA(253, 246, 236, 255), xc.RGBA(230, 162, 60, 255), xc.RGBA(207, 146, 54, 255), 0, xc.RGBA(253, 246, 236, 255)),
	ButtonStyle_Danger:  JoinColorString(xc.RGBA(254, 240, 240, 255), xc.RGBA(245, 108, 108, 255), xc.RGBA(221, 97, 97, 255), 0, xc.RGBA(254, 240, 240, 255)),
}

// ButtonTextColors_Plain 存放朴素按钮不同样式的字体颜色字符串, 不包含 default 样式的
//   - 顺序: Leave, Stay, Down, Check, Disable
var ButtonTextColors_Plain = map[int]string{
	ButtonStyle_Primary: JoinColorString(xc.RGBA(64, 158, 255, 255), xc.RGBA(255, 255, 255, 255), xc.RGBA(255, 255, 255, 255), 0, xc.RGBA(140, 197, 255, 255)),
	ButtonStyle_Success: JoinColorString(xc.RGBA(103, 194, 58, 255), xc.RGBA(255, 255, 255, 255), xc.RGBA(255, 255, 255, 255), 0, xc.RGBA(164, 218, 137, 255)),
	ButtonStyle_Info:    JoinColorString(xc.RGBA(144, 147, 153, 255), xc.RGBA(255, 255, 255, 255), xc.RGBA(255, 255, 255, 255), 0, xc.RGBA(188, 190, 194, 255)),
	ButtonStyle_Warning: JoinColorString(xc.RGBA(230, 162, 60, 255), xc.RGBA(255, 255, 255, 255), xc.RGBA(255, 255, 255, 255), 0, xc.RGBA(240, 199, 138, 255)),
	ButtonStyle_Danger:  JoinColorString(xc.RGBA(245, 108, 108, 255), xc.RGBA(255, 255, 255, 255), xc.RGBA(255, 255, 255, 255), 0, xc.RGBA(249, 167, 167, 255)),
}

// 默认按钮和朴素默认按钮 style 0
func onDrawButton_Default(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	var rc xc.RECT
	rc.Right = xc.XEle_GetWidth(hEle)
	rc.Bottom = xc.XEle_GetHeight(hEle)
	xc.XDraw_EnableSmoothingMode(hDraw, true)

	var textColor, borderColor, bgColor uint32
	nState := xc.XBtn_GetStateEx(hEle)
	isPlain := xc.XC_GetProperty(hEle, "element-plain") == "true"
	if isPlain { // 朴素按钮
		bgColor = xcc.COLOR_WHITE
	} else {
		if bgColorsText := xc.XC_GetProperty(hEle, "element-bg-colors"); bgColorsText != "" {
			colors := strings.Split(bgColorsText, ",")
			if int(nState) < len(colors) {
				bgColor = common.AtoUint32(colors[nState])
			}
		}
	}

	switch nState {
	case xcc.Button_State_Leave:
		borderColor = xc.RGBA(220, 223, 229, 255)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = xc.OffsetRect(rc3, svgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = xc.OffsetRect(rc3, imgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = xc.OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
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

	textColor := common.AtoUint32(xc.XC_GetProperty(hEle, "element-text-color"))
	var bgColor uint32 // 背景颜色
	if bgColorsText := xc.XC_GetProperty(hEle, "element-bg-colors"); bgColorsText != "" {
		colors := strings.Split(bgColorsText, ",")
		nState := int(xc.XBtn_GetStateEx(hEle))
		if nState < len(colors) {
			bgColor = common.AtoUint32(colors[nState])
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = xc.OffsetRect(rc3, svgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = xc.OffsetRect(rc3, imgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = xc.OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
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
	var bgColor, textColor, borderColor uint32
	if bgColorsText := xc.XC_GetProperty(hEle, "element-bg-colors"); bgColorsText != "" {
		colors := strings.Split(bgColorsText, ",")
		if int(nState) < len(colors) {
			bgColor = common.AtoUint32(colors[nState])
		}
	}
	if textColorsText := xc.XC_GetProperty(hEle, "element-text-colors"); textColorsText != "" {
		colors := strings.Split(textColorsText, ",")
		if int(nState) < len(colors) {
			textColor = common.AtoUint32(colors[nState])
		}
	}
	if borderColorsText := xc.XC_GetProperty(hEle, "element-border-colors"); borderColorsText != "" {
		colors := strings.Split(borderColorsText, ",")
		if int(nState) < len(colors) {
			borderColor = common.AtoUint32(colors[nState])
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = xc.OffsetRect(rc3, svgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = xc.OffsetRect(rc3, imgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = xc.OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
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

	var textColor uint32 // 文本颜色
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-svgWidth-space)/2, 0, svgWidth, 0)
		xc.XDraw_DrawSvg(hDraw, hSvg, rc3.Left, (rc.Bottom-xc.XSvg_GetHeight(hSvg))/2)

		rc3 = xc.OffsetRect(rc3, svgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left-defaultFontShowSize.CX-imgWidth-space)/2, 0, imgWidth, 0)
		xc.XDraw_Image(hDraw, hImage, rc3.Left, (rc.Bottom-xc.XImage_GetHeight(hImage))/2)

		rc3 = xc.OffsetRect(rc3, imgWidth+space, 0, 0, 0)
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
		rc3 := xc.OffsetRect(rc, (rc.Right-rc.Left)/2-(defaultFontShowSize.CX+hFontAwesomeShowSizeCx)/2, 0, hFontAwesomeShowSizeCx, 0)
		xc.XDraw_DrawText(hDraw, iconFa, &rc3)

		rc3 = xc.OffsetRect(rc3, hFontAwesomeShowSizeCx, 0, 0, 0)
		xc.XDraw_SetFont(hDraw, defaultFont)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap)
		xc.XDraw_DrawText(hDraw, btnText, &rc3)
	} else if btnText != "" { // 纯文本
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)
		xc.XDraw_DrawText(hDraw, btnText, &rc)
	}
	return 0
}
