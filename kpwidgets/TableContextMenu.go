package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
	"os/exec"
	"runtime"
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
	moveTopAction := rearrangeMenu.AddAction("Move Entry to Top")
	moveUpAction := rearrangeMenu.AddAction("Move Entry One Up")
	moveDownAction := rearrangeMenu.AddAction("Move Entry One Down")
	moveBottomAction := rearrangeMenu.AddAction("Move Entry to Bottom")

	moveUpAction.SetShortcut(gui.NewQKeySequence2("Alt+", gui.QKeySequence__NativeText))

	setMoveTopAction(tableWidget, moveTopAction)
	setMoveUpAction(tableWidget, moveUpAction)
	setMoveDownAction(tableWidget, moveDownAction)
	setMoveBottomAction(tableWidget, moveBottomAction)

	setCopyUserNameAction(tableWidget, copyUserNameAction)
	setCopyPasswordAction(tableWidget, copyPasswordAction)

	setOpenUrlAction(tableWidget, openUrlAction)

	setCopyUrlAction(tableWidget, copyUrlAction)
	performAutoTypeAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
	})
	setEditOrViewEntryAction(tableWidget, editOrViewEntryAction)
	duplicateAction.ConnectTriggered(func(bool) {

		dialog := widgets.NewQDialog(nil, 0)
		dialog.SetWindowTitle("Duplication Options")

		// Create the tab widget

		appendCopyCheckBox := widgets.NewQCheckBox2("Apend \"-Copy\" to entry titles", nil)
		repeatUserNameCheckBox := widgets.NewQCheckBox2("Enable auto-type for this entry", nil)
		helpLabel1 := widgets.NewQLabel2("If this option is enabled, the copies will reference \nthe user names and passwords of the original entries.\nWhen a user name or password is changed in an original \nentry, the copy will automatically use the new data, too.", nil, 0)
		helpLabel2 := widgets.NewQLabel2("Help: Field References", nil, 0)

		helpLabel2.SetText("<a href=\"D:\\Program Files (x86)\\Microsoft Visual Studio\\2019\\Community\\Common7\\Tools\\spyxx.chm\">Help: Field References</a>")

		helpLabel2.SetTextInteractionFlags(core.Qt__LinksAccessibleByMouse)

		helpLabel2.ConnectLinkActivated(func(link string) {
			//link = "D:\\Program Files\\TotalCMDPortable\\App\\totalcmd\\Plugins\\wcx\\Total7zip\\7-zip.chm"
			/*if err := exec.Command("hh.exe", link).Start(); err != nil {
				fmt.Println("Error opening .chm file:", err)
			}*/
			//chmFilePath := "path/to/helpfile.chm"

			// Replace "your-topic" with the topic you want to display in the CHM file.
			topic := "Working In Spy++"

			// Use different commands based on the operating system.
			var command string
			var args []string
			switch runtime.GOOS {
			case "windows":
				command = "hh.exe"
				args = []string{link, fmt.Sprintf("-#%s", topic)}
			default:
				fmt.Println("Unsupported operating system.")
				os.Exit(1)
			}

			cmd := exec.Command(command, args...)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error opening CHM file:", err)
				os.Exit(1)
			}
		})

		hBoxLayout := widgets.NewQHBoxLayout2(nil)
		hBoxLayout.AddSpacing(20)
		hBoxLayout.AddWidget(helpLabel1, 0, core.Qt__AlignLeft)
		//helpLabel2 := widgets.NewQLabel2("<a href=\"https://keepass.info\">KeePass website</a>", nil, 0)

		hBoxLayout2 := widgets.NewQHBoxLayout2(nil)
		hBoxLayout2.AddSpacing(20)
		hBoxLayout2.AddWidget(helpLabel2, 0, core.Qt__AlignLeft)

		copyHistoryCheckBox := widgets.NewQCheckBox2("Copy history", nil)
		vBoxLayout := widgets.NewQVBoxLayout2(dialog)
		vBoxLayout.AddWidget(appendCopyCheckBox, 0, core.Qt__AlignLeft)
		vBoxLayout.AddWidget(repeatUserNameCheckBox, 0, core.Qt__AlignLeft)
		vBoxLayout.AddLayout(hBoxLayout, 0)
		vBoxLayout.AddLayout(hBoxLayout2, 0)
		vBoxLayout.AddWidget(copyHistoryCheckBox, 0, core.Qt__AlignLeft)

		separator := widgets.NewQFrame(nil, 0)
		separator.SetFrameShape(widgets.QFrame__HLine)
		separator.SetLineWidth(20)
		vBoxLayout.AddWidget(separator, 0, core.Qt__AlignLeft)
		vBoxLayout.AddWidget(separator, 0, core.Qt__AlignLeft)
		// Create the button box
		buttonBox := widgets.NewQDialogButtonBox(dialog)
		okButton := buttonBox.AddButton3(widgets.QDialogButtonBox__Ok)
		cancelButton := buttonBox.AddButton3(widgets.QDialogButtonBox__Cancel)
		// Set the button text
		okButton.SetText("OK")
		cancelButton.SetText("Cancel")
		// Connect the button box's accepted signal
		buttonBox.ConnectAccepted(func() {
			fmt.Println("OK button clicked")
			dialog.Accept()

			row := tableWidget.CurrentRow()
			if row <= tableWidget.RowCount()-1 {
				fieldName := tableWidget.Item(row, 0).Text()
				fieldValue := tableWidget.Item(row, 1).Text()
				fieldValue2 := tableWidget.Item(row, 2).Text()
				fieldValue3 := tableWidget.Item(row, 3).Text()

				if appendCopyCheckBox.IsChecked() {
					fieldName += "-Copy"
				}

				tableWidget.InsertRow(row + 1)
				tableWidget.SetItem(row+1, 0, widgets.NewQTableWidgetItem2(fieldName, 0))
				tableWidget.SetItem(row+1, 1, widgets.NewQTableWidgetItem2(fieldValue, 0))
				tableWidget.SetItem(row+1, 2, widgets.NewQTableWidgetItem2(fieldValue2, 0))
				tableWidget.SetItem(row+1, 3, widgets.NewQTableWidgetItem2(fieldValue3, 0))
				tableWidget.SelectRow(row + 1)
				/*tableWidget.InsertRow(row - 1)
				for column := 0; column < tableWidget.ColumnCount(); column++ {
					tableWidget.SetItem(row-1, column, tableWidget.Item(row, column))
				}*/
			}

		})

		// Connect the button box's rejected signal
		buttonBox.ConnectRejected(func() {
			fmt.Println("Cancel button clicked")
			dialog.Reject()
		})

		layout := widgets.NewQVBoxLayout2(dialog)
		layout.AddWidget(buttonBox, 0, core.Qt__AlignRight)

		vBoxLayout.AddLayout(layout, 0)
		dialog.Resize2(450, 250)
		dialog.Exec()

	})

	selectAllAction.ConnectTriggered(func(bool) {
		tableWidget.SelectAll()
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

		row := tableWidget.CurrentRow()
		if row > 0 {
			tableWidget.RemoveRow(row)
		}

		// Get the selection model from the table view
		// Get the selection model from the table view
		//selectionModel := tableWidget.SelectionModel()
		//selectedRows := selectionModel.Selection()
		//tableWidget.Model().RemoveRow(0, core.NewQModelIndex())

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

func setEditOrViewEntryAction(tableWidget *widgets.QTableWidget, editOrViewEntryAction *widgets.QAction) {
	editOrViewEntryAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
		//row := tableWidget.CurrentRow()
		//objectName := tableWidget.ObjectName()
		//GetKeePassEntry(objectName, row)
	})
}

