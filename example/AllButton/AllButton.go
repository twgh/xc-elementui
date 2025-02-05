// 所有按钮例子
package main

import (
	"github.com/twgh/xc-elementui/eui"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"math/rand"
	"time"
)

var (
	a *app.App
	w *window.Window
	e *eui.Elementui
)

func main() {
	a = app.New(true)
	a.EnableAutoDPI(true).EnableDPI(true)
	// 设置默认字体
	a.SetDefaultFont(font.NewEX("微软雅黑", 10, xcc.FontStyle_Regular).Handle)
	// 设置默认窗口图标
	a.SetWindowIcon(imagex.NewBySvgStringW(svg_element).Handle)

	// 创建窗口
	w = window.New(0, 0, 740, 696, "xc-elementui 按钮例子", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 设置窗口边框
	w.SetBorderSize(0, 32, 0, 0)
	// 设置窗口阴影, 圆角
	w.SetTransparentType(xcc.Window_Transparent_Shadow).SetShadowInfo(8, 255, 10, false, 0).SetTransparentAlpha(255)
	// 窗口启用布局, 水平垂直居中, 行列间距10
	w.EnableLayout(true).SetSpace(10).SetSpaceRow(10).SetAlignH(xcc.Layout_Align_Center).SetAlignV(xcc.Layout_Align_Center)

	// 创建Elementui对象
	e = eui.NewElementui(12, w.GetDPI())
	// 创建按钮
	createButton(w.Handle)

	a.ShowAndRun(w.Handle)
	a.Exit()
}

// 创建按钮
func createButton(hParent int) {
	// 基础按钮
	{
		layout := widget.NewLayoutEle(0, 0, 0, 0, hParent)
		layout.EnableDrawBorder(true).SetBorderSize(1, 1, 1, 1)
		layout.SetWidth(704).SetHeight(514).SetPadding(10, 0, 0, 0)
		layout.SetSpace(10).SetSpaceRow(10).SetAlignV(xcc.Layout_Align_Center)
		{
			// 默认按钮
			e.CreateButton(0, 0, 0, 0, "默认按钮", layout.Handle)
			// 主要按钮
			e.CreateButton(0, 0, 0, 0, "主要按钮", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Primary})
			// 成功按钮
			e.CreateButton(0, 0, 0, 0, "成功按钮", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Success})
			// 信息按钮
			e.CreateButton(0, 0, 0, 0, "信息按钮", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Info})
			// 警告按钮
			e.CreateButton(0, 0, 0, 0, "警告按钮", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Warning})
			// 危险按钮
			e.CreateButton(0, 0, 0, 0, "危险按钮", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Danger})
		}

		{
			// 图标+文字按钮 默认按钮
			e.CreateButton(0, 0, 0, 0, "图标文字", layout.Handle, eui.ButtonOption{IconName: "fa-copyright"}).LayoutItem_EnableWrap(true)
			// 图标+文字按钮 主要按钮
			e.CreateButton(0, 0, 0, 0, "图标文字", layout.Handle, eui.ButtonOption{IconName: "fa-house-medical-flag", Style: eui.ButtonStyle_Primary})
			// 图标+文字按钮 成功按钮
			e.CreateButton(0, 0, 0, 0, "图标文字", layout.Handle, eui.ButtonOption{IconName: "fa-address-card", Style: eui.ButtonStyle_Success})
			// 图标+文字按钮 信息按钮
			e.CreateButton(0, 0, 0, 0, "图标文字", layout.Handle, eui.ButtonOption{IconName: "fa-jet-fighter", Style: eui.ButtonStyle_Info})
			// 图标+文字按钮 警告按钮
			e.CreateButton(0, 0, 0, 0, "图标文字", layout.Handle, eui.ButtonOption{IconName: "fa-chart-pie", Style: eui.ButtonStyle_Warning})
			// 图标+文字按钮 危险按钮
			e.CreateButton(0, 0, 0, 0, "图标文字", layout.Handle, eui.ButtonOption{IconName: "fa-road-circle-exclamation", Style: eui.ButtonStyle_Danger})
		}

		{
			// 图标+文字按钮 默认按钮 朴素按钮
			e.CreateButton(0, 0, 0, 0, "朴素按钮", layout.Handle, eui.ButtonOption{IconName: "fa-copyright", IsPlain: true}).LayoutItem_EnableWrap(true)
			// 图标+文字按钮 主要按钮 朴素按钮
			e.CreateButton(0, 0, 0, 0, "朴素按钮", layout.Handle, eui.ButtonOption{IconName: "fa-house-medical-flag", Style: eui.ButtonStyle_Primary, IsPlain: true})
			// 图标+文字按钮 成功按钮 朴素按钮
			e.CreateButton(0, 0, 0, 0, "朴素按钮", layout.Handle, eui.ButtonOption{IconName: "fa-address-card", Style: eui.ButtonStyle_Success, IsPlain: true})
			// 图标+文字按钮 信息按钮 朴素按钮
			e.CreateButton(0, 0, 0, 0, "朴素按钮", layout.Handle, eui.ButtonOption{IconName: "fa-jet-fighter", Style: eui.ButtonStyle_Info, IsPlain: true})
			// 图标+文字按钮 警告按钮 朴素按钮
			e.CreateButton(0, 0, 0, 0, "朴素按钮", layout.Handle, eui.ButtonOption{IconName: "fa-chart-pie", Style: eui.ButtonStyle_Warning, IsPlain: true})
			// 图标+文字按钮 危险按钮 朴素按钮
			e.CreateButton(0, 0, 0, 0, "朴素按钮", layout.Handle, eui.ButtonOption{IconName: "fa-road-circle-exclamation", Style: eui.ButtonStyle_Danger, IsPlain: true})
		}

		{
			// 图标+文字按钮 禁用状态 默认按钮
			e.CreateButton(0, 0, 0, 0, "禁用状态", layout.Handle, eui.ButtonOption{IconName: "fa-copy"}).Enable(false).LayoutItem_EnableWrap(true)
			// 图标+文字按钮 禁用状态 主要按钮
			e.CreateButton(0, 0, 0, 0, "禁用状态", layout.Handle, eui.ButtonOption{IconName: "fa-italic", Style: eui.ButtonStyle_Primary}).Enable(false)
			// 图标+文字按钮 禁用状态 成功按钮
			e.CreateButton(0, 0, 0, 0, "禁用状态", layout.Handle, eui.ButtonOption{IconName: "fa-location-dot", Style: eui.ButtonStyle_Success}).Enable(false)
			// 图标+文字按钮 禁用状态 信息按钮
			e.CreateButton(0, 0, 0, 0, "禁用状态", layout.Handle, eui.ButtonOption{IconName: "fa-invision", Style: eui.ButtonStyle_Info}).Enable(false)
			// 图标+文字按钮 禁用状态 警告按钮
			e.CreateButton(0, 0, 0, 0, "禁用状态", layout.Handle, eui.ButtonOption{IconName: "fa-wpexplorer", Style: eui.ButtonStyle_Warning}).Enable(false)
			// 图标+文字按钮 禁用状态 危险按钮
			e.CreateButton(0, 0, 0, 0, "禁用状态", layout.Handle, eui.ButtonOption{IconName: "fa-house-lock", Style: eui.ButtonStyle_Danger}).Enable(false)
		}

		{
			// 图标+文字 圆角按钮 默认按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IconName: "fa-file-export", IsRound: true}).LayoutItem_EnableWrap(true)
			// 图标+文字 圆角按钮 主要按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IconName: "fa-z", IsRound: true, Style: eui.ButtonStyle_Primary})
			// 图标+文字 圆角按钮 成功按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IconName: "fa-apple", IsRound: true, Style: eui.ButtonStyle_Success})
			// 图标+文字 圆角按钮 信息按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IconName: "fa-face-kiss-beam", IsRound: true, Style: eui.ButtonStyle_Info})
			// 图标+文字 圆角按钮 警告按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IconName: "fa-circle-radiation", IsRound: true, Style: eui.ButtonStyle_Warning})
			// 图标+文字 圆角按钮 危险按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IconName: "fa-wifi", IsRound: true, Style: eui.ButtonStyle_Danger})
		}

		{
			// 文字圆角按钮 默认按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IsRound: true}).LayoutItem_EnableWrap(true)
			// 文字圆角按钮 主要按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IsRound: true, Style: eui.ButtonStyle_Primary})
			// 文字圆角按钮 成功按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IsRound: true, Style: eui.ButtonStyle_Success})
			// 文字圆角按钮 信息按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IsRound: true, Style: eui.ButtonStyle_Info})
			// 文字圆角按钮 警告按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IsRound: true, Style: eui.ButtonStyle_Warning})
			// 文字圆角按钮 危险按钮
			e.CreateButton(0, 0, 0, 0, "圆角按钮", layout.Handle, eui.ButtonOption{IsRound: true, Style: eui.ButtonStyle_Danger})
		}

		{
			// 图标按钮 默认按钮
			e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-regular fa-copyright"}).LayoutItem_EnableWrap(true)
			// 图标按钮 主要按钮
			e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-buffer", Style: eui.ButtonStyle_Primary})
			// 图标按钮 成功按钮
			e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-building-un", Style: eui.ButtonStyle_Success})
			// 图标按钮 信息按钮
			e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-file-code", Style: eui.ButtonStyle_Info})
			// 图标按钮 警告按钮
			e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-hand-lizard", Style: eui.ButtonStyle_Warning})
			// 图标按钮 危险按钮
			e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-volcano", Style: eui.ButtonStyle_Danger})
		}

		{
			// 圆形图标按钮 默认按钮
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-magnifying-glass", IsCircle: true}).LayoutItem_EnableWrap(true)
			// 圆形图标按钮 主要按钮
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-text-height", IsCircle: true, Style: eui.ButtonStyle_Primary})
			// 圆形图标按钮 成功按钮
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-square-poll-vertical", IsCircle: true, Style: eui.ButtonStyle_Success})
			// 圆形图标按钮 信息按钮
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-arrow-right-to-bracket", IsCircle: true, Style: eui.ButtonStyle_Info})
			// 圆形图标按钮 警告按钮
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-pump-medical", IsCircle: true, Style: eui.ButtonStyle_Warning})
			// 圆形图标按钮 危险按钮
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{IconName: "fa-gear", IsCircle: true, Style: eui.ButtonStyle_Danger})
		}

		{ // 改变现有按钮
			// 普通按钮 改变为 图标+文字 圆角按钮 默认按钮
			btn := widget.NewButton(0, 0, 100, 30, "现有按钮", layout.Handle)
			e.ChangeButton(btn.Handle, eui.ButtonOption{Size: eui.ButtonSize_Default}).LayoutItem_EnableWrap(true)

			// 点击改变按钮样式
			btn1 := e.CreateButton(0, 0, 0, 0, "随机变样式", layout.Handle)
			rand.Seed(time.Now().UnixNano())
			btn1.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn1.SetStyle(rand.Intn(7))
				return 0
			})
		}

		{
			svgDel := svg.NewByStringW(svg_del).SetSize(20, 20)
			// 图标按钮 默认按钮 自定义svg图标
			btn0 := e.CreateButton(0, 0, 56, 40, "", layout.Handle, eui.ButtonOption{HSvg: svgDel.Handle})
			btn0.LayoutItem_EnableWrap(true)
			btn0.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn0.SetLoading(true, 0, "")
				go func() {
					time.Sleep(2 * time.Second)
					/*
						在这里做一些加载数据的操作, 比如读取数据库数据
					*/
					a.CallUT(func() {
						/*
							拿到数据库数据后, 如果要赋予ui元素数据, 要在这里操作ui元素
						*/
						btn0.SetLoading(false, 0, "")
					})
				}()
				return 0
			})

			// 圆形图标按钮 危险按钮 自定义svg图标
			btn4 := e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{HSvg: svgDel.Handle, Style: eui.ButtonStyle_Danger, IsCircle: true})
			btn4.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn4.SetLoading(true, 0, "")
				go func() {
					time.Sleep(2 * time.Second)
					xc.XC_CallUT(func() {
						btn4.SetLoading(false, 0, "")
					})
				}()
				return 0
			})

			// 图标+文字按钮 默认按钮 自定义svg图标
			btn1 := e.CreateButton(0, 0, 0, 0, "svg图标", layout.Handle, eui.ButtonOption{HSvg: svgDel.Handle})
			btn1.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn1.SetLoading(true, 0, "加载中")
				go func() {
					time.Sleep(2 * time.Second)
					xc.XC_CallUT(func() {
						btn1.SetLoading(false, 0, "")
					})
				}()
				return 0
			})

			// 图标+文字按钮 默认按钮 自定义炫彩图片句柄 这个图标颜色不会随按钮风格变化
			img1 := imagex.NewBySvgStringW(svg_del)
			btn2 := e.CreateButton(0, 0, 0, 0, "炫彩图片", layout.Handle, eui.ButtonOption{HImage: img1.Handle})
			btn2.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn2.SetLoading(true, 0, "加载中")
				go func() {
					time.Sleep(2 * time.Second)
					xc.XC_CallUT(func() {
						btn2.SetLoading(false, 0, "")
					})
				}()
				return 0
			})

			// 图标+文字按钮 警告按钮 自定义svg图标
			btn3 := e.CreateButton(0, 0, 0, 0, "点我加载", layout.Handle, eui.ButtonOption{HSvg: svgDel.Handle, Style: eui.ButtonStyle_Primary})
			btn3.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn3.SetLoading(true, 0, "加载中")
				go func() {
					time.Sleep(2 * time.Second)
					xc.XC_CallUT(func() {
						btn3.SetLoading(false, 0, "")
					})
				}()
				return 0
			})

			// 图标+文字按钮 警告按钮 朴素按钮 自定义svg图标
			btn5 := e.CreateButton(0, 0, 0, 0, "点我加载", layout.Handle, eui.ButtonOption{HSvg: svgDel.Handle, Style: eui.ButtonStyle_Warning, IsPlain: true})
			btn5.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn5.SetLoading(true, 0, "")
				go func() {
					time.Sleep(2 * time.Second)
					xc.XC_CallUT(func() {
						btn5.SetLoading(false, 0, "")
					})
				}()
				return 0
			})

			// 图标+文字按钮 无边框无背景按钮按钮 自定义svg图标
			btn6 := e.CreateButton(0, 0, 100, 40, "点我加载", layout.Handle, eui.ButtonOption{HSvg: svgDel.Handle, Style: eui.ButtonStyle_Text})
			btn6.Event_BnClick1(func(hEle int, pbHandled *bool) int {
				btn6.SetLoading(true, 0, "")
				go func() {
					time.Sleep(2 * time.Second)
					xc.XC_CallUT(func() {
						btn6.SetLoading(false, 0, "")
					})
				}()
				return 0
			})
		}
	}

	// 无边框无背景按钮
	{
		layout := widget.NewLayoutEle(0, 0, 0, 0, hParent)
		layout.EnableDrawBorder(true).SetBorderSize(1, 1, 1, 1)
		layout.SetWidth(704).SetHeight(100).SetPadding(10, 0, 0, 0)
		layout.SetSpace(10).SetSpaceRow(10).SetAlignV(xcc.Layout_Align_Center)

		{ // 正常状态
			// 文字按钮
			e.CreateButton(0, 0, 120, 40, "无边框无背景", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Text})
			// 图标+文字
			e.CreateButton(0, 0, 120, 40, "无边框无背景", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Text, IconName: "fa-leaf"})
			// 只有图标
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Text, IconName: "fa-paw"})
		}

		{ // 禁用状态
			// 文字按钮
			e.CreateButton(0, 0, 120, 40, "无边框无背景", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Text}).Enable(false).LayoutItem_EnableWrap(true)
			// 图标+文字
			e.CreateButton(0, 0, 120, 40, "无边框无背景", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Text, IconName: "fa-leaf"}).Enable(false)
			// 只有图标
			e.CreateButton(0, 0, 40, 40, "", layout.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Text, IconName: "fa-paw"}).Enable(false)
		}
	}
}

