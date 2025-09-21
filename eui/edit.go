package eui

import (
	"strconv"

	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// Edit 是 Elementui 风格的编辑框, 继承 widget.Edit.
type Edit struct {
	widget.Edit
	objBase
}

// CreateEdit 创建编辑框.
//   - 内部设置了边框大小, 这个影响的是文本的位置. 如果觉得不合适可自行调用 SetBorderSize 进行设置.
//   - 无图标时左右边框大小都是 15.
//   - 左边图标时, 左边框大小是 29, 右边框大小是 15.
//   - 右边图标时, 左边框大小是 15, 右边框大小是 29.
//   - 内部注册了元素绘制事件, 鼠标进入/离开事件, 编辑框光标位置改变事件.
//
// hParent: 父元素或父窗口句柄.
//
// opts: EditOption 编辑框选项, 可不填.
func (e *Elementui) CreateEdit(hParent int, opts ...EditOption) *Edit {
	return updateEdit(e, false, hParent, 0, opts...)
}

// ChangeEdit 改变现有的编辑框.
//   - 可配合界面设计器来使用, 设计器里放编辑框, 然后在代码里调用改变样式.
//
// hEdit: 编辑框句柄. 如果不是编辑框句柄, 函数会返回 nil.
//
// opts: EditOption 编辑框选项, 可不填. 只有填写了其中的 Size 字段, 才会改变现有编辑框的宽高.
func (e *Elementui) ChangeEdit(hEdit int, opts ...EditOption) *Edit {
	return updateEdit(e, true, 0, hEdit, opts...)
}

// 修改编辑框.
//
// isChange: true 是改变模式, false 是创建模式.
//
// hParent: 父元素或父窗口句柄. [创建模式]
//
// hEdit: 编辑框句柄. 如果不是编辑框句柄, 函数会返回 nil. [改变模式]
//
// opts: ButtonOption 按钮选项, 可不填.
func updateEdit(e *Elementui, isChange bool, hParent, hEdit int, opts ...EditOption) *Edit {
	if isChange && xc.XC_GetObjectType(hEdit) != xcc.XC_EDIT {
		return nil
	}
	var opt EditOption
	if len(opts) > 0 {
		opt = opts[0]
	}
	if !isChange && opt.Size < EditSize_Default || opt.Size > EditSize_Mini {
		opt.Size = EditSize_Default
	}

	// 创建编辑框对象
	edit := &Edit{}
	edit.hFontAwesomeMap = e.hFontAwesomeMap
	edit.dpi = e.dpi
	if !isChange {
		hEdit = xc.XEdit_Create(opt.X, opt.Y, opt.Width, opt.Height, hParent)
	}
	edit.SetHandle(hEdit)
	edit.H = edit.Handle
	if isChange {
		// 正确填写 Size 时才改变宽高
		edit.SetSizeEle(opt.Size)
	} else {
		// 设置大小
		if opt.Width < 1 && opt.Height < 1 {
			edit.SetSizeEle(opt.Size)
		}
	}

	// 设置圆角大小
	edit.SetRound(4)
	// 左边填充
	edit.SetProperty("element-space-left", xc.Itoa(4*e.dpi/96))

	// 启用背景透明
	edit.EnableBkTransparent(true)
	// 置默认文本
	if opt.DefaultText != "" {
		edit.SetDefaultText(opt.DefaultText)
	}
	// 置默认文本颜色
	edit.SetDefaultTextColor(xc.RGBA(193, 196, 203, 255))
	// 置插入符颜色
	edit.SetCaretColor(xc.RGBA(96, 98, 102, 255))
	// 置文本颜色
	edit.SetTextColor(xc.RGBA(96, 98, 102, 255))

	hasIcon := true // 是否有图标
	// 自定义炫彩 svg 句柄优先级最高, 其次是炫彩图片句柄, 再然后是 iconFa
	if opt.HSvg > 0 && xc.XC_IsHXCGUI(opt.HSvg, xcc.XC_SVG) {
		edit.SetHSvg(opt.HSvg)
	} else if opt.HImage > 0 && xc.XC_IsHXCGUI(opt.HImage, xcc.XC_IMAGE_FRAME) {
		edit.SetHImage(opt.HImage)
	} else { // 确定 iconFa 图标和字体类型
		if opt.IconUnicode > 0 {
			edit.SetIconUnicode(opt.IconUnicode)
		} else if opt.IconHex != "" {
			edit.SetIconHex(opt.IconHex)
		} else if opt.Icon != "" {
			edit.SetIconName(opt.Icon)
		} else { // 无图标
			hasIcon = false
		}
	}

	if hasIcon { // 有图标
		if opt.IsRight {
			edit.SetBorderSize(15, 0, 29, 0)
		} else {
			edit.SetBorderSize(29, 0, 15, 0)
		}
		edit.EnableRight(opt.IsRight)
		edit.EnableAutoColor(opt.IsAutoColor)
	} else {
		edit.SetBorderSize(15, 0, 15, 0)
	}

	edit.SetProperty("element-func-draw-ele", "onDrawEdit")

	// 注册元素鼠标进入事件
	edit.Event_MOUSESTAY1(onMouseStayEle)
	// 注册元素鼠标离开事件
	edit.Event_MOUSELEAVE1(onMouseLeaveEle)
	// 注册元素绘制事件
	edit.Event_PAINT1(onDrawEle)
	// 注册编辑框光标位置改变事件, 用于在光标移动时重绘
	edit.Event_EDIT_POS_CHANGED1(onEditPosChanged)
	return edit
}

// SetSizeEle 设置编辑框的大小. 只能使用预设好的常量.
//
// size: 预设好的大小, 可使用常量: EditSize_.
//   - 1 = default (180x40)
//   - 2 = medium (180x36)
//   - 3 = small (180x32)
//   - 4 = mini (180x28)
func (e *Edit) SetSizeEle(size int) *Edit {
	if size >= EditSize_Default && size <= EditSize_Mini {
		heights := []int32{40, 36, 32, 28}
		nHeight := heights[size-1]
		e.SetSize(180, nHeight, false, xcc.AdjustLayout_All, 0)
	}
	return e
}

// SetRound 设置编辑框的圆角大小, 没有设置时的默认圆角是 4.
//
// round: 圆角大小, 小于 1 时为直角.
func (e *Edit) SetRound(round int32) *Edit {
	if round < 0 {
		round = 0
	}
	e.SetProperty("element-round", xc.Itoa(round*e.dpi/96))
	return e
}

// GetRound 获取编辑框的圆角大小.
func (e *Edit) GetRound() int32 {
	return xc.Atoi(e.GetProperty("element-round")) * 96 / e.dpi
}

// EnableRight 设置编辑框的图标是否在右边.
//
// isRight: 图标是否在右边.
func (e *Edit) EnableRight(isRight bool) *Edit {
	e.SetProperty("element-icon-right", common.BoolToString(isRight))
	return e
}

// IsRight 判断编辑框的图标是否在右边.
func (e *Edit) IsRight() bool {
	return e.GetProperty("element-icon-right") == "true"
}

// EnableAutoColor 设置编辑框的图标颜色是否根据焦点颜色自动改变.
//
// isAutoColor: 图标颜色是否根据焦点颜色自动改变.
func (e *Edit) EnableAutoColor(isAutoColor bool) *Edit {
	e.SetProperty("element-icon-autocolor", common.BoolToString(isAutoColor))
	return e
}

// IsAutoColor 判断编辑框的图标颜色是否根据焦点颜色自动改变.
func (e *Edit) IsAutoColor() bool {
	return e.GetProperty("element-icon-autocolor") == "true"
}

// EditOption 编辑框选项.
type EditOption struct {
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
	//  - 图标大全: https://fa6.dashgame.com, 在网页里点导航栏图标, 然后点免费, 可筛选出 2000+ 免费图标, 点击图标会复制完整风格+图标名到剪贴板, 可直接使用. 内置 FontAwesome 版本为 6.6.0
	// 	- 注意: HSvg, HImage, IconUnicode, IconHex, Icon 这几个参数只需要填一个即可, 填多个的话, 生效顺序优先级为: HSvg > HImage > IconUnicode > IconHex > Icon.
	Icon string

	// 当无内容时显示的文本.
	DefaultText string

	X, Y, Width, Height int32

	// 编辑框尺寸, 默认为 EditSize_Default, 可使用常量: EditSize_
	//  - 使用预设的尺寸效果会比较好.
	//  - 如果 Width 或 Height 字段 > 0 那么本字段就无效.
	//  - 1 = default (180x40)
	//  - 2 = medium (180x36)
	//  - 3 = small (180x32)
	//  - 4 = mini (180x28)
	Size int

	// 图标是否在右边.
	IsRight bool
	// 图标颜色是否根据焦点颜色自动改变.
	//  - 如果你使用的是 HImage, 则此参数无效.
	IsAutoColor bool
}

// https://element.eleme.cn/2.14/#/zh-CN/component/input
// todo: Clearable 可清空的输入框
// todo: IsPassword 可切换显示隐藏的密码框
// todo: Textarea 多行输入框
// todo: 复合型输入框
// todo: 带输入建议

// 编辑框尺寸. 已经预设好的.

const (
	EditSize_Default = iota + 1 // 180x40
	EditSize_Mdeium             // 180x36
	EditSize_Small              // 180x32
	EditSize_Mini               // 180x28
)

// 编辑框绘制事件.
func onDrawEdit(hEle int, hDraw int, pbHandled *bool) int {
	*pbHandled = true
	eleWidth := xc.XEle_GetWidth(hEle)
	eleHeight := xc.XEle_GetHeight(hEle)
	var rc xc.RECT
	rc.Right = eleWidth
	rc.Bottom = eleHeight
	xc.XDraw_EnableSmoothingMode(hDraw, true)

	nState := xc.Atoi(xc.XC_GetProperty(hEle, "element-mouse-state"))
	var borderColor uint32
	if xc.XEle_IsFocus(hEle) { // 判断是否拥有焦点改变边框颜色和文本颜色
		borderColor = xc.RGBA(64, 158, 255, 255)
		if xc.XEle_GetStateFlags(hEle) == xcc.Element_State_Flag_Down {
			xc.XEle_SetTextColor(hEle, xc.RGBA(193, 196, 203, 255))
		} else {
			xc.XEle_SetTextColor(hEle, xc.RGBA(96, 98, 102, 255))
		}
	} else if nState == 0 {
		borderColor = xc.RGBA(220, 223, 229, 255)
	} else if nState == 1 {
		borderColor = xc.RGBA(193, 196, 203, 255)
	}
	bgColor := xcc.COLOR_WHITE

	if !xc.XEle_IsEnable(hEle) { // 元素为禁用状态改变各种颜色
		borderColor = xc.RGBA(228, 231, 237, 255)
		bgColor = xc.RGBA(245, 247, 250, 255)
		xc.XEle_SetTextColor(hEle, xc.RGBA(193, 196, 203, 255))
	} else {
		xc.XEle_SetTextColor(hEle, xc.RGBA(96, 98, 102, 255))
	}

	round := xc.Atoi(xc.XC_GetProperty(hEle, "element-round"))
	// 绘制圆角矩形边框
	xc.XDraw_SetBrushColor(hDraw, borderColor)
	xc.XDraw_DrawRoundRect(hDraw, &rc, round, round)

	// 绘制填充圆角矩形
	xc.XDraw_SetBrushColor(hDraw, bgColor)
	rc.Top = 1
	rc.Left = 1
	rc.Right = rc.Right - 1
	rc.Bottom = rc.Bottom - 1
	xc.XDraw_FillRoundRect(hDraw, &rc, round, round)

	IsRight := xc.XC_GetProperty(hEle, "element-icon-right") == "true"
	AutoColor := xc.XC_GetProperty(hEle, "element-icon-autocolor") == "true"
	spaceLeft := round
	if spaceLeft < 1 {
		spaceLeft = xc.Atoi(xc.XC_GetProperty(hEle, "element-space-left"))
	}

	if hSvg, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-hsvg")); hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		if AutoColor { // 图标颜色是否根据焦点颜色自动改变.
			xc.XSvg_SetUserFillColor(hSvg, borderColor, true)
		} else {
			xc.XSvg_SetUserFillColor(hSvg, xc.RGBA(193, 196, 203, 255), true)
		}

		rc.Top = (eleHeight - xc.XSvg_GetHeight(hSvg)) / 2
		if IsRight { // 图标是否在右边.
			svgWidth := xc.XSvg_GetWidth(hSvg)
			rc.Left = eleWidth - 1 - spaceLeft - svgWidth
			xc.XDraw_DrawSvg(hDraw, hSvg, rc.Left, rc.Top)
		} else {
			rc.Left += spaceLeft
			xc.XDraw_DrawSvg(hDraw, hSvg, rc.Left, rc.Top)
		}
	} else if hImage, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-icon-himage")); hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		rc.Top = (eleHeight - xc.XImage_GetHeight(hImage)) / 2
		if IsRight { // 图标是否在右边.
			imageWidth := xc.XImage_GetWidth(hImage)
			rc.Left = eleWidth - 1 - spaceLeft - imageWidth
			xc.XDraw_Image(hDraw, hImage, rc.Left, rc.Top)
		} else {
			rc.Left += spaceLeft
			xc.XDraw_Image(hDraw, hImage, rc.Left, rc.Top)
		}
	} else if iconFa := xc.XC_GetProperty(hEle, "element-icon-fa"); iconFa != "" {
		hFontAwesome, _ := strconv.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome"))
		xc.XDraw_SetFont(hDraw, hFontAwesome)
		xc.XDraw_SetTextAlign(hDraw, xcc.TextAlignFlag_Vcenter|xcc.TextFormatFlag_NoWrap|xcc.TextAlignFlag_Center)

		if AutoColor { // 图标颜色是否根据焦点颜色自动改变.
			xc.XDraw_SetBrushColor(hDraw, borderColor)
		} else {
			xc.XDraw_SetBrushColor(hDraw, xc.RGBA(193, 196, 203, 255))
		}

		hFontAwesomeShowSizeCx := xc.Atoi(xc.XC_GetProperty(hEle, "element-hfontawesome-showsize-cx"))
		if IsRight { // 图标是否在右边.
			rc.Left = eleWidth - 1 - spaceLeft - hFontAwesomeShowSizeCx
			rc.Right = rc.Left + hFontAwesomeShowSizeCx
			xc.XDraw_DrawText(hDraw, iconFa, &rc)
		} else {
			rc.Left += spaceLeft
			rc.Right = rc.Left + hFontAwesomeShowSizeCx
			xc.XDraw_DrawText(hDraw, iconFa, &rc)
		}
		defaultFont := xc.XC_GetDefaultFont()
		xc.XDraw_SetFont(hDraw, defaultFont)
	}
	return 0
}

// 元素鼠标进入事件
func onMouseStayEle(hEle int, pbHandled *bool) int {
	xc.XC_SetProperty(hEle, "element-mouse-state", "1")
	xc.XEle_Redraw(hEle, false)
	return 0
}

// 元素鼠标离开事件
func onMouseLeaveEle(hEle int, hEleStay int, pbHandled *bool) int {
	xc.XC_SetProperty(hEle, "element-mouse-state", "0")
	xc.XEle_Redraw(hEle, false)
	return 0
}

// 编辑框光标位置改变事件
func onEditPosChanged(hEle int, iPos int32, pbHandled *bool) int {
	xc.XEle_Redraw(hEle, false)
	return 0
}
