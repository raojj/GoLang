package main

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"
)

type Window interface {
	ShowWindow() // 展示窗体界面
}

type ComWindow struct {
	// 创建压缩解压缩界面类
	Window
	*walk.MainWindow // 主窗体
}

type LabWindow struct {
	// 创建展示提示信息的类
	Window
}

func Show(WindowType string) {
	// 创建界面类对象
	var Win Window
	switch WindowType {
	case "main_window":
		var comWindow ComWindow = ComWindow{
			Window:     nil,
			MainWindow: &walk.MainWindow{},
		}
		Win = &comWindow
	case "lab_window":
		Win = &LabWindow{}
	default:
		fmt.Println("参数传入错误")
	}
	Win.ShowWindow()
}

// ShowWindow 首先实现ShowWindow方法，展示空白的窗口
func (comWindow *ComWindow) ShowWindow() {
	pathWindow := new(ComWindow)
	var mainWindow = declarative.MainWindow{
		AssignTo: &pathWindow.MainWindow,                    // 关联主窗体
		Title:    "文件压缩",                                    // 窗口的标题名称
		MinSize:  declarative.Size{Width: 480, Height: 230}, // 设置窗口的大小，注意Size来自declarative包
	}
	err := mainWindow.Create()
	if err != nil {
		return
	}
	// 窗口的展示需要指定坐标
	err = pathWindow.SetX(650) // x坐标
	if err != nil {
		return
	}
	err = pathWindow.SetY(300) // y坐标
	if err != nil {
		return
	}
	pathWindow.Run() // 运行窗口

}
