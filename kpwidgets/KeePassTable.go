package kpwidgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
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
	//SetTableContextMenu(tableWidget)
	SetTableContextMenu(tableWidget)

	keePassTable := &KeePassTable{
		QTableWidget: tableWidget,
	}

	return keePassTable
}