const (
	svg_del = `<svg t="1731387404919" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1120" width="24" height="24"><path d="M561.1 708.454V413.798c0-47.155 49.102-47.155 49.102 0v294.708c0.05 34.099-49.101 38.656-49.101-0.052z m-147.302 0V413.798c0-47.155 49.101-47.155 49.101 0v294.708c0 38.656-49.1 34.099-49.1-0.052z m442.01-442.01H708.454v-49.151c0-71.731-22.988-98.253-98.252-98.253H413.798c-74.035 0-98.252 24.166-98.252 98.253v49.152H168.192c-53.094 0-53.094 49.1 0 49.1h687.616c53.094 0 53.094-49.1 0-49.1z m-491.162-49.151c0-47.667 2.97-49.101 49.101-49.101h196.455c46.08 0 49.1 1.126 49.1 49.1v49.153H364.646v-49.152z m343.91 687.616h-393.01c-70.964 0-98.253-27.239-98.253-98.253V413.798c0-49.612 49.1-49.612 49.1 0v392.91c0 47.718-0.102 49.151 49.101 49.151h392.91c47.718 0 49.1 0.154 49.1-49.152V413.798c0-48.486 49.1-48.486 49.1 0v392.91c0.103 69.426-23.96 98.2-98.047 98.2z" fill="#333333" p-id="1121"></path></svg>`

	svg_element = `<svg t="1731392844936" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5912" width="24" height="24"><path d="M903.84 705.78c-0.23 37.18-19.54 45.46-19.54 45.46S551.44 943.85 529.8 955.9c-21.43 9.22-35.79 0-35.79 0S145.74 753.64 133.03 744.75C120.31 735.86 120 722 120 722s0.35-400.51 0-419.07c-0.36-18.53 22.77-32.47 22.77-32.47l348-201.44c21.43-11.32 42.27 0 42.27 0s307.45 178.96 341.5 198.18c33.42 15.89 29.29 48.71 29.29 48.71s0.2 355.3 0.01 389.87z m-138.96-402c-71.26-41.08-239.11-138.46-239.11-138.46s-16.39-8.87-33.21 0L219.33 322.95s-18.15 10.92-17.89 25.42c0.28 14.52 0 327.98 0 327.98s0.24 10.85 10.22 17.8c9.99 6.96 283.45 165.24 283.45 165.24s11.26 7.22 28.09 0c17-9.43 278.34-160.17 278.34-160.17s15.15-6.48 15.31-35.6c0.07-8.37 0.1-40.96 0.1-81.86L509.24 768.42V697c0-29.33 22.67-48.69 22.67-48.69l272.17-164.2c10.27-10.76 12.39-27.94 12.81-34.45v-72.39L509.24 563.88v-74.65c0-29.36 19.45-42.21 19.45-42.21l236.2-143.26v0.02z m0 0" fill="#FDDD48" p-id="5913"></path></svg>`
)