func setCopyUrlAction(tableWidget *widgets.QTableWidget, copyUrlAction *widgets.QAction) {
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
}

func setOpenUrlAction(tableWidget *widgets.QTableWidget, openUrlAction *widgets.QAction) {
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
}

func setCopyPasswordAction(tableWidget *widgets.QTableWidget, copyPasswordAction *widgets.QAction) {
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
}

func setCopyUserNameAction(tableWidget *widgets.QTableWidget, copyUserNameAction *widgets.QAction) {
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
}

func setMoveTopAction(tableWidget *widgets.QTableWidget, moveTopAction *widgets.QAction) {
	moveTopAction.ConnectTriggered(func(checked bool) {
		row := tableWidget.CurrentRow()
		if row > 0 {
			rowData := getRowData(tableWidget, row)
			tableWidget.RemoveRow(row)
			newRow := 0
			tableWidget.InsertRow(newRow)
			setTableRowData(tableWidget, newRow, rowData)
		}
	})
}

func setMoveBottomAction(tableWidget *widgets.QTableWidget, moveBottomAction *widgets.QAction) {
	moveBottomAction.ConnectTriggered(func(checked bool) {
		row := tableWidget.CurrentRow()
		if row < tableWidget.RowCount()-1 {
			rowData := getRowData(tableWidget, row)
			tableWidget.RemoveRow(row)
			newRow := tableWidget.RowCount()
			tableWidget.InsertRow(newRow)
			setTableRowData(tableWidget, newRow, rowData)
		}
	})
}

