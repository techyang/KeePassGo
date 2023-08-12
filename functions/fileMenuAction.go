package functions

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strings"
)

var kpResources = NewKpResources()

func InitFileMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	fileMenu := menuBar.AddMenu2("File")
	core.QCoreApplication_Translate("myapp", "", "", -1)

	// Create actions for the file menu
	newAction := fileMenu.AddAction("&New...")
	//newActionIcon := gui.NewQIcon5("Ext/Icons_15_VA/KeePass_Round/KeePass_Round_24.png")
	newActionIcon := window.Style().StandardIcon(widgets.QStyle__SP_FileIcon, nil, nil)

	newAction.SetIcon(newActionIcon)
	//newAction.SetIcon(gui.NewQIcon5("D:\\workspace_finmall\\autotest\\metersphere-2.10.3-lts\\api-test\\frontend\\node_modules\\mobius1-selectr\\docs\\favicon.ico"))

	// Connect the actions and tool buttons to their respective triggered events
	newAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the file menu
	//openAction := fileMenu.AddAction("&Open")

	openMenu := fileMenu.AddMenu2("&Open")
	openRecentMenu := fileMenu.AddMenu2("&Open Recent")
	openFileAction := openMenu.AddAction("&Open File...")
	openUrlAction := openMenu.AddAction("&Open Url...")

	fileMenu.AddSeparator()
	saveAction := fileMenu.AddAction("&Save")
	saveAsMenu := fileMenu.AddMenu2("&Save As ...")
	saveAsMenu.AddAction("&Save To File ...")
	saveAsMenu.AddAction("&Save To Url ...")
	saveAsMenu.AddSeparator()
	saveAsMenu.AddAction("&Save Copy To File ...")
	saveAction.ConnectTriggered(func(checked bool) {
		msgBox := widgets.NewQMessageBox(window)
		msgBox.SetWindowTitle("退出确认")
		msgBox.SetText("是否退出?")
		msgBox.SetInformativeText("真的要退出吗?")
		msgBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
		msgBox.SetDefaultButton2(widgets.QMessageBox__Cancel)
		// 添加自定义按钮
		openButton := msgBox.AddButton2("打开", widgets.QMessageBox__ActionRole)
		openButton.ConnectClicked(func(checked bool) {
			// 在这里添加自定义按钮的逻辑
			fmt.Println("点击了打开按钮")
		})

		// 获取 "OK" 按钮和 "Cancel" 按钮
		okButton := msgBox.Button(widgets.QMessageBox__Ok)
		cancelButton := msgBox.Button(widgets.QMessageBox__Cancel)

		// 修改按钮的文本
		okButton.SetText("确定")
		cancelButton.SetText("取消")

		result := msgBox.Exec()

		if result == int(widgets.QMessageBox__Ok) {
			// 用户点击了 "OK" 按钮
			fmt.Println("点击了 OK 按钮")
		} else if result == int(widgets.QMessageBox__Cancel) {
			// 用户点击了 "Cancel" 按钮
			fmt.Println("点击了 Cancel 按钮")
		} else {
			// 用户点击了其他按钮或关闭了消息框
			fmt.Println("关闭了消息框")
		}
	})

	fileMenu.AddSeparator()
	exitAction := fileMenu.AddAction("Exit")
	//exitAction.SetIcon(gui.QIcon_FromTheme("window-close"))
	//exitAction.SetShortcut(widgets.NewQKeySequence2("Ctrl+Q"))
	exitAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__Quit))
	//exitAction.SetShortcut(gui.NewQKeySequence5(gui.key))
	exitAction.SetShortcut(gui.NewQKeySequence2("Ctrl+S", gui.QKeySequence__NativeText))
	/*closeAction := widgets.NewQAction3(gui.QIcon_FromTheme("window-close"), "Close", nil)
	fileMenu.AddActions([]*widgets.QAction{closeAction})*/

	exitAction.SetIcon(gui.QIcon_FromTheme("edit-copy"))
	//closeAction := widgets.NewQAction3(gui.QIcon_FromTheme("window-close"), "Close", nil)
	//fileMenu.AddAction(closeAction,"close")

	openRecentMenu.AddSeparator()
	openRecentMenu.AddAction("&Clear List...")
	// Connect the actions and tool buttons to their respective triggered events
	openFileAction.ConnectTriggered(func(checked bool) {
		//widgets
		dialog := widgets.NewQDialog(nil, 0)
		dialog.SetWindowTitle("Open Dialog")

		// Create the tab widget
		//entryTabWidget := widgets.NewQWidget(nil, 0)
		//tabWidget := kpwidgets.NewKeePassDialog(dialog)
		//tabWidget.Resize(700, 400)

		//initEntryTab(a)

		//vBoxLayout := widgets.NewQVBoxLayout2(dialog)
		//	vBoxLayout.AddWidget(tabWidget.TabWidget, 0, core.Qt__AlignLeft)

		dialog.Resize2(600, 400)
		dialog.Exec()
	})

	openUrlAction.ConnectTriggered(func(checked bool) {
		// Action logic for "New"
	})

	exitAction.ConnectTriggered(func(checked bool) {
		//widgets.QMessageBox_Question(window, "是否退出?", "真的要退出吗?", widgets.QMessageBox__Ok, widgets.QMessageBox__Cancel)
		// 弹出确认对话框
		result := widgets.QMessageBox_Question(window, "确认退出", "确定要退出应用程序吗？", widgets.QMessageBox__Ok|widgets.QMessageBox__Cancel, widgets.QMessageBox__Cancel)
		if result == widgets.QMessageBox__Ok {
			// 用户点击了确定按钮，退出应用程序
			window.Close()
		}
	})

}

