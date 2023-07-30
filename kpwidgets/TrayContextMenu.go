package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/widgets"
)

type TrayContextMenu struct {
	*widgets.QMenu
}

func NewTrayContextMenu() *TrayContextMenu {
	menu := &TrayContextMenu{
		QMenu: widgets.NewQMenu(nil),
	}

	exit := menu.AddAction("Exit")
	exit.ConnectTriggered(func(bool) {
		//app.Exit(0)
	})
	//添加分隔符
	menu.AddSeparator()
	help := menu.AddAction("help")
	//定义二级菜单
	menuChild := widgets.NewQMenu(nil)
	menuChild.AddAction("option")
	about := menuChild.AddAction("about")
	about.ConnectTriggered(func(bool) {
		//button := widgets.QMessageBox_Information(nil, "title", "text", widgets.QMessageBox__Ok, widgets.QMessageBox__Yes)
		fmt.Println("click me")
		//widgets.QMessageBox_Information(nil, "title", "text", widgets.QMessageBox__Ok, widgets.QMessageBox__Yes)
		widgets.NewQFileDialog2(nil, "打开", "d:", "*.txt").Show()
	})
	//设置子项
	help.SetMenu(menuChild)
	return menu
}
