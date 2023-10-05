package kpwidgets

import (
	"fmt"
	"github.com/techyang/keepassgo/entity"
	"github.com/techyang/keepassgo/functions"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"github.com/tobischo/gokeepasslib/v3/wrappers"
	log "log/slog"
	"os"
	"strings"
)

var kpResources = functions.NewKpResources()

func InitFileMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	fileMenu := menuBar.AddMenu2("File")
	core.QCoreApplication_Translate("myapp", "", "", -1)

	// Create actions for the file menu
	newAction := fileMenu.AddAction("&New...")
	newActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_FileNew.png")
	newAction.SetIcon(newActionIcon)

	// Connect the actions and tool buttons to their respective triggered events
	newAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})
	//newAction.SetShortcut(gui.NewQKeySequence2("Ctrl+N", gui.QKeySequence__NativeText))
	newAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__New))

	// Create actions for the file menu
	//openAction := fileMenu.AddAction("&Open")

	openMenu := fileMenu.AddMenu2("Open")
	openRecentMenu := fileMenu.AddMenu2("Open Recent")
	openFileAction := openMenu.AddAction("Open File...")
	openUrlAction := openMenu.AddAction("Open Url...")

	closeAction := fileMenu.AddAction("Close")
	closeActionIcon := gui.NewQIcon5("Resources/Nuvola_Derived/B16x16_File_Close.png")
	closeAction.SetIcon(closeActionIcon)
	closeAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})
	closeAction.SetShortcut(gui.NewQKeySequence2("Ctrl+W", gui.QKeySequence__NativeText))

	fileMenu.AddSeparator()
	saveAction := fileMenu.AddAction("&Save")
	saveAsMenu := fileMenu.AddMenu2("&Save As ...")
	saveActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_FileSave.png")
	saveAction.SetIcon(saveActionIcon)
	saveAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__Save))

	saveAsMenu.AddAction("&Save To File ...")
	saveAsMenu.AddAction("&Save To Url ...")
	saveAsMenu.AddSeparator()
	saveAsMenu.AddAction("&Save Copy To File ...")
	saveAction.ConnectTriggered(func(checked bool) {
		filename := "D:\\workspace_go\\gokeepasslib-master\\example-writing2023.kdbx"
		//masterPassword := "111111"

		// create root group

		file, _ := os.Open(filename)

		db := gokeepasslib.NewDatabase()
		db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
		_ = gokeepasslib.NewDecoder(file).Decode(db)

		db.UnlockProtectedEntries()

		gp := entity.FindGroupByUUID(db.Content.Root.Groups, "UpnnUcQ5/kKbVGEJ8PFqMA==")

		entry := gokeepasslib.NewEntry()
		entry.Values = append(entry.Values, mkValue("Title", "My GMail password"))
		entry.Values = append(entry.Values, mkValue("UserName", "example@gmail.com"))
		entry.Values = append(entry.Values, mkProtectedValue("Password", "hunter2"))
		gp.Entries = append(gp.Entries, entry)

		db = &gokeepasslib.Database{
			Header:      gokeepasslib.NewHeader(),
			Credentials: gokeepasslib.NewPasswordCredentials("111111"),
			Content: &gokeepasslib.DBContent{
				Meta: gokeepasslib.NewMetaData(),
				Root: &gokeepasslib.RootData{
					Groups: db.Content.Root.Groups,
				},
			},
		}

		// Lock entries using stream cipher
		db.LockProtectedEntries()

		// and encode it into the file
		keepassEncoder := gokeepasslib.NewEncoder(file)
		if err := keepassEncoder.Encode(db); err != nil {
			panic(err)
		}

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

	DatabaseSettingsAction := fileMenu.AddAction("Database Settings...")
	DatabaseSettingsActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_Package_Development.png")
	DatabaseSettingsAction.SetIcon(DatabaseSettingsActionIcon)
	DatabaseSettingsAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	changeMasterKeyAction := fileMenu.AddAction("Change Master Key...")
	changeMasterKeyActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_File_Locked.png")
	changeMasterKeyAction.SetIcon(changeMasterKeyActionIcon)
	changeMasterKeyAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	fileMenu.AddSeparator()
	printAction := fileMenu.AddAction("Print...")
	printActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_FilePrint.png")
	printAction.SetIcon(printActionIcon)
	printAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})
	printAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__Print))

	fileMenu.AddSeparator()

	importAction := fileMenu.AddAction("Import...")
	importActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_Folder_Inbox.png")
	importAction.SetIcon(importActionIcon)
	importAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	exportAction := fileMenu.AddAction("Export...")
	exportActionIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_Reload_Page.png")
	exportAction.SetIcon(exportActionIcon)
	exportAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	synchronizeMenu := fileMenu.AddMenu2("Synchronize")
	synchronizeWithFileAction := synchronizeMenu.AddAction("Synchronize with File...")
	synchronizeWithUrlAction := synchronizeMenu.AddAction("Synchronize with Url...")
	synchronizeMenuIcon := gui.NewQIcon5("Resources/Nuvola/B16x16_Reload_Page.png")
	synchronizeWithFileAction.SetIcon(synchronizeMenuIcon)
	synchronizeWithUrlAction.SetIcon(synchronizeMenuIcon)
	synchronizeWithFileAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})
	synchronizeWithUrlAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})

	fileMenu.AddSeparator()
	lockWorkspaceAction := fileMenu.AddAction("Lock Workspace")
	lockWorkspaceActionIcon := gui.NewQIcon5("Resources/images/B16x16_LockWorkspace.png")
	lockWorkspaceAction.SetIcon(lockWorkspaceActionIcon)
	lockWorkspaceAction.ConnectTriggered(func(checked bool) {
		//window.Close()
	})
	lockWorkspaceAction.SetShortcut(gui.NewQKeySequence2("Ctrl+L", gui.QKeySequence__NativeText))

	exitAction := fileMenu.AddAction("Exit")
	exitAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__Quit))
	exitAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))

	exitAction.SetIcon(gui.QIcon_FromTheme("edit-copy"))

	openRecentMenu.AddSeparator()
	openRecentMenu.AddAction("&Clear List...")
	// Connect the actions and tool buttons to their respective triggered events
	openFileAction.ConnectTriggered(func(checked bool) {
		functions.OpenDatabase(TreeWidget, TableWidget)
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

func mkProtectedValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{
		Key:   key,
		Value: gokeepasslib.V{Content: value, Protected: wrappers.BoolWrapper(NewBoolWrapper(true))},
	}
}

type BoolWrapper struct {
	Bool bool
}

func NewBoolWrapper(value bool) BoolWrapper {
	return BoolWrapper{
		Bool: value,
	}
}

func DoNewAction2(window *widgets.QMainWindow) {
	dialog := NewNewEntryTipsDialog()

	dialog.Show()
}

func DoNewAction(window *widgets.QMainWindow) {
	messageBox := widgets.NewQMessageBox(nil)
	messageBox.SetTextFormat(core.Qt__RichText) // Use rich text format
	messageBox.SetWindowTitle(window.WindowTitle())
	// Use HTML-style formatting to set font size
	messageText := "<font size=\"5\">" + kpResources.NewDatabase + "</font>"
	messageBox.SetText(messageText)
	messageBox.SetIcon(widgets.QMessageBox__Information)
	//	fileInfo := insertAfter(kpResources.DatabaseFileIntro, "KeePass database file", "\n")

	messageBox.SetInformativeText(kpResources.DatabaseFileIntro + "\n\n" + kpResources.DatabaseFileRem + "\n\n" + kpResources.BackupDatabase)
	//messageBox.SetDetailedText("Here is the detailed information about the error:\n\nLine 1: Something went wrong.\nLine 2: Please try again later.")
	messageBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
	messageBox.SetDefaultButton2(widgets.QMessageBox__Ok)
	okButton := messageBox.Button(widgets.QMessageBox__Ok)
	okButton.ConnectClicked(func(checked bool) {
		// Handle OK button clicked event
		log.Info("Handle OK button clicked event")
		functions.NewDatabase()
		messageBox.Accept()
	})

	// Connect Cancel button clicked signal
	cancelButton := messageBox.Button(widgets.QMessageBox__Cancel)
	cancelButton.ConnectClicked(func(checked bool) {
		// Handle Cancel button clicked event
		log.Info("Handle Cancel button clicked event")
		messageBox.Reject()
	})
	messageBox.Exec()
}
func insertAfter(originalStr, searchString, insertString string) string {
	index := strings.Index(originalStr, searchString)
	if index != -1 {
		return originalStr[:index+len(searchString)] + insertString + originalStr[index+len(searchString):]
	}
	return originalStr
}