func DoNewAction2(window *widgets.QMainWindow) {
	dialog := NewNewEntryTipsDialog()

	dialog.Show()
}

func DoNewAction(window *widgets.QMainWindow) {
	messageBox := widgets.NewQMessageBox(nil)
	messageBox.SetTextFormat(core.Qt__RichText) // Use rich text format

	// Use HTML-style formatting to set font size
	messageText := "<font size=\"5\">" + kpResources.NewDatabase + "</font>"
	messageBox.SetText(messageText)
	messageBox.SetIcon(widgets.QMessageBox__Information)
	//	fileInfo := insertAfter(kpResources.DatabaseFileIntro, "KeePass database file", "\n")

	messageBox.SetInformativeText(kpResources.DatabaseFileIntro + "\n\n" + kpResources.DatabaseFileRem + "\n\n" + kpResources.BackupDatabase)
	//messageBox.SetDetailedText("Here is the detailed information about the error:\n\nLine 1: Something went wrong.\nLine 2: Please try again later.")
	messageBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
	messageBox.Exec()
}
func insertAfter(originalStr, searchString, insertString string) string {
	index := strings.Index(originalStr, searchString)
	if index != -1 {
		return originalStr[:index+len(searchString)] + insertString + originalStr[index+len(searchString):]
	}
	return originalStr
}
func DoNewAction3(window *widgets.QMainWindow) {
	msgBox := widgets.NewQMessageBox(nil)
	msgBox.SetWindowTitle("Message Box")
	msgBox.SetIcon(widgets.QMessageBox__Information)
	msgBox.SetText(kpResources.CreateNewDatabase2)
	msgBox.SetInformativeText(kpResources.DatabaseFileIntro)
	msgBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
	msgBox.SetDefaultButton2(widgets.QMessageBox__Ok)

	// Connect the clicked signal of the buttons
	msgBox.ConnectButtonClicked(func(button *widgets.QAbstractButton) {
		if button.Text() == "OK" {
			newFileBox := widgets.NewQFileDialog2(window, "新建", "", "*.txt")
			newFileBox.Show()
			newFileBox.ConnectFileSelected(func(file string) {
				fmt.Print(file)
			})
		} else if button.Text() == "&Cancel" {
			// Handle the logic for the Cancel button
			// ...
		}

		// Close the message box
		msgBox.Close()
	})
	msgBox.Resize2(800, 600)
	// Show the message box
	msgBox.Exec()
}