func setMoveUpAction(tableWidget *widgets.QTableWidget, moveUpAction *widgets.QAction) {
	moveUpAction.ConnectTriggered(func(checked bool) {
		row := tableWidget.CurrentRow()
		if row > 0 {
			rowData := getRowData(tableWidget, row)
			tableWidget.RemoveRow(row)
			newRow := row - 1
			tableWidget.InsertRow(newRow)
			setTableRowData(tableWidget, newRow, rowData)
		}
	})
}

func setMoveDownAction(tableWidget *widgets.QTableWidget, moveDownAction *widgets.QAction) {
	moveDownAction.ConnectTriggered(func(checked bool) {
		row := tableWidget.CurrentRow()
		if row < tableWidget.RowCount()-1 {
			rowData := getRowData(tableWidget, row)
			tableWidget.RemoveRow(row)
			newRow := row + 1
			tableWidget.InsertRow(newRow)
			setTableRowData(tableWidget, newRow, rowData)
		}
	})
}

func setTableRowData(tableWidget *widgets.QTableWidget, newRow int, rowData []string) {
	for column := 0; column < tableWidget.ColumnCount(); column++ {
		tableWidget.SetItem(newRow, column, widgets.NewQTableWidgetItem2(rowData[column], 0))
	}
	tableWidget.SelectRow(newRow)
}

// Function to retrieve the data of a specific row in a QTableWidget and store it in an array
func getRowData(tableWidget *widgets.QTableWidget, row int) []string {
	columnCount := tableWidget.ColumnCount()
	rowData := make([]string, columnCount)

	for col := 0; col < columnCount; col++ {
		item := tableWidget.Item(row, col)
		if item != nil {
			rowData[col] = item.Text()
		} else {
			rowData[col] = ""
		}
	}

	return rowData
}

func initDetailWidget(tableWidget *widgets.QTableWidget) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	keePassTabWidget := NewKeePassTabWidget(dialog)
	keePassTabWidget.Resize(600, 400)
	keePassTabWidget.EntryTab.InitEntryTab(tableWidget)
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

func initBottomButton(keePassDialog *KeePassTabWidget, tableWidget *widgets.QTableWidget, dialog *widgets.QDialog) *widgets.QHBoxLayout {
	entryTab := keePassDialog.EntryTab
	advancedTab := keePassDialog.AdvancedTab
	advancedTab.Widget.Parent()

	hBoxLayout := widgets.NewQHBoxLayout2(nil)
	toolButton := widgets.NewQPushButton2("Tool", nil)

	toolButton.SetIcon(gui.NewQIcon5("/Ext/Images_Client_16/C59_Package_Development.png")) // Replace with the path to your icon file
	toolButton.SetIconSize(core.NewQSize2(32, 32))                                         // Set the size of the icon
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
		ReAddTableItem(entryTab, tableWidget)

		file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

		db := gokeepasslib.NewDatabase()
		db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
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
