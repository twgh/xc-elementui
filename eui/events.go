package eui

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
)

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
