package kpwidgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CommonTableWidget struct {
	*widgets.QTableWidget
}

func NewCommonTableWidget(headerLabels []string) *CommonTableWidget {
	// Create a QTableWidget
	// Create the table widget

	tableWidget := widgets.NewQTableWidget(nil)
	tableWidget.SetColumnCount(len(headerLabels))

	// Set the header labels
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

	keePassTable := &CommonTableWidget{
		QTableWidget: tableWidget,
	}
	//SetTableContextMenu(tableWidget)
	//keePassTable.SetTableContextMenu()
	return keePassTable
}

func (tableWidget *CommonTableWidget) MoveTop() {
	row := tableWidget.CurrentRow()
	if row > 0 {
		rowData := tableWidget.getRowData(row)
		tableWidget.RemoveRow(row)
		newRow := 0
		tableWidget.InsertRow(newRow)
		tableWidget.setTableRowData(newRow, rowData)
	}
}

func (tableWidget *CommonTableWidget) MoveBottom() {
	row := tableWidget.CurrentRow()
	if row < tableWidget.RowCount()-1 {
		rowData := tableWidget.getRowData(row)
		tableWidget.RemoveRow(row)
		newRow := tableWidget.RowCount()
		tableWidget.InsertRow(newRow)
		tableWidget.setTableRowData(newRow, rowData)
	}
}

func (tableWidget *CommonTableWidget) MoveUp() {
	row := tableWidget.CurrentRow()
	if row > 0 {
		rowData := tableWidget.getRowData(row)
		tableWidget.RemoveRow(row)
		newRow := row - 1
		tableWidget.InsertRow(newRow)
		tableWidget.setTableRowData(newRow, rowData)
	}
}

func (tableWidget *CommonTableWidget) MoveDown() {
	row := tableWidget.CurrentRow()
	if row < tableWidget.RowCount()-1 {
		rowData := tableWidget.getRowData(row)
		tableWidget.RemoveRow(row)
		newRow := row + 1
		tableWidget.InsertRow(newRow)
		tableWidget.setTableRowData(newRow, rowData)
	}
}

func (tableWidget *CommonTableWidget) setTableRowData(newRow int, rowData []string) {
	for column := 0; column < tableWidget.ColumnCount(); column++ {
		tableWidget.SetItem(newRow, column, widgets.NewQTableWidgetItem2(rowData[column], 0))
	}
	tableWidget.SelectRow(newRow)
}

// Function to retrieve the data of a specific row in a QTableWidget and store it in an array
func (tableWidget *CommonTableWidget) getRowData(row int) []string {
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
