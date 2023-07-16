package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"math/rand"
	"os"
	"time"
)

type EntryTab struct {
	FormLayout         *widgets.QFormLayout
	FirstNameLineEdit  *widgets.QLineEdit
	LastNameLineEdit   *widgets.QLineEdit
	UserNameEdit       *widgets.QLineEdit
	PasswordEdit       *widgets.QLineEdit
	RepeatPasswordEdit *widgets.QLineEdit
	ProgressBar        *widgets.QProgressBar
	URLEdit            *widgets.QLineEdit
	NotesEdit          *widgets.QTextEdit
	DateTimeEdit       *widgets.QDateTimeEdit
}

func clearChildItems(item *widgets.QTreeWidgetItem) {
	for item.ChildCount() > 0 {
		item.TakeChild(0)
	}
}
func (entry *EntryTab) InitEntryTab2(entryTab *widgets.QWidget) {
	// Create the entry tab struct
	//entryTabWidget := widgets.NewQWidget(nil, 0)

	// Create the form layout
	entry.FormLayout = widgets.NewQFormLayout(entryTab)

	// Create and add widgets to the form layout
	entry.FirstNameLineEdit = widgets.NewQLineEdit(nil)
	entry.LastNameLineEdit = widgets.NewQLineEdit(nil)

	nameLayout := widgets.NewQHBoxLayout2(nil)
	nameLayout.AddWidget(entry.FirstNameLineEdit, 0, core.Qt__AlignLeft)
	nameLayout.AddWidget(entry.LastNameLineEdit, 0, core.Qt__AlignLeft)
	label2 := widgets.NewQLabel2("Title:", nil, 0)
	entry.FormLayout.AddRow2(label2, nameLayout)

	entry.UserNameEdit = widgets.NewQLineEdit(nil)
	entry.FormLayout.AddRow3("User name:", entry.UserNameEdit)

	entry.PasswordEdit = widgets.NewQLineEdit(nil)
	entry.FormLayout.AddRow3("Password:", entry.PasswordEdit)

	entry.RepeatPasswordEdit = widgets.NewQLineEdit(nil)
	entry.FormLayout.AddRow3("Repeat:", entry.RepeatPasswordEdit)

	entry.ProgressBar = widgets.NewQProgressBar(nil)
	entry.ProgressBar.SetRange(0, 100)
	entry.FormLayout.AddRow3("Quality:", entry.ProgressBar)

	entry.URLEdit = widgets.NewQLineEdit(nil)
	entry.FormLayout.AddRow3("URL:", entry.URLEdit)

	entry.NotesEdit = widgets.NewQTextEdit(nil)
	entry.NotesEdit.Resize2(300, 200)
	entry.FormLayout.AddRow3("Notes:", entry.NotesEdit)

	entry.DateTimeEdit = widgets.NewQDateTimeEdit(nil)
	entry.FormLayout.AddRow3("Expires:", entry.DateTimeEdit)

	button := widgets.NewQPushButton2("Get DateTime", nil)
	button.ConnectClicked(func(checked bool) {
		selectedDateTime := entry.DateTimeEdit.DateTime().ToString("2006-01-02 15:04:05")
		entry.DateTimeEdit.SetDateTime(core.QDateTime_CurrentDateTime())
		fmt.Println("Selected DateTime:", selectedDateTime)
	})

	entry.FormLayout.AddRow3("Change datetime:", button)
}

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Create the main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("KeePass")
	// Create the menu bar
	initMenuBar(window)
	// Create the QKeySequence for the shortcut
	//shortcut := gui.NewQKeySequence2("core.Qt__Key_H", "")

	// Set the shortcut for the "About" action
	//aboutAction.SetShortcut(gui.QKeySequence_ITF().QKeySequence_PTR())

	// Create the toolbar with a title
	initToolbar(window)

	initMainContent(window)

	// 创建状态栏
	statusBar := widgets.NewQStatusBar(window)
	window.SetStatusBar(statusBar)

	// 在状态栏中显示文本
	statusBar.ShowMessage("Ready", 0)
	// Show the main window
	window.Resize2(800, 600)

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

