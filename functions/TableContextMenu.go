package functions

import (
	"fmt"
	"github.com/techyang/keepassgo/kpwidgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

func SetTableContextMenu(tableWidget *widgets.QTableWidget) {
	contextMenu := widgets.NewQMenu(nil)
	copyUserNameAction := contextMenu.AddAction("Copy User Name \tCtrl+B+C")
	copyUserNameAction.SetShortcut(gui.NewQKeySequence2("Ctrl+B", gui.QKeySequence__NativeText))
	//copyUserNameAction.SetShortcut(widgets.QKeySequence_fromString("Ctrl+O"))

	copyPasswordAction := contextMenu.AddAction("Copy Password")
	copyPasswordAction.SetShortcut(gui.NewQKeySequence2("Ctrl+C", gui.QKeySequence__NativeText))

	copyPasswordAction.SetMenuRole(widgets.QAction__TextHeuristicRole) // Show shortcut in the context menu
	//copyPasswordAction.

	urlsMenu := contextMenu.AddMenu2("URS(S)")
	openUrlAction := urlsMenu.AddAction("Open")
	copyUrlAction := urlsMenu.AddAction("Copy to ClipBoard")
	contextMenu.AddSeparator()

	performAutoTypeAction := contextMenu.AddAction("Perform Auto-Type")
	contextMenu.AddSeparator()
	addItemAction := contextMenu.AddAction("Add Entry...")
	editOrViewEntryAction := contextMenu.AddAction("Edit/View Entry...")
	duplicateAction := contextMenu.AddAction("Duplicate Entry...")

	//addItemAction := contextMenu.AddAction("Add Entry...")
	deleteItemAction := contextMenu.AddAction("Delete Entry")
	selectEntriesMenu := contextMenu.AddMenu2("Select Entries")
	selectEntriesMenu.AddAction("Duplicate Entry...")
	selectAllAction := contextMenu.AddAction("Select All")
	contextMenu.AddSeparator()
	clipbordMenu := contextMenu.AddMenu2("Clipbord")
	clipbordMenu.AddAction("Copy Entries...")
	clipbordMenu.AddAction("Paste Entries...")

	rearrangeMenu := contextMenu.AddMenu2("Rearrange")
	rearrangeMenu.AddAction("Move Entry to Top")
	rearrangeMenu.AddAction("Move Entry One Up")
	rearrangeMenu.AddAction("Move Entry One Down")
	rearrangeMenu.AddAction("Move Entry to Bottom")

	copyUserNameAction.ConnectTriggered(func(bool) {
		selectedRow := tableWidget.CurrentRow()

		// Retrieve the item at the first column of the selected row
		item := tableWidget.Item(selectedRow, 1)

		// Get the text of the item
		if item != nil {
			firstItemText := item.Text()
			fmt.Println("Text of the first item in the selected row:", firstItemText)
			clipboard := gui.QGuiApplication_Clipboard()
			if clipboard != nil {
				clipboard.SetText(firstItemText, gui.QClipboard__Clipboard)
			}
		}

	})
	copyPasswordAction.ConnectTriggered(func(bool) {
		selectedRow := tableWidget.CurrentRow()

		// Retrieve the item at the first column of the selected row
		item := tableWidget.Item(selectedRow, 2)

		// Get the text of the item
		if item != nil {
			firstItemText := item.Text()
			fmt.Println("Text of the first item in the selected row:", firstItemText)
			clipboard := gui.QGuiApplication_Clipboard()
			if clipboard != nil {
				clipboard.SetText(firstItemText, gui.QClipboard__Clipboard)
			}
		}
	})

	openUrlAction.ConnectTriggered(func(bool) {
		selectedRow := tableWidget.CurrentRow()

		// Retrieve the item at the first column of the selected row
		item := tableWidget.Item(selectedRow, 3)

		// Get the text of the item
		if item != nil {
			firstItemText := item.Text()
			fmt.Println("Text of the first item in the selected row:", firstItemText)
			gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput(firstItemText))
		}
	})

	copyUrlAction.ConnectTriggered(func(bool) {
		selectedRow := tableWidget.CurrentRow()

		// Retrieve the item at the first column of the selected row
		item := tableWidget.Item(selectedRow, 3)

		// Get the text of the item
		if item != nil {
			firstItemText := item.Text()
			fmt.Println("Text of the first item in the selected row:", firstItemText)
			clipboard := gui.QGuiApplication_Clipboard()
			if clipboard != nil {
				clipboard.SetText(firstItemText, gui.QClipboard__Clipboard)
			}
		}
	})
	performAutoTypeAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
	})
	editOrViewEntryAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
	})
	duplicateAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
	})

	selectAllAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
	})

	/*clipbordAction := contextMenu.AddMenu2("Copy User Name")
	clipbordAction := contextMenu.AddAction("Copy User Name")

	copyUserNameAction := contextMenu.AddAction("Copy User Name")

	copyUserNameAction := contextMenu.AddAction("Copy User Name")

	copyUserNameAction := contextMenu.AddAction("Copy User Name")
	*/

	tableWidget.ConnectCustomContextMenuRequested(func(pos *core.QPoint) {
		contextMenu.Exec2(tableWidget.MapToGlobal(pos), nil)
	})

	// Connect the triggered signal of the menu actions
	addItemAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
	})

	deleteItemAction.ConnectTriggered(func(bool) {
		// Get the selection model from the table view
		// Get the selection model from the table view
		//selectionModel := tableWidget.SelectionModel()
		//selectedRows := selectionModel.Selection()
		tableWidget.Model().RemoveRow(0, core.NewQModelIndex())

		//selectedIndexes := selectionModel.SelectedRows()

		/*qModelIndex := selectionModel.SelectedRows(0).
		qModelIndex.
		tableWidget.*/

		/*for _, index := range selectedRows {
			tableWidget.Model().RemoveRow(index.Row(), core.NewQModelIndex())
			fmt.Println("第", index, "行删除了")
		}*/

	})
}

