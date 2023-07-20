package functions

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func InitEditMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	editMenu := menuBar.AddMenu2("Edit")

	// Create actions for the file menu
	addGroupAction := editMenu.AddAction("&Add Group...")
	// Connect the actions and tool buttons to their respective triggered events
	addGroupAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the file menu
	editGroupAction := editMenu.AddAction("&Edit Group...")
	// Connect the actions and tool buttons to their respective triggered events
	editGroupAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the file menu
	DelteteGroupAction := editMenu.AddAction("&Delete Group...")
	// Connect the actions and tool buttons to their respective triggered events
	DelteteGroupAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	editMenu.AddSeparator()
	// Create actions for the file menu
	//openAction := fileMenu.AddAction("&Open")

	// Create actions for the edit menu
	addEntityAction := editMenu.AddAction("&Add Entity...")
	// Connect the actions and tool buttons to their respective triggered events
	addEntityAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	addEntityAction.SetShortcut(gui.NewQKeySequence2("Ctrl+I", gui.QKeySequence__NativeText))

	// Create actions for the edit menu
	editEntityAction := editMenu.AddAction("&Edit Entity...")
	// Connect the actions and tool buttons to their respective triggered events
	editEntityAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the edit menu
	duplicateEntityAction := editMenu.AddAction("&Duplicate Entity...")
	// Connect the actions and tool buttons to their respective triggered events
	duplicateEntityAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the edit menu
	DelteteEntityAction := editMenu.AddAction("&Delete Entity...")
	// Connect the actions and tool buttons to their respective triggered events
	DelteteEntityAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	editMenu.AddSeparator()

	// Create actions for the edit menu
	selectAllAction := editMenu.AddAction("&Select All...")
	// Connect the actions and tool buttons to their respective triggered events
	selectAllAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})
	selectAllAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__SelectAll))
	editMenu.AddSeparator()

	showEntriesAction := editMenu.AddAction("&Select All...")
	// Connect the actions and tool buttons to their respective triggered events
	showEntriesAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})
	showEntriesByTagAction := editMenu.AddAction("&Select All...")
	// Connect the actions and tool buttons to their respective triggered events
	showEntriesByTagAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	editMenu.AddSeparator()
	findAction := editMenu.AddAction("&Find...")
	// Connect the actions and tool buttons to their respective triggered events
	findAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})
	//helpAction.SetShortcuts2(core.NewQKeySequence2("Ctrl+O", core.QKeySequence__NativeText))
	//helpAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Alt+S", gui.QKeySequence__NativeText))
	findAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__Find))
	/*openMenu := fileMenu.AddMenu2("&Open")
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
	fileMenu.AddActions([]*widgets.QAction{closeAction})

	exitAction.SetIcon(gui.QIcon_FromTheme("edit-copy"))
	//closeAction := widgets.NewQAction3(gui.QIcon_FromTheme("window-close"), "Close", nil)
	//fileMenu.AddAction(closeAction,"close")

	openRecentMenu.AddSeparator()
	openRecentMenu.AddAction("&Clear List...")
	// Connect the actions and tool buttons to their respective triggered events
	openFileAction.ConnectTriggered(func(checked bool) {
		//widgets
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
	*/
}
