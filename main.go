package main

import (
	"fmt"
	"github.com/techyang/keepassgo/functions"
	"github.com/techyang/keepassgo/kpwidgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

func clearChildItems(item *widgets.QTreeWidgetItem) {
	for item.ChildCount() > 0 {
		item.TakeChild(0)
	}
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Create the main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowIcon(gui.NewQIcon5("Ext/Icons_15_VA/KeePass_Round/KeePass_Round_24.png"))
	window.SetWindowTitle("KeePass")

	// Create the menu bar
	initMenuBar(window)
	// Create the toolbar with a title
	initToolbar(window)
	initMainContent(window)

	// 创建状态栏
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)
	// 在状态栏中显示文本
	statusBar.ShowMessage("Ready", 0)

	// Show the main window
	window.Resize2(800, 650)

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
	menu := widgets.NewQMenu(nil)
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

	//设置菜单
	sys.SetContextMenu(menu)
	sys.Show()

	window.Show()

	widgets.QApplication_Exec()
}

func initMainContent(window *widgets.QMainWindow) {
	// Create a QVBoxLayout and a QWidget as the container

	splitter := widgets.NewQSplitter2(core.Qt__Horizontal, nil)

	// Create the left and right widgets
	//leftWidget := widgets.NewQWidget(nil, 0)
	hBoxLayout := widgets.NewQVBoxLayout()
	//	hBoxLayout2 := widgets.NewQVBoxLayout()

	container := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQHBoxLayout2(container)

	// Create the QHBoxLayout for the buttons
	//	hBoxLayout := widgets.NewQVBoxLayout()

	// Create the buttons
	//button1 := widgets.NewQPushButton2("Button 1", nil)
	//button2 := widgets.NewQPushButton2("Button 2", nil)

	// Create a QTreeWidget
	tableWidget := initTableWidget()
	treeWidget := initTreeWidget(tableWidget)

	//restTable(tableWidget)

	// Add the buttons to the QHBoxLayout
	//hBoxLayout.AddWidget(button1, 0, 0)
	hBoxLayout.AddWidget(treeWidget, 0, 0)

	button3 := widgets.NewQPushButton2("Button 3", nil)
	// Add the QHBoxLayout to the QVBoxLayout
	layout.AddLayout(hBoxLayout, 0)
	layout.AddWidget(button3, 0, 0)

	splitter.AddWidget(treeWidget)
	splitter.AddWidget(tableWidget)
	// Set the container as the central widget of the main window
	window.SetCentralWidget(splitter)

}