func initDetailWidget(tableWidget *widgets.QTableWidget) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	keePassTabWidget := kpwidgets.NewKeePassTabWidget(dialog)
	keePassTabWidget.Resize(600, 400)

	hBoxLayout := initBottomButton(keePassTabWidget, tableWidget, dialog)

	vBoxLayout := widgets.NewQVBoxLayout2(dialog)
	vBoxLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(keePassTabWidget.TabWidget, 0, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(hBoxLayout, 0)

	dialog.Resize2(600, 400)
	dialog.Exec()

	return dialog
}

func initKeePassImage() *widgets.QLabel {
	imageLabel := widgets.NewQLabel(nil, 0)
	imagePixmap := gui.NewQPixmap3("src\\Hello\\img\\keepass.png", "", core.Qt__AutoColor)
	imageLabel.SetPixmap(imagePixmap)
	return imageLabel
}

func initBottomButton(keePassDialog *kpwidgets.KeePassTabWidget, tableWidget *widgets.QTableWidget, dialog *widgets.QDialog) *widgets.QHBoxLayout {
	entryTab := keePassDialog.EntryTab
	advancedTab := keePassDialog.AdvancedTab
	advancedTab.Widget.Parent()

	hBoxLayout := widgets.NewQHBoxLayout2(nil)
	toolButton := widgets.NewQPushButton2("Tool", nil)

	toolButton.SetIcon(gui.NewQIcon5(":/Ext/Images_Client_16/C59_Package_Development.png")) // Replace with the path to your icon file
	toolButton.SetIconSize(core.NewQSize2(32, 32))                                          // Set the size of the icon
	//toolButton.SetText("Tool")

	okButton := widgets.NewQPushButton2("Ok", nil)
	cancelButton := widgets.NewQPushButton2("Cancel", nil)
	spacer := widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)

	hBoxLayout.AddWidget(toolButton, 0, core.Qt__AlignLeft)
	hBoxLayout.AddSpacerItem(spacer)
	hBoxLayout.AddWidget(okButton, 0, core.Qt__AlignRight)
	hBoxLayout.AddWidget(cancelButton, 0, core.Qt__AlignRight)

	toolButton.ConnectClicked(func(bool) {
		// Code to handle cancelButton click event
		fmt.Println("toolButton clicked")
		//dialog.Close()
	})

	okButton.ConnectClicked(func(bool) {

		msgBox := widgets.NewQMessageBox(nil)
		msgBox.SetWindowTitle("提示信息")
		//msgBox.SetText(keePassDialog.EntryTab.UserNameEdit.Text())
		msgBox.SetInformativeText(keePassDialog.EntryTab.UserNameEdit.Text())
		msgBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
		msgBox.SetIcon(widgets.QMessageBox__Information)
		msgBox.Exec()

		// Code to handle cancelButton click event
		//tabWidget.get
		fmt.Println("okButton clicked")
		kpwidgets.ReAddTableItem(entryTab, tableWidget)

		file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

		db := gokeepasslib.NewDatabase()
		db.Credentials = gokeepasslib.NewPasswordCredentials("supersecret")
		_ = gokeepasslib.NewDecoder(file).Decode(db)

		db.UnlockProtectedEntries()

		fmt.Println("Password entry saved successfully.")

		dialog.Close()
	})

	cancelButton.ConnectClicked(func(bool) {
		// Code to handle cancelButton click event
		fmt.Println("Cancel Button clicked")
		dialog.Close()
	})
	return hBoxLayout
}
