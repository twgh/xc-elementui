package eui

import (
	"strconv"
	"strings"

	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

// objBase 对象基类.
type objBase struct {
	hFontAwesomeMap map[string]int // FontAwesome字体句柄
	H               int            // 句柄.
	dpi             int32          // 窗口dpi
}

// ClearIcon 清除掉已设置的 Font Awesome 图标, hsvg 和 himage.
func (o *objBase) ClearIcon() *objBase {
	xc.XC_SetProperty(o.H, "element-icon-hsvg", "")
	xc.XC_SetProperty(o.H, "element-icon-himage", "")
	xc.XC_SetProperty(o.H, "element-icon-fa", "")
	return o
}

// SetHSvg 设置炫彩 svg 句柄.
//
// hSvg: 炫彩 svg 句柄.
func (o *objBase) SetHSvg(hSvg int) *objBase {
	o.ClearIcon()
	if hSvg > 0 && xc.XC_IsHXCGUI(hSvg, xcc.XC_SVG) {
		xc.XC_SetProperty(o.H, "element-icon-hsvg", strconv.Itoa(hSvg))
	}
	return o
}

// GetHSvg 获取已设置的炫彩 svg 句柄.
func (o *objBase) GetHSvg() int {
	hSvg, _ := strconv.Atoi(xc.XC_GetProperty(o.H, "element-icon-hsvg"))
	return hSvg
}

// SetHImage 设置炫彩图片句柄.
//
// hImage: 炫彩图片句柄.
func (o *objBase) SetHImage(hImage int) *objBase {
	o.ClearIcon()
	if hImage > 0 && xc.XC_IsHXCGUI(hImage, xcc.XC_IMAGE_FRAME) {
		xc.XC_SetProperty(o.H, "element-icon-himage", strconv.Itoa(hImage))
	}
	return o
}

// GetHImage 获取已设置的炫彩图片句柄.
func (o *objBase) GetHImage() int {
	hImage, _ := strconv.Atoi(xc.XC_GetProperty(o.H, "element-icon-himage"))
	return hImage
}

// SetIconName 设置 Font Wesome 图标名.
//
// iconName: Font Awesome 图标名.
//   - 如'fa-solid fa-paw', 前面是风格, 后面是图标名, 用空格分开, 其中风格可省略, 没有风格时会自动根据'fa-solid', 'fa-brands', 'fa-regular'的顺序尝试添加风格.
//   - 图标大全: https://fa6.dashgame.com, 在网页里点导航栏图标, 然后点免费, 可筛选出 2000+ 免费图标, 点击图标会复制完整风格+图标名到剪贴板, 可直接使用. 内置 FontAwesome 版本为 6.6.0
func (o *objBase) SetIconName(iconName string) *objBase {
	// 删首尾空
	iconName = strings.TrimSpace(iconName)
	var iconFaStr, fontType string
	// 判断 IconName 是否存在, 如不存在就尝试加上所有风格前缀, 有就使用
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
	setIconFa(o, iconFaStr, fontType)
	return o
}

// GetIconName 获取已设置的 Font Awesome 图标名.
func (o *objBase) GetIconName() string {
	return xc.XC_GetProperty(o.H, "element-icon-fa")
}

// SetIconHex 设置 Font Awesome 图标对应的 Unicode 码点十六进制文本.
//
// iconHex: Font Wesome 图标对应的 Unicode 码点十六进制文本, 如'f1b0'相当于'fa-solid fa-paw'.
func (o *objBase) SetIconHex(iconHex string) *objBase {
	iconInt, _ := strconv.ParseInt(iconHex, 16, 32)
	iconUnicode := int32(iconInt)
	iconFaStr := string(iconUnicode)
	var fontType string
	// 遍历 map 找出图标 name, 得到字体类型
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
	setIconFa(o, iconFaStr, fontType)
	return o
}

// SetIconUnicode 设置 Font Awesome 图标对应的 Unicode 码点十进制数字.
//
// iconUnicode: Font Awesome 图标对应的 Unicode 码点十进制数字, 如 61872 相当于'fa-solid fa-paw'.
func (o *objBase) SetIconUnicode(iconUnicode int32) *objBase {
	iconFaStr := string(iconUnicode)
	var fontType string
	// 遍历 map 找出图标 name, 得到字体类型
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
	setIconFa(o, iconFaStr, fontType)
	return o
}

// 设置 iconfa 的相关信息.
//
// iconFaStr: Font Awesome 图标字符串.
//
// fontType: 字体类型, 可为'fa-solid', 'fa-brands', 'fa-regular'.
func setIconFa(o *objBase, iconFaStr, fontType string) *objBase {
	o.ClearIcon()
	xc.XC_SetProperty(o.H, "element-icon-fa", iconFaStr)
	if iconFaStr != "" { // 确定字体句柄和字体显示大小
		hFontAwesome := o.hFontAwesomeMap[fontType]
		xc.XC_SetProperty(o.H, "element-hfontawesome", strconv.Itoa(hFontAwesome))
		var hFontAwesomeShowSize xc.SIZE
		xc.XC_GetTextShowSize(iconFaStr, 1, hFontAwesome, &hFontAwesomeShowSize)
		xc.XC_SetProperty(o.H, "element-hfontawesome-showsize-cx", xc.Itoa(hFontAwesomeShowSize.CX))
	}
	return o
}
