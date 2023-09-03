package kpwidgets

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
)

func InitMainWindow() *widgets.QMainWindow {
	// Create the main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowIcon(gui.NewQIcon5("Ext\\Icons_04_CB\\Finals2\\plockb.ico"))
	window.SetWindowTitle("KeePass")

	icon := gui.NewQIcon5("Ext\\Icons_04_CB\\Finals2\\plockb.ico")
	widgets.QApplication_SetWindowIcon(icon)

	// Create the menu bar
	initMenuBar(window)
	// Create the toolbar with a title
	InitToolbar(window)
	InitMainContent(window)

	// 创建状态栏
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	// 在状态栏中显示文本
	statusBar.ShowMessage("Ready", 0)

	// Show the main window
	window.Resize2(800, 650)

	sys := InitTrayIcon(window)
	//menu := getTrayContextMenu()
	trayContextMenu := NewTrayContextMenu()
	//设置菜单
	sys.SetContextMenu(trayContextMenu)
	sys.Show()

	return window
}

func InitTrayIcon(window *widgets.QMainWindow) *widgets.QSystemTrayIcon {
	//系统托盘
	sys := widgets.NewQSystemTrayIcon(nil)
	//设置托盘图标
	//sys.SetIcon(window.Style().StandardIcon(widgets.QStyle__SP_MediaPlay, nil, nil))

	sys.SetIcon(gui.NewQIcon5("Ext/Icons_15_VA/KeePass_Round/KeePass_Round_24.png"))
	sys.ConnectActivated(func(reason widgets.QSystemTrayIcon__ActivationReason) {
		//单击系统托盘
		if reason == widgets.QSystemTrayIcon__Trigger {
			window.Show()
		}
	})
	return sys
}

func initMenuBar(window *widgets.QMainWindow) {
	menuBar := window.MenuBar()

	InitFileMenu(menuBar, window)
	InitEditMenu(menuBar, window)
	InitViewMenu(menuBar, window)
	InitToolsMenu(menuBar, window)
	InitHelpMenu(menuBar, window)
}

func clearChildItems(item *widgets.QTreeWidgetItem) {
	for item.ChildCount() > 0 {
		item.TakeChild(0)
	}
}

func mkValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{Key: key, Value: gokeepasslib.V{Content: value}}
}

// Function to calculate the password complexity score
func calculatePasswordComplexity(password string) int {
	// Dummy implementation, replace with your own logic
	// Calculate the complexity based on the password strength criteria
	// Return a score between 0 and 100
	return len(password) * 10
}
