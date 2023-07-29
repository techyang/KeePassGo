package functions

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

type PasswordDelegate struct {
	widgets.QStyledItemDelegate
}

func NewPasswordDelegate() *PasswordDelegate {
	return &PasswordDelegate{}
}

func (delegate *PasswordDelegate) DisplayText(value *core.QVariant, locale *core.QLocale) string {
	// If the value is a string, return "***"
	if value.Type() == core.QVariant__String {
		return "***"
	}

	// Otherwise, call the base class method to display the default text
	return delegate.QStyledItemDelegate.DisplayText(value, locale)
}

func InitMainContent(window *widgets.QMainWindow) {
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

func initTableWidget() *widgets.QTableWidget {
	// Create a QTableWidget
	// Create the table widget
	tableWidget := widgets.NewQTableWidget(nil)
	tableWidget.SetColumnCount(5)

	// Set the header labels
	headerLabels := []string{"Title", "User Name", "Password", "URL", "Notes"}
	tableWidget.SetHorizontalHeaderLabels(headerLabels)

	// when click one table item, select the whole columns of the row
	tableWidget.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)

	// Enable sorting
	tableWidget.SetSortingEnabled(true)

	// Set the password delegate for the second column
	passwordDelegate := NewPasswordDelegate()
	tableWidget.SetItemDelegateForColumn(0, passwordDelegate)

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

	// Create the right-click menu
	tableWidget.SetContextMenuPolicy(core.Qt__CustomContextMenu)

	SetTableContextMenu(tableWidget)

	return tableWidget
}

func initTreeWidget(tableWidget *widgets.QTableWidget) *widgets.QTreeWidget {
	treeWidget := widgets.NewQTreeWidget(nil)
	//treeWidget.SetHeaderLabels([]string{"yangwl"})
	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
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

func TableItemClicked(tableWidget *widgets.QTableWidget, treeWidget *widgets.QTreeWidget, rootGroups []gokeepasslib.Group) {

	treeWidget.ConnectItemClicked(func(item *widgets.QTreeWidgetItem, column int) {

		groupUUID := item.Data(1, 0).ToString()
		fmt.Println(item.Text(0), "点击了", groupUUID)
		tableWidget.SetObjectName(groupUUID)
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

	passwordDelegate := NewPasswordDelegate()
	tableWidget.SetItemDelegateForColumn(0, passwordDelegate)
}

func setTableItems(group *gokeepasslib.Group, tableWidget *widgets.QTableWidget) {
	// Set the password delegate for the second column
	passwordDelegate := NewPasswordDelegate()
	tableWidget.SetItemDelegateForColumn(0, passwordDelegate)
	for i, entry := range group.Entries {
		username := entry.Get("UserName").Value.Content
		url := entry.Get("URL").Value.Content
		note := entry.Get("Notes").Value.Content
		tableWidget.SetRowCount(i + 1)
		tableWidget.SetItem(i, 0, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
		tableWidget.SetItem(i, 1, widgets.NewQTableWidgetItem2(username, 0))

		passwordItem := widgets.NewQTableWidgetItem2(entry.GetPassword(), 0)
		//passwordItem.SetFlags(core.Qt__ItemIsSelectable | core.Qt__ItemIsEditable)
		//passwordItem.SetFlags(passwordItem.Flags() | core.Qt__ItemIsUserCheckable)
		//passwordItem.SetCheckState(core.Qt__Checked)

		tableWidget.SetItem(i, 2, passwordItem)
		tableWidget.SetItem(i, 3, widgets.NewQTableWidgetItem2(url, 0))
		tableWidget.SetItem(i, 4, widgets.NewQTableWidgetItem2(note, 0))
	}
}

func InitToolbar(window *widgets.QMainWindow) {
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
