package main

import (
	"fmt"
	"github.com/techyang/keepassgo/functions"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
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

		findGroupByName(groups, "公共111")
		//	findGroupByUUID(groups, "[97 82 122 100 116 72 67 110 76 107 87 79 110 82 90 57 119 81 103 101 85 81 61 61]")
		//
		buildGroupTree(rootItem, groups)

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

func buildGroupTree(parent *widgets.QTreeWidgetItem, groups []gokeepasslib.Group) {
	for _, group := range groups {
		txt, _ := group.UUID.MarshalText()
		fmt.Println("group.UUID -----------:", group.UUID)
		fmt.Println("group.UUID:", txt)
		treeItem := widgets.NewQTreeWidgetItem2([]string{group.Name}, 0)
		parent.AddChild(treeItem)
		buildGroupTree(treeItem, group.Groups)
	}
}

func findGroupByName(groups []gokeepasslib.Group, name string) *gokeepasslib.Group {
	for _, group := range groups {
		if group.Name == name {
			fmt.Println("找到的名称是:", group.Name)
			return &group
		}
		if foundGroup := findGroupByName(group.Groups, name); foundGroup != nil {
			return foundGroup
		}
	}
	return nil
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

/*
	func buildGroupTree(parent *widgets.QTreeWidgetItem, groups []*gokeepasslib.Group) {
		for _, group := range groups {
			treeItem := widgets.NewQTreeWidgetItem2([]string{group.Name}, 0)
			parent.AddChild(treeItem)
			//item := widgets.NewQTreeWidgetItem2(parent, []string{group.Name()})
			buildGroupTree(treeItem, group.Groups)
		}
	}
*/
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
	imagePixmap := gui.NewQPixmap3("D:\\workspace_go\\KeePassGo\\src\\Hello\\img\\keepass.png", "", core.Qt__AutoColor)
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

	functions.InitFileMenu(menuBar, window)
	functions.InitEditMenu(menuBar, window)
	functions.InitHelpMenu(menuBar, window)
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