func setGroupKeePassItem(group *gokeepasslib.Group, tableWidget *widgets.QTableWidget) {

	entries := group.Entries

	tableWidget.Clear()
	tableWidget.SetRowCount(len(entries) + 1)
	//tableWidget.
	for i, entry := range entries {

		tableWidget.SetItem(i, 0, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
		tableWidget.SetItem(i, 1, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
		tableWidget.SetItem(i, 2, widgets.NewQTableWidgetItem2(entry.GetPassword(), 0))
		fmt.Println(entry.GetTitle())
		fmt.Println(entry.GetPassword())
	}

}

func initTreeWidget(tableWidget *widgets.QTableWidget) *widgets.QTreeWidget {
	treeWidget := widgets.NewQTreeWidget(nil)
	//treeWidget.SetHeaderLabels([]string{"yangwl"})
	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("supersecret")
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()
	rootGroups := db.Content.Root.Groups
	for i, rootGroup := range rootGroups {
		fmt.Println(i, "rootGroup:", rootGroup.Name)
		// Create the root item
		rootItem := widgets.NewQTreeWidgetItem4(treeWidget, []string{rootGroup.Name, "1.1"}, 0)
		rootItem.SetExpanded(true) // Set the root item initially expanded
		groups := rootGroup.Groups
		buildGroupTree(rootItem, groups)

		treeWidget.InsertTopLevelItem(i, rootItem)
	}
	// Create the root item

	// Set the root item as the top-level item of the tree widget

	treeWidget.SetHeaderHidden(true)

	// Connect the itemClicked signal of the tree widget
	TableItemClicked(tableWidget, treeWidget, rootGroups)
	return treeWidget
}

func TableItemClicked(tableWidget *widgets.QTableWidget, treeWidget *widgets.QTreeWidget, rootGroups []gokeepasslib.Group) {
	treeWidget.ConnectItemClicked(func(item *widgets.QTreeWidgetItem, column int) {

		groupUUID := item.Data(1, 0).ToString()
		fmt.Println(item.Text(0), "点击了", groupUUID)

		group := findGroupByUUID(rootGroups, groupUUID)

		if group != nil && group.Entries != nil {
			headerLabels := []string{"Title", "User Name", "Password", "URL", "Notes"}
			tableWidget.SetHorizontalHeaderLabels(headerLabels)
			setTableItems(group, tableWidget)
		} else {
			headerLabels := []string{"Title", "User Name", "Password", "URL", "Notes"}
			tableWidget.SetHorizontalHeaderLabels(headerLabels)
			//tableWidget.Clear()
			tableWidget.SetRowCount(0)
		}

	})
}

func setTableItems(group *gokeepasslib.Group, tableWidget *widgets.QTableWidget) {
	for i, entry := range group.Entries {
		username := entry.Get("UserName").Value.Content
		url := entry.Get("URL").Value.Content
		note := entry.Get("Notes").Value.Content
		tableWidget.SetRowCount(i + 1)
		tableWidget.SetItem(i, 0, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
		tableWidget.SetItem(i, 1, widgets.NewQTableWidgetItem2(username, 0))

		passwordItem := widgets.NewQTableWidgetItem2(entry.GetPassword(), 0)
		passwordItem.SetFlags(core.Qt__ItemIsSelectable | core.Qt__ItemIsEditable)
		//passwordItem.SetFlags(passwordItem.Flags() | core.Qt__ItemIsUserCheckable)
		//passwordItem.SetCheckState(core.Qt__Checked)

		tableWidget.SetItem(i, 2, passwordItem)
		tableWidget.SetItem(i, 3, widgets.NewQTableWidgetItem2(url, 0))
		tableWidget.SetItem(i, 4, widgets.NewQTableWidgetItem2(note, 0))
	}
}

func buildGroupTree(parent *widgets.QTreeWidgetItem, groups []gokeepasslib.Group) {
	for _, group := range groups {
		txt, _ := group.UUID.MarshalText()

		fmt.Println("group.UUID -----------:", group.Name, ":", string(txt))
		treeItem := widgets.NewQTreeWidgetItem2([]string{group.Name}, 0)
		//treeItem.SetData(0, core.Qt__UserRole, core.NewQVariant1(group.UUID.String()))

		treeItem.SetData(1, 0, core.NewQVariant1(string(txt)))

		parent.AddChild(treeItem)
		buildGroupTree(treeItem, group.Groups)
	}
}

func findGroupByUUID(groups []gokeepasslib.Group, uuid string) *gokeepasslib.Group {
	//uuids := uuid.MustParse(uuid)
	for _, group := range groups {
		txt, _ := group.UUID.MarshalText()
		if string(txt) == uuid {
			fmt.Println("找到的名称是:", group.Name)
			return &group
		}
		foundGroup := findGroupByUUID(group.Groups, uuid)
		if foundGroup != nil {
			return foundGroup
		}
	}
	return nil
}

/*func initDetailWidget2(tableWidget *widgets.QTableWidget) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	tabWidget := widgets.NewQTabWidget(dialog)
	entryTab, advancedTab := initTabWidget(tabWidget)
	//entryTabWidget := widgets.NewQWidget(nil, 0)
	entry := &kpwidgets.EntryTab{}
	entry.InitEntryTab2(entryTab)

	//initEntryTab(a)
	initAdvancedTab(advancedTab)

	hBoxLayout := initBottomButton(entry, tableWidget, tabWidget, dialog)

	vBoxLayout := widgets.NewQVBoxLayout2(dialog)
	vBoxLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(tabWidget, 0, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(hBoxLayout, 0)

	dialog.Resize2(600, 400)
	dialog.Exec()

	return dialog
}*/

func initDetailWidget(tableWidget *widgets.QTableWidget) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	keePassDialog := kpwidgets.NewKeePassDialog(dialog)
	keePassDialog.Resize(600, 400)
	//entryTab, advancedTab := initTabWidget(keePassDialog.TabWidget)
	//entryTabWidget := widgets.NewQWidget(nil, 0)
	//entry := &kpwidgets.EntryTab{}
	//entry.InitEntryTab2(keePassDialog.EntryTab)

	//initEntryTab(a)
	//initAdvancedTab(keePassDialog.AdvancedTab)
	//keePassDialog.AdvancedTab =
	//kpwidgets.NewAdvanceTab(keePassDialog.AdvancedTab)
	hBoxLayout := initBottomButton(keePassDialog, tableWidget, keePassDialog.TabWidget, dialog)

	vBoxLayout := widgets.NewQVBoxLayout2(dialog)
	vBoxLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(keePassDialog.TabWidget, 0, core.Qt__AlignLeft)
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

func initBottomButton(keePassDialog *kpwidgets.KeePassDialog, tableWidget *widgets.QTableWidget, tabWidget *widgets.QTabWidget, dialog *widgets.QDialog) *widgets.QHBoxLayout {
	entryTab := keePassDialog.EntryTab
	advancedTab := keePassDialog.AdvancedTab
	advancedTab.Widget.Parent()

	hBoxLayout := widgets.NewQHBoxLayout2(nil)
	toolButton := widgets.NewQPushButton2("Tool", nil)
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
		// Code to handle cancelButton click event
		//tabWidget.get
		fmt.Println("okButton clicked")
		kpwidgets.ReAddTableItem(entryTab, tableWidget)

		file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

		db := gokeepasslib.NewDatabase()
		db.Credentials = gokeepasslib.NewPasswordCredentials("supersecret")
		_ = gokeepasslib.NewDecoder(file).Decode(db)

		db.UnlockProtectedEntries()

		// Find the group by UUID
		targetGroup := findGroupByUUID(db.Content.Root.Groups, "your-group-uuid")
		/*if err != nil {
			fmt.Println("Error finding the group:", err)
			return
		}*/

		// Create a new password entry
		entry := gokeepasslib.NewEntry()
		entry.Values = append(entry.Values, mkValue("Title", "My GMail password"))
		entry.Values = append(entry.Values, mkValue("UserName", "example@gmail.com"))
		//entry.Values = append(entry.Values, mkProtectedValue("Password", "hunter2"))

		targetGroup.Entries = append(targetGroup.Entries, entry)
		/*// Add the new entry to the group
		targetGroup.Entries = append(targetGroup.Entries, newEntry)

		// Add the new entry to the group
		targetGroup.Entries = append(targetGroup.Entries, newEntry)*/

		// Save the Keepass database
		/*newFile, err := os.OpenFile("path/to/your/new_keepass.kdbx", os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Println("Error creating the new Keepass database:", err)
			return
		}*/

		/*	err = writer.WriteDatabase(newFile, db, gokeepasslib.WithPassword("your-master-password"))
			if err != nil {
				fmt.Println("Error saving the Keepass database:", err)
				return
			}*/

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

func mkValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{Key: key, Value: gokeepasslib.V{Content: value}}
}

/*func mkProtectedValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{
		Key:   key,
		Value: gokeepasslib.V{Content: value, Protected: NewBoolWrapper(true)},
	}
}*/

// Function to calculate the password complexity score
func calculatePasswordComplexity(password string) int {
	// Dummy implementation, replace with your own logic
	// Calculate the complexity based on the password strength criteria
	// Return a score between 0 and 100
	return len(password) * 10
}

func initAdvancedTab(advancedTab *widgets.QWidget) {

	// Set the layout for the second tab
	tab2Layout := widgets.NewQVBoxLayout2(advancedTab)

	// Create and add widgets to the second tab
	label3 := widgets.NewQLabel2("This is Tab 2", nil, 0)

	tab2Layout.AddWidget(label3, 0, core.Qt__AlignCenter)

}

func initTopImage(vBoxLayout *widgets.QVBoxLayout, tabWidget *widgets.QTabWidget, hBoxLayout *widgets.QHBoxLayout) {
	// Load and set the image pixmap
	imageLabel := widgets.NewQLabel(nil, 0)
	imagePixmap := gui.NewQPixmap3("D:\\workspace_go\\KeePassGo\\src\\Hello\\img\\keepass.png", "", core.Qt__AutoColor)
	imageLabel.SetPixmap(imagePixmap)
	vBoxLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(tabWidget, 0, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(hBoxLayout, 0)
}

func initTableWidget() *widgets.QTableWidget {
	// Create a QTableWidget
	// Create the table widget
	tableWidget := widgets.NewQTableWidget(nil)
	tableWidget.SetColumnCount(5)

	// Set the header labels
	headerLabels := []string{"Title", "User Name", "Password", "URL", "Notes"}
	tableWidget.SetHorizontalHeaderLabels(headerLabels)

	// Set the row count, must set to show table content
	/*tableWidget.SetRowCount(2)

	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(0, 0, widgets.NewQTableWidgetItem2("KeePass", 0))
	tableWidget.SetItem(0, 1, widgets.NewQTableWidgetItem2("Status 1", 0))
	tableWidget.SetItem(0, 2, widgets.NewQTableWidgetItem2("2.37", 0))
	tableWidget.SetItem(1, 0, widgets.NewQTableWidgetItem2("keePassLicC", 0))
	tableWidget.SetItem(1, 1, widgets.NewQTableWidgetItem2("Status 2", 0))
	tableWidget.SetItem(1, 2, widgets.NewQTableWidgetItem2("1.34", 0))
	*/
	// Enable sorting
	tableWidget.SetSortingEnabled(true)

	// Store the current sort order for each column
	sortOrders := make([]core.Qt__SortOrder, tableWidget.ColumnCount())

	// Connect the sectionClicked signal of horizontalHeader
	tableWidget.HorizontalHeader().ConnectSectionClicked(func(index int) {
		// Get the current sort order for the column
		currentOrder := sortOrders[index]

		// Toggle the sort order
		if currentOrder == core.Qt__AscendingOrder {
			sortOrders[index] = core.Qt__DescendingOrder
		} else {
			sortOrders[index] = core.Qt__AscendingOrder
		}

		// Sort the table by the column with the new sort order
		tableWidget.SortByColumn(index, sortOrders[index])
	})

	// Connect the itemClicked signal
	tableWidget.ConnectItemClicked(func(item *widgets.QTableWidgetItem) {
		// Get the row index of the clicked item
		row := item.Row()

		// Select the entire row
		tableWidget.SetRangeSelected(widgets.NewQTableWidgetSelectionRange2(row, 0, row, tableWidget.ColumnCount()-1), true)
	})

	// Create the right-click menu
	tableWidget.SetContextMenuPolicy(core.Qt__CustomContextMenu)

	setTableContextMenu(tableWidget)

	return tableWidget
}

func setTableContextMenu(tableWidget *widgets.QTableWidget) {
	contextMenu := widgets.NewQMenu(nil)
	copyUserNameAction := contextMenu.AddAction("Copy User Name \tCtrl+B+C")
	copyUserNameAction.SetShortcut(gui.NewQKeySequence2("Ctrl+B", gui.QKeySequence__NativeText))
	//copyUserNameAction.SetShortcut(widgets.QKeySequence_fromString("Ctrl+O"))

	copyPasswordAction := contextMenu.AddAction("Copy Password")
	copyPasswordAction.SetShortcut(gui.NewQKeySequence2("Ctrl+C", gui.QKeySequence__NativeText))

	copyPasswordAction.SetMenuRole(widgets.QAction__TextHeuristicRole) // Show shortcut in the context menu
	//copyPasswordAction.

	urlsMenu := contextMenu.AddMenu2("URS(S)")
	urlsMenu.AddAction("Open")
	urlsMenu.AddAction("Copy to ClipBoard")
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
		initDetailWidget(tableWidget)
	})
	copyPasswordAction.ConnectTriggered(func(bool) {
		initDetailWidget(tableWidget)
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

func initMenuBar(window *widgets.QMainWindow) {
	menuBar := window.MenuBar()

	functions.InitFileMenu(menuBar, window)
	functions.InitEditMenu(menuBar, window)
	functions.InitHelpMenu(menuBar, window)
}

func initToolbar(window *widgets.QMainWindow) {
	toolBar := widgets.NewQToolBar("Toolbar", window)

	// Add tool buttons to the toolbar
	newToolButton := widgets.NewQToolButton(nil)
	newToolButton.SetText("New")
	newToolIcon := window.Style().StandardIcon(widgets.QStyle__SP_FileIcon, nil, nil)
	newToolButton.SetIcon(newToolIcon)
	//.
	//newToolButton.SetIcon(gui.NewQIcon5("tr\\docs\\favicon.ico"))

	// Add tool buttons to the toolbar
	openToolButton := widgets.NewQToolButton(nil)
	openToolButton.SetToolTip("Open")

	openToolButton.SetIcon(gui.NewQIcon5("Ext\\Images_Client_16\\C49_Folder_Blue_Open.png"))

	/*newAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-new"), "New", nil)
	openAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-open"), "Open", nil)
	saveAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-save"), "Save", nil)

	// Add the actions to the toolbar
	toolBar.AddActions([]*widgets.QAction{newAction, openAction, saveAction})*/

	toolBar.AddWidget(newToolButton)
	toolBar.AddWidget(openToolButton)

	// Add the toolbar to the main window
	window.AddToolBar2(toolBar)
}