func restTable(tableWidget *widgets.QTableWidget) {
	tableWidget.SetRowCount(2)
	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(0, 0, widgets.NewQTableWidgetItem2("搜狐", 0))
	tableWidget.SetItem(0, 1, widgets.NewQTableWidgetItem2("sohu", 0))
	tableWidget.SetItem(0, 2, widgets.NewQTableWidgetItem2("2.37", 0))
	tableWidget.SetItem(1, 0, widgets.NewQTableWidgetItem2("新浪", 0))
	tableWidget.SetItem(1, 1, widgets.NewQTableWidgetItem2("sina", 0))
	tableWidget.SetItem(1, 2, widgets.NewQTableWidgetItem2("1.34", 0))
}

func restTable2(tableWidget *widgets.QTableWidget) {
	tableWidget.SetRowCount(3)
	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(0, 0, widgets.NewQTableWidgetItem2("头条", 0))
	tableWidget.SetItem(0, 1, widgets.NewQTableWidgetItem2("toutiao", 0))
	tableWidget.SetItem(0, 2, widgets.NewQTableWidgetItem2("2.37", 0))
	tableWidget.SetItem(1, 0, widgets.NewQTableWidgetItem2("抖音", 0))
	tableWidget.SetItem(1, 1, widgets.NewQTableWidgetItem2("douyin", 0))
	tableWidget.SetItem(1, 2, widgets.NewQTableWidgetItem2("1.34", 0))
	tableWidget.SetItem(2, 0, widgets.NewQTableWidgetItem2("西瓜视频", 0))
	tableWidget.SetItem(2, 1, widgets.NewQTableWidgetItem2("西瓜", 0))
	tableWidget.SetItem(2, 2, widgets.NewQTableWidgetItem2("1.34", 0))
}

