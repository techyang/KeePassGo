package kpwidgets

import (
	"github.com/techyang/keepassgo/entity"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func InitEditMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	editMenu := menuBar.AddMenu2("Edit")

	// Create actions for the file menu
	addGroupAction := editMenu.AddAction("&Add Group...")
	//addGroupAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_Folder_Txt.png"))
	addGroupAction.SetIcon(gui.NewQIcon5("Ext/Images_App_HighRes/Nuvola_Derived/B48x48_Folder_New_Ex.png"))
	// Connect the actions and tool buttons to their respective triggered events
	addGroupAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the file menu
	editGroupAction := editMenu.AddAction("&Edit Group...")
	editGroupAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_Folder_Txt.png"))

	// Connect the actions and tool buttons to their respective triggered events
	editGroupAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	delteteGroupAction := editMenu.AddAction("&Delete Group...")
	delteteGroupAction.SetIcon(gui.NewQIcon5("Ext/Images_App_HighRes/Nuvola/B48x48_Folder_Locked.png"))
	delteteGroupAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	editMenu.AddSeparator()
	// Create actions for the file menu
	//openAction := fileMenu.AddAction("&Open")

	// Create actions for the edit menu
	addEntityAction := editMenu.AddAction("&Add Entity...")
	addEntityAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Import.png"))
	// Connect the actions and tool buttons to their respective triggered events
	addEntityAction.ConnectTriggered(func(checked bool) {
		entity.NewDetailWidget(TableWidget)
	})

	addEntityAction.SetShortcut(gui.NewQKeySequence2("Ctrl+I", gui.QKeySequence__NativeText))

	// Create actions for the edit menu
	editEntityAction := editMenu.AddAction("&Edit Entity...")
	editEntityAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Sign.png"))
	// Connect the actions and tool buttons to their respective triggered events
	editEntityAction.ConnectTriggered(func(checked bool) {
		entity.InitDetailWidget(TableWidget)
	})

	// Create actions for the edit menu
	duplicateEntityAction := editMenu.AddAction("&Duplicate Entity...")
	duplicateEntityAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Key2.png"))

	// Connect the actions and tool buttons to their respective triggered events
	duplicateEntityAction.ConnectTriggered(func(checked bool) {
		entity.DuplicationEntity(TableWidget)
	})

	DelteteEntityAction := editMenu.AddAction("&Delete Entity...")
	DelteteEntityAction.SetIcon(gui.NewQIcon5("Resources/Nuvola_Derived/B16x16_DeleteEntry.png"))
	DelteteEntityAction.SetShortcut(gui.NewQKeySequence2("Delete", gui.QKeySequence__NativeText))
	DelteteEntityAction.ConnectTriggered(func(checked bool) {
		TableWidget.Delete()
	})

	editMenu.AddSeparator()

	// Create actions for the edit menu
	selectAllAction := editMenu.AddAction("&Select All...")
	// Connect the actions and tool buttons to their respective triggered events
	selectAllAction.ConnectTriggered(func(checked bool) {
		TableWidget.SelectAll()
	})
	selectAllAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__SelectAll))
	editMenu.AddSeparator()

	showEntriesMenu := editMenu.AddMenu2("Show Entries...")
	showEntriesByTagMenu := editMenu.AddMenu2("Show Entries by Tag...")
	// Connect the actions and tool buttons to their respective triggered events

	showEntriesAllAction := showEntriesMenu.AddAction("All")
	showEntriesGroupAction := showEntriesMenu.AddAction("Selected Entry's Group")
	showEntriesAllIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Key3.png")
	showEntriesGroupIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_Folder.png")

	showEntriesAllAction.SetIcon(showEntriesAllIcon)
	showEntriesGroupAction.SetIcon(showEntriesGroupIcon)
	showEntriesGroupAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	noTagFoundAction := showEntriesByTagMenu.AddAction("(No Tag Found)")
	noTagFoundAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	editMenu.AddSeparator()
	findAction := editMenu.AddAction("&Find...")
	findAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_XMag.png"))

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
