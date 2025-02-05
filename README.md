<h1 align="center">xc-elementui</h1>
<p align="center">
    <a href="https://github.com/twgh/xc-elementui/releases"><img src="https://img.shields.io/badge/release-早期版本-blue" alt="release"></a>
    <a href="https://github.com/twgh/xcgui"><img src="https://img.shields.io/badge/xcgui-1.3.394-blue" alt="xcgui"></a>
   <a href="https://golang.org"> <img src="https://img.shields.io/badge/golang-≥1.16-blue" alt="golang"></a>
    <a href="https://pkg.go.dev/github.com/twgh/xc-elementui"><img src="https://img.shields.io/badge/go.dev-reference-brightgreen" alt="GoDoc"></a>
    <a href="https://github.com/twgh/xc-elementui/blob/main/LICENSE"><img src="https://img.shields.io/badge/License-MIT-brightgreen" alt="License"></a>
    <br><br>
    <a href="https://github.com/twgh/xc-elementui/tree/main/example">程序示例</a>&nbsp;&nbsp;
	<a href="https://pkg.go.dev/github.com/twgh/xc-elementui">项目文档</a>&nbsp;&nbsp;
</p>







## 介绍

- 使用 [xcgui](https://github.com/twgh/xcgui) 仿 [Elementui](https://element.eleme.cn/#/zh-CN/component/installation)，元素样式具有一致性，快速开发界面。
- 内置 2000+ [Font Awesome](https://fa6.dashgame.com/) 免费图标，风格一致，开箱即用。

## 获取

```bash
go get -u github.com/twgh/xc-elementui@latest
```

## 动态链接库下载

程序运行时需要把 `xcgui.dll` 放到程序运行目录。

在开发时最好是放到 `C:\Windows\System32` 目录，这样就不需要频繁把dll放到不同程序的运行目录了。

#### 文件直链

| 64位 | [点击下载](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui.dll?version=latest) |
| ---- | ------------------------------------------------------------ |
| 32位 | [点击下载](https://pkggo-generic.pkg.coding.net/xcgui/file/xcgui-32.dll?version=latest) |


## 按钮例子

[![按钮例子](https://gitee.com/twgh/xc-elementui/raw/main/example/AllButton/1.png)](https://github.com/twgh/xc-elementui/tree/main/example/AllButton)

## 简单用法

```go
package main

import (
	"github.com/twgh/xc-elementui/eui"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/font"
	"github.com/twgh/xcgui/imagex"
	"github.com/twgh/xcgui/svg"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
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
	w = window.New(0, 0, 500, 400, "xc-elementui 简单例子", 0, xcc.Window_Style_Default|xcc.Window_Style_Drag_Window)
	// 设置窗口边框大小
	w.SetBorderSize(0, 32, 0, 0)
	// 设置窗口阴影, 圆角
	w.SetTransparentType(xcc.Window_Transparent_Shadow).SetShadowInfo(8, 255, 10, false, 0).SetTransparentAlpha(255)
	// 窗口启用布局, 水平垂直居中, 自动换行
	w.EnableLayout(true).SetSpace(10).SetSpaceRow(10).SetAlignH(xcc.Layout_Align_Center).SetAlignV(xcc.Layout_Align_Center).EnableAutoWrap(true)

	// 创建Elementui对象
	e = eui.NewElementui(12, w.GetDPI())

	// 默认按钮
	e.CreateButton(0, 0, 0, 0, "默认按钮", w.Handle)
	// 主要按钮
	e.CreateButton(0, 0, 0, 0, "主要按钮", w.Handle, eui.ButtonOption{Style: eui.ButtonStyle_Primary})
	// 图标+文字按钮 成功按钮
	e.CreateButton(0, 0, 0, 0, "图标文字", w.Handle, eui.ButtonOption{IconName: "fa-house-medical-flag", Style: eui.ButtonStyle_Success})
	// 图标+文字按钮 信息按钮 朴素按钮
	e.CreateButton(0, 0, 0, 0, "朴素按钮", w.Handle, eui.ButtonOption{IconName: "fa-jet-fighter", Style: eui.ButtonStyle_Info, IsPlain: true})
	// 图标+文字按钮 禁用状态 警告按钮
	e.CreateButton(0, 0, 0, 0, "禁用状态", w.Handle, eui.ButtonOption{IconName: "fa-wpexplorer", Style: eui.ButtonStyle_Warning}).Enable(false)
	// 图标+文字 圆角按钮 警告按钮
	e.CreateButton(0, 0, 0, 0, "圆角按钮", w.Handle, eui.ButtonOption{IconName: "fa-circle-radiation", IsRound: true, Style: eui.ButtonStyle_Warning})
	// 图标按钮 危险按钮
	e.CreateButton(0, 0, 56, 40, "", w.Handle, eui.ButtonOption{IconName: "fa-volcano", Style: eui.ButtonStyle_Danger})
	// 圆形图标按钮 主要按钮
	e.CreateButton(0, 0, 40, 40, "", w.Handle, eui.ButtonOption{IconName: "fa-text-height", IsCircle: true, Style: eui.ButtonStyle_Primary})

	// 图标按钮 默认按钮 自定义svg图标
	svgElement := svg.NewByStringW(svg_element).SetSize(20, 20)
	btn0 := e.CreateButton(0, 0, 0, 0, "点我加载", w.Handle, eui.ButtonOption{HSvg: svgElement.Handle, Style: eui.ButtonStyle_Primary})
	btn0.Event_BnClick1(func(hEle int, pbHandled *bool) int {
		btn0.SetLoading(true, 0, "")
		go func() {
			time.Sleep(2 * time.Second)
			/*
				在这里做一些加载数据的操作, 比如读取数据库数据
			*/
			xc.XC_CallUT(func() {
				/*
					拿到数据库数据后, 如果要赋予ui元素数据, 要在这里操作ui元素
				*/
				btn0.SetLoading(false, 0, "")
			})
		}()
		return 0
	})

	w.Show(true)
	a.Run()
	a.Exit()
}

const svg_element = `<svg t="1731392844936" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5912" width="24" height="24"><path d="M903.84 705.78c-0.23 37.18-19.54 45.46-19.54 45.46S551.44 943.85 529.8 955.9c-21.43 9.22-35.79 0-35.79 0S145.74 753.64 133.03 744.75C120.31 735.86 120 722 120 722s0.35-400.51 0-419.07c-0.36-18.53 22.77-32.47 22.77-32.47l348-201.44c21.43-11.32 42.27 0 42.27 0s307.45 178.96 341.5 198.18c33.42 15.89 29.29 48.71 29.29 48.71s0.2 355.3 0.01 389.87z m-138.96-402c-71.26-41.08-239.11-138.46-239.11-138.46s-16.39-8.87-33.21 0L219.33 322.95s-18.15 10.92-17.89 25.42c0.28 14.52 0 327.98 0 327.98s0.24 10.85 10.22 17.8c9.99 6.96 283.45 165.24 283.45 165.24s11.26 7.22 28.09 0c17-9.43 278.34-160.17 278.34-160.17s15.15-6.48 15.31-35.6c0.07-8.37 0.1-40.96 0.1-81.86L509.24 768.42V697c0-29.33 22.67-48.69 22.67-48.69l272.17-164.2c10.27-10.76 12.39-27.94 12.81-34.45v-72.39L509.24 563.88v-74.65c0-29.36 19.45-42.21 19.45-42.21l236.2-143.26v0.02z m0 0" fill="#FDDD48" p-id="5913"></path></svg>`
```


## 其它组件预期效果

![预期效果](https://gitee.com/twgh/image/raw/master/xc-elementui/yuqi.jpg)

## 开发进度

- [x] 按钮
- [x] 输入框
- [ ] 单选框
- [ ] 多选框
- [ ] 开关按钮
- [ ] 文字链接
- [ ] 选择器
- [ ] 时间选择器
- [ ] 日期选择器
- [ ] 滑块
- [ ] 计数器
- [ ] 颜色选择器
- [ ] 标签
- [ ] 进度条
- [ ] 导航菜单
- [ ] 标签页
- [ ] 步骤条
- [ ] 文字提示
- [ ] 弹出框
- [ ] 卡片
- [ ] 折叠面板
- [ ] ...