func initKeePass(tableWidget *widgets.QTableWidget) {
	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("supersecret")
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()

	// Note: This is a simplified example and the groups and entries will depend on the specific file.
	// bound checking for the slices is recommended to avoid panics.

	entries := db.Content.Root.Groups[0].Groups[0].Entries

	db.LockProtectedEntries()

	tableWidget.Clear()
	for i, entry := range entries {
		tableWidget.SetRowCount(tableWidget.RowCount() + 1)
		tableWidget.SetItem(i, 0, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
		tableWidget.SetItem(i, 1, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
		tableWidget.SetItem(i, 2, widgets.NewQTableWidgetItem2(entry.GetPassword(), 0))
		fmt.Println(entry.GetTitle())
		fmt.Println(entry.GetPassword())
	}

}

func initKeePassItem(qTreeWidgetItem *widgets.QTreeWidgetItem, tableWidget *widgets.QTableWidget) {
	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("supersecret")
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()
	index := qTreeWidgetItem.IndexOfChild(qTreeWidgetItem.Parent())
	// Note: This is a simplified example and the groups and entries will depend on the specific file.
	// bound checking for the slices is recommended to avoid panics.

	entries := db.Content.Root.Groups[0].Groups[index].Entries

	db.LockProtectedEntries()

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

func reAddTableItem(entry *EntryTab, tableWidget *widgets.QTableWidget) {
	tableWidget.SetRowCount(tableWidget.RowCount() + 1)
	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(tableWidget.RowCount()-1, 0, widgets.NewQTableWidgetItem2(entry.UserNameEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 1, widgets.NewQTableWidgetItem2(entry.PasswordEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 2, widgets.NewQTableWidgetItem2("2.37", 0))

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
	//groupMap := make(map[string]int, 10) //创建map
	//groupMap[group.Name] = i
	for i, rootGroup := range rootGroups {
		fmt.Println(i, "rootGroup:", rootGroup.Name)
		// Create the root item
		rootItem := widgets.NewQTreeWidgetItem4(treeWidget, []string{rootGroup.Name, "1.1"}, 0)
		rootItem.SetExpanded(true) // Set the root item initially expanded
		groups := rootGroup.Groups
		//group4 := rootGroup.Groups[4]
		for i, group := range groups {
			fmt.Println(i, "subGroup:", group.Name)
			treeItem := widgets.NewQTreeWidgetItem2([]string{group.Name}, 0)
			search1(group, 1, 1)

			rootItem.AddChild(treeItem)
			/*entries := group.Entries

			for i, entry := range entries {
				fmt.Println(i, entry.GetTitle(), entry.GetPassword())
				/*	fmt.Println(entry.GetTitle())
					fmt.Println(entry.GetPassword())
			}
			*/
		}
		// Add child items to the root item

		treeWidget.InsertTopLevelItem(i, rootItem)
	}
	// Create the root item

	// Set the root item as the top-level item of the tree widget

	treeWidget.SetHeaderHidden(true)

	// Connect the itemClicked signal of the tree widget
	treeWidget.ConnectItemClicked(func(item *widgets.QTreeWidgetItem, column int) {
		fmt.Println(item.Text(0), "点击了")

		parentItem := item.Parent()
		//item.AddChild(widgets.NewQTreeWidgetItem2([]string{"group.Name"}, 0))

		if parentItem == nil {
			// Clicked item is a top-level item
			topLevelIndex := treeWidget.IndexOfTopLevelItem(item)
			fmt.Printf("Clicked top-level item: Level %d, Index %d\n", 0, topLevelIndex)
		} else {
			// Clicked item is a child item
			//topLevelItem := treeWidget.InvisibleRootItem()
			topLevelIndex := treeWidget.IndexOfTopLevelItem(item)

			childIndex := parentItem.IndexOfChild(item)
			fmt.Printf("aaaaaaaaaaaaaaaaClicked child item: Level %d, Index %d\n", topLevelIndex, childIndex)
			//db.Content.Root.Groups[0].Groups[1]
			//search1(db, topLevelIndex, childIndex)
			rootGroups := db.Content.Root.Groups
			//groupMap := make(map[string]int, 10) //创建map
			//groupMap[group.Name] = i
			for _, rootGroup := range rootGroups {
				fmt.Println("level:", topLevelIndex, ",当前group:", rootGroup.Name)
				// Create the root item
				groups := rootGroup.Groups
				//group4 := rootGroup.Groups[4]

				for m, group := range groups {
					//fmt.Println(topLevelIndex, "当前点击是分组:", group.Name)
					if childIndex == m {

						clearChildItems(item)
						for _, grp := range group.Groups {

							item.AddChild(widgets.NewQTreeWidgetItem2([]string{grp.Name}, 0))
							fmt.Println(m, "level:", topLevelIndex, "-----------childIndex:", childIndex)

						}
					}

				}

			}

		}

		//index := treeWidget.CurrentIndex().Row()
		//it := treeWidget.SelectedItems()[0]
		//selectedItems := treeWidget.SelectedItems()
		// Iterate over the selected items
		//for _, item := range selectedItems {
		// Perform operations on the selected item
		//	item.
		//}
		/*if len(items) > 0 {
			item := items[0] //取第一个选中节点
			//node := item.Data(0, widgets.Qt__UserRole).(myNodeType) // 强制转换为我的节点类型
			itemData := item.Data(0, 0).ToString()
			fmt.Println(itemData)
			//item.Data()
			// 使用节点对象
			//fmt.Println(node.Name)
		}*/

		//initKeePassItem(it, tableWidget)
		//expanded := it.IsExpanded() // Get the current expansion state
		//it.SetExpanded(!expanded)   // Toggle the expansion state

	})
	return treeWidget
}

func search1(group gokeepasslib.Group, level, index int) int {

	rootGroups := group.Groups

	for i, rootGroup := range rootGroups {
		fmt.Println(i, "files-------:", level, "-", index, "-", rootGroup.Name)
	}

	/*	files, err := ioutil.ReadDir(path)
		fmt.Println("files-------:", files)
		if err != nil {
			fmt.Println("目录读取失败！", err.Error())
			return matches
		}
		if len(files) <= 0 {
			return matches
		}
		for _, file := range files {
			name := file.Name()
			fmt.Println("name-----:", name)
			if name == queryName {
				matches++
			}
			if file.IsDir() {
				dir := path + "/" + name
				if path == "/" {
					dir = path + name
				}
				search1(dir, queryName)
			}
		}*/
	return 0
}

func initDetailWidget(tableWidget *widgets.QTableWidget) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	tabWidget := widgets.NewQTabWidget(dialog)
	entryTab, advancedTab := initTabWidget(tabWidget)
	//entryTabWidget := widgets.NewQWidget(nil, 0)
	entry := &EntryTab{}
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
}

func initKeePassImage() *widgets.QLabel {
	imageLabel := widgets.NewQLabel(nil, 0)
	imagePixmap := gui.NewQPixmap3("D:\\workspace_go\\KeePassGo\\basic\\Hello\\img\\keepass.png", "", core.Qt__AutoColor)
	imageLabel.SetPixmap(imagePixmap)
	return imageLabel
}

func initBottomButton(entryTab *EntryTab, tableWidget *widgets.QTableWidget, tabWidget *widgets.QTabWidget, dialog *widgets.QDialog) *widgets.QHBoxLayout {
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
		reAddTableItem(entryTab, tableWidget)

		dialog.Close()
	})

	cancelButton.ConnectClicked(func(bool) {
		// Code to handle cancelButton click event
		fmt.Println("Cancel Button clicked")
		dialog.Close()
	})
	return hBoxLayout
}

func initTabWidget(tabWidget *widgets.QTabWidget) (*widgets.QWidget, *widgets.QWidget) {
	// Create and add tabs to the tab widget
	entryTab := widgets.NewQWidget(nil, 0)
	advancedTab := widgets.NewQWidget(nil, 0)
	propertiesTab := widgets.NewQWidget(nil, 0)
	autoTypeTab := widgets.NewQWidget(nil, 0)
	historyTab := widgets.NewQWidget(nil, 0)

	tabWidget.AddTab(entryTab, "Entry")
	tabWidget.AddTab(advancedTab, "Advanced")
	tabWidget.AddTab(propertiesTab, "Properties")
	tabWidget.AddTab(autoTypeTab, "Auto-Type")
	tabWidget.AddTab(historyTab, "History")
	tabWidget.Resize2(700, 400)
	return entryTab, advancedTab
}

func initTabWidget2(entryTab *widgets.QWidget, tabWidget *widgets.QTabWidget) {
	// Create and add tabs to the tab widget
	tabWidget.AddTab(entryTab, "Entry")

}

func initEntryTab(entryTab *widgets.QWidget) {
	//widgets.newqla
	// Create and add widgets to the form layout
	//nameLabel := widgets.NewQLabel2("Name:", nil, 0)
	// Create the form layout
	formLayout := widgets.NewQFormLayout(entryTab)
	firstNameLineEdit := widgets.NewQLineEdit(nil)
	lastNameLineEdit := widgets.NewQLineEdit(nil)

	nameLayout := widgets.NewQHBoxLayout2(nil)
	nameLayout.AddWidget(firstNameLineEdit, 0, core.Qt__AlignLeft)
	nameLayout.AddWidget(lastNameLineEdit, 0, core.Qt__AlignLeft)
	label2 := widgets.NewQLabel2("Title:", nil, 0)
	//formLayout.AddRow3("nameLabel", nameLayout.Widget())
	formLayout.AddRow2(label2, nameLayout)
	userNameEdit := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("User name:", userNameEdit)

	passwordEdit := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("Password:", passwordEdit)

	repeatPasswordEdit := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("Repeat:", repeatPasswordEdit)

	// Create a progress bar
	progressBar := widgets.NewQProgressBar(nil)
	progressBar.SetRange(0, 100)

	// Create a palette for the progress bar
	palette := progressBar.Palette()
	//palette.SetColor(gui.QPalette__Base, core.Qt__GlobalColor(gui.QPalette__Dark))
	// Create a color gradient from orange to green
	gradient := gui.NewQLinearGradient3(0, 0, 1, 0)
	gradient.SetColorAt(0.0, gui.NewQColor3(255, 165, 0, 0)) // Orange
	gradient.SetColorAt(1.0, gui.NewQColor3(0, 128, 0, 0))   // Green

	// Create a brush with the gradient
	brush := gui.NewQBrush10(gradient)

	// Set the color gradient as the background of the progress bar
	palette.SetBrush(gui.QPalette__Highlight, brush)
	progressBar.SetPalette(palette)

	// Create a line edit for entering the password
	passwordEdit.ConnectTextChanged(func(text string) {
		// Calculate the password complexity score
		complexity := calculatePasswordComplexity(text)

		// Set the value of the progress bar based on the complexity score
		progressBar.SetValue(complexity)
	})

	formLayout.AddRow3("Quality:", progressBar)

	urlEdit := widgets.NewQLineEdit(nil)
	formLayout.AddRow3("Url:", urlEdit)

	notesEdit := widgets.NewQTextEdit(nil)
	notesEdit.Resize2(300, 200)
	formLayout.AddRow3("Notes:", notesEdit)

	dateTimeEdit := widgets.NewQDateTimeEdit(nil)
	//dateEdit.enab
	formLayout.AddRow3("Expires:", dateTimeEdit)
	button := widgets.NewQPushButton2("Get DateTime", nil)
	button.ConnectClicked(func(checked bool) {
		selectedDateTime := dateTimeEdit.DateTime().ToString("2006-01-02 15:04:05")
		dateTimeEdit.SetDateTime(core.QDateTime_CurrentDateTime())
		fmt.Println("Selected DateTime:", selectedDateTime)
	})

	formLayout.AddRow3("chage datetime:", button)

}

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
	imagePixmap := gui.NewQPixmap3("D:\\workspace_go\\KeePassGo\\basic\\Hello\\img\\keepass.png", "", core.Qt__AutoColor)
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
	tableWidget.SetRowCount(2)

	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(0, 0, widgets.NewQTableWidgetItem2("KeePass", 0))
	tableWidget.SetItem(0, 1, widgets.NewQTableWidgetItem2("Status 1", 0))
	tableWidget.SetItem(0, 2, widgets.NewQTableWidgetItem2("2.37", 0))
	tableWidget.SetItem(1, 0, widgets.NewQTableWidgetItem2("keePassLicC", 0))
	tableWidget.SetItem(1, 1, widgets.NewQTableWidgetItem2("Status 2", 0))
	tableWidget.SetItem(1, 2, widgets.NewQTableWidgetItem2("1.34", 0))

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
	addItemAction := contextMenu.AddAction("Add Item")
	deleteItemAction := contextMenu.AddAction("Delete Item")
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

	initFileMenu(menuBar, window)

	initHelpMenu(menuBar, window)
}

func initToolbar(window *widgets.QMainWindow) {
	toolBar := widgets.NewQToolBar("Toolbar", window)

	// Add tool buttons to the toolbar
	newToolButton := widgets.NewQToolButton(nil)
	newToolButton.SetText("New")

	// Add tool buttons to the toolbar
	openToolButton := widgets.NewQToolButton(nil)
	openToolButton.SetText("Open")

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

func initHelpMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the help menu
	helpMenu := menuBar.AddMenu2("Help")

	// Create the "About" action for the help menu
	helpAction := helpMenu.AddAction("Help")
	//helpAction.SetShortcuts2(core.NewQKeySequence2("Ctrl+O", core.QKeySequence__NativeText))
	//helpAction.SetShortcuts(gui.NewQKeySequence("Ctrl+O", core.QKeySequence__NativeText)))
	// Connect the "About" action to its triggered event
	helpAction.ConnectTriggered(func(checked bool) {

		dialog := widgets.NewQDialog(window, 0)
		dialog.SetWindowTitle("QDialogButtonBox Example")

		// Create the button box
		buttonBox := widgets.NewQDialogButtonBox(dialog)
		okButton := buttonBox.AddButton3(widgets.QDialogButtonBox__Ok)
		cancelButton := buttonBox.AddButton3(widgets.QDialogButtonBox__Cancel)

		// Connect the button box's accepted signal
		buttonBox.ConnectAccepted(func() {
			fmt.Println("OK button clicked")
			dialog.Accept()
		})

		// Connect the button box's rejected signal
		buttonBox.ConnectRejected(func() {
			fmt.Println("Cancel button clicked")
			dialog.Reject()
		})

		// Set the button text
		okButton.SetText("OK")
		cancelButton.SetText("Cancel")

		// Create the layout
		layout := widgets.NewQVBoxLayout2(dialog)
		layout.AddWidget(buttonBox, 0, core.Qt__AlignCenter)

		// Set the layout for the dialog
		dialog.SetLayout(layout)

		// Show the main window and the dialog

		dialog.Exec()

		//url := core.NewQUrl3("https://keepass.info/help/base/index.html", core.QUrl__StrictMode)
		//gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput("https://keepass.info/help/base/index.html"))

		// Open the URL in the default web browser
		//widgets.QDesktopServices_OpenUrl(url)
		//widgets.QMessageBox_About(window, "帮助", "请参考:https://www.csdn.net")
	})

	helpMenu.AddSeparator()
	helpMenu.AddAction("KeePass Website").ConnectTriggered(func(checked bool) {
		gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput("https://keepass.info/"))
	})
	helpMenu.AddAction("Donate...").ConnectTriggered(func(checked bool) {
		gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput("https://keepass.info/donate.html"))
	})

	helpMenu.AddSeparator()
	// Create the "About" action for the help menu
	aboutAction := helpMenu.AddAction("About KeePassGo")
	// Connect the "About" action to its triggered event
	aboutAction.ConnectTriggered(func(checked bool) {
		// 创建自定义对话框
		dialog := widgets.NewQDialog(nil, 0)
		dialog.SetWindowTitle("About")

		// 创建版权文本 QLabel
		copyrightLabel := widgets.NewQLabel2("Copyright Text", nil, 0)

		// 创建 KeePass 官网链接 QLabel
		websiteLabel := widgets.NewQLabel2("<a href=\"https://keepass.info\">KeePass website</a>", nil, 0)
		websiteLabel.SetOpenExternalLinks(true)

		// 创建 Component/Status/Version 表格
		table := widgets.NewQTableWidget2(2, 3, nil)
		table.SetHorizontalHeaderLabels([]string{"Component", "Status", "Version"})
		table.SetItem(0, 0, widgets.NewQTableWidgetItem2("KeePass", 0))
		table.SetItem(0, 1, widgets.NewQTableWidgetItem2("Status 1", 0))
		table.SetItem(0, 2, widgets.NewQTableWidgetItem2("2.37", 0))
		table.SetItem(1, 0, widgets.NewQTableWidgetItem2("keePassLicC", 0))
		table.SetItem(1, 1, widgets.NewQTableWidgetItem2("Status 2", 0))
		table.SetItem(1, 2, widgets.NewQTableWidgetItem2("1.34", 0))

		// 创建 OK 按钮
		okButton := widgets.NewQPushButton2("OK", nil)

		// 连接 OK 按钮的点击事件
		okButton.ConnectClicked(func(bool) {
			dialog.Close()
		})

		// 创建主布局
		layout := widgets.NewQVBoxLayout()

		// 添加部件到主布局
		layout.AddWidget(copyrightLabel, 0, core.Qt__AlignCenter)
		layout.AddSpacing(10)
		layout.AddWidget(websiteLabel, 0, core.Qt__AlignCenter)
		layout.AddSpacing(10)
		layout.AddWidget(table, 0, 0)

		/*gridLayout := widgets.NewQGridLayout(nil)
		gridLayout.AddWidget(copyrightLabel)
		gridLayout.AddWidget(copyrightLabel)
		gridLayout.AddWidget(copyrightLabel)
		gridLayout.AddWidget(copyrightLabel)
		gridLayout.AddWidget(copyrightLabel)
		gridLayout.AddWidget(copyrightLabel)*/

		// 创建网格布局
		gridLayout := widgets.NewQGridLayout(nil)

		// 门户网站连接列表
		websites := []string{
			"https://www.google.com",
			"https://www.github.com",
			"https://www.microsoft.com",
			"https://www.apple.com",
			"https://www.openai.com",
			"https://www.spotify.com",
			"https://www.amazon.com",
			"https://www.facebook.com",
			"https://www.twitter.com",
		}

		// 创建随机数生成器
		rand.Seed(time.Now().UnixNano())

		// 在网格布局中随机放置门户网站连接
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// 从门户网站连接列表中随机选择一个连接
				randomWebsite := websites[rand.Intn(len(websites))]

				// 创建 QLabel 和 QUrl
				label := widgets.NewQLabel(nil, 0)
				url := core.NewQUrl3(randomWebsite, core.QUrl__TolerantMode)

				// 设置 QLabel 的文本和打开外部链接
				label.SetText(randomWebsite)
				label.SetOpenExternalLinks(true)

				// 将 QLabel 添加到网格布局的指定位置
				gridLayout.AddWidget2(label, i, j, core.Qt__AlignLeft)

				// 释放 QUrl
				url.DestroyQUrl()
			}
		}

		// 创建底部布局
		bottomLayout := widgets.NewQHBoxLayout2(nil)
		bottomLayout.AddSpacing(10)
		bottomLayout.AddStretch(1)
		bottomLayout.AddWidget(okButton, 0, core.Qt__AlignRight|core.Qt__AlignBottom)

		// 添加底部布局到主布局
		//layout.AddSpacing(10)
		//layout.AddLayout(bottomLayout)

		layout.AddSpacing(10)
		layout.AddLayout(bottomLayout, 0)
		layout.AddSpacing(10)
		layout.AddLayout(gridLayout, 0)
		// 设置主布局为对话框的布局
		dialog.SetLayout(layout)

		// 显示对话框
		dialog.Exec()
	})

}

func initFileMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	fileMenu := menuBar.AddMenu2("File")

	// Create actions for the file menu
	newAction := fileMenu.AddAction("&New...")
	// Connect the actions and tool buttons to their respective triggered events
	newAction.ConnectTriggered(func(checked bool) {
		doNewAction(window)
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
	//shortcut := gui.NewQKeySequence()

	// Set the key code for the "Open" action
	//shortcut.FromString("ctrl+A", gui.QKeySequence__NativeText)

	//openAction.SetShortcuts([]*gui.QKeySequence{shortcut})

	// Set the shortcut for the "Open" action
	//openAction.SetShortcuts([]*gui.QKeySequence{shortcut})

	// Connect the actions and tool buttons to their respective triggered events
	/*openAction.ConnectTriggered(func(checked bool) {
		widgets.NewQFileDialog2(window, "打开", "d:", "*.*").Show()
	})*/
}

func doNewAction(window *widgets.QMainWindow) {
	msgBox := widgets.NewQMessageBox(nil)
	msgBox.SetWindowTitle("Message Box")
	msgBox.SetText("This is a message box.")
	msgBox.SetInformativeText("This is a message box with text information.")
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

	// Show the message box
	msgBox.Exec()
}
