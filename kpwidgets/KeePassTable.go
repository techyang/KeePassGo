package kpwidgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
)

type KeePassTable struct {
	*widgets.QTableWidget
}

func NewKeePassTable() *KeePassTable {
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
	//passwordDelegate := NewPasswordDelegate()
	//tableWidget.SetItemDelegateForColumn(0, passwordDelegate)

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

	passwordDelegate := NewPasswordDelegate()
	tableWidget.SetItemDelegateForColumn(0, passwordDelegate)

	keePassTable := &KeePassTable{
		QTableWidget: tableWidget,
	}
	//SetTableContextMenu(tableWidget)
	keePassTable.SetTableContextMenu()
	return keePassTable
}

func (keePassTable *KeePassTable) SetTableContextMenu() {
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

	moveTopAction.ConnectTriggered(func(checked bool) {
		moveTop(keePassTable)
	})

	setMoveTopAction(keePassTable, moveTopAction)
	setMoveUpAction(keePassTable, moveUpAction)
	setMoveDownAction(keePassTable, moveDownAction)
	setMoveBottomAction(keePassTable, moveBottomAction)

	setCopyUserNameAction(keePassTable, copyUserNameAction)
	setCopyPasswordAction(keePassTable, copyPasswordAction)

	setOpenUrlAction(keePassTable, openUrlAction)

	setCopyUrlAction(keePassTable, copyUrlAction)
	performAutoTypeAction.ConnectTriggered(func(bool) {
		initDetailWidget(keePassTable)
	})
	setEditOrViewEntryAction(keePassTable, editOrViewEntryAction)
	setDuplicateAction(keePassTable, duplicateAction)

	selectAllAction.ConnectTriggered(func(bool) {
		keePassTable.SelectAll()
	})

	/*clipbordAction := contextMenu.AddMenu2("Copy User Name")
	clipbordAction := contextMenu.AddAction("Copy User Name")

	copyUserNameAction := contextMenu.AddAction("Copy User Name")

	copyUserNameAction := contextMenu.AddAction("Copy User Name")

	copyUserNameAction := contextMenu.AddAction("Copy User Name")
	*/

	keePassTable.ConnectCustomContextMenuRequested(func(pos *core.QPoint) {
		contextMenu.Exec2(keePassTable.MapToGlobal(pos), nil)
	})

	// Connect the triggered signal of the menu actions
	addItemAction.ConnectTriggered(func(bool) {
		initDetailWidget(keePassTable)
	})

	deleteItemAction.ConnectTriggered(func(bool) {

		row := keePassTable.CurrentRow()
		if row > 0 {
			keePassTable.RemoveRow(row)
		}

		// Get the selection model from the table view
		// Get the selection model from the table view
		//selectionModel := keePassTable.SelectionModel()
		//selectedRows := selectionModel.Selection()
		//keePassTable.Model().RemoveRow(0, core.NewQModelIndex())

		//selectedIndexes := selectionModel.SelectedRows()

		/*qModelIndex := selectionModel.SelectedRows(0).
		qModelIndex.
		keePassTable.*/

		/*for _, index := range selectedRows {
			keePassTable.Model().RemoveRow(index.Row(), core.NewQModelIndex())
			fmt.Println("第", index, "行删除了")
		}*/

	})
}

func moveTop(keePassTable *KeePassTable) {
	row := keePassTable.CurrentRow()
	if row > 0 {
		rowData := getRowData(keePassTable, row)
		keePassTable.RemoveRow(row)
		newRow := 0
		keePassTable.InsertRow(newRow)
		setTableRowData(keePassTable, newRow, rowData)
	}
}

func (tableWidget *KeePassTable) setTableItems(group *gokeepasslib.Group) {
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
