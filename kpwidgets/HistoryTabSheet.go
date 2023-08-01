package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
)

type HistoryTabSheet struct {
	*widgets.QWidget
	TableWidget   *widgets.QTableWidget
	ViewButton    *widgets.QPushButton
	DeleteButton  *widgets.QPushButton
	RestoreButton *widgets.QPushButton
}

func NewHistoryTabSheet() *HistoryTabSheet {
	cw := &HistoryTabSheet{
		QWidget:       widgets.NewQWidget(nil, 0),
		TableWidget:   widgets.NewQTableWidget(nil),
		ViewButton:    widgets.NewQPushButton2("View", nil),
		DeleteButton:  widgets.NewQPushButton2("Delete", nil),
		RestoreButton: widgets.NewQPushButton2("Restore", nil),
	}

	// Layout
	vBoxLayout := widgets.NewQVBoxLayout2(cw)

	setHistoryTableWidget(cw.TableWidget, vBoxLayout)
	/*vBoxLayout.AddWidget(cw.TagsLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.TagsInput, 0, core.Qt__AlignLeft)*/

	twoChannelAutoTypeLayout := widgets.NewQHBoxLayout2(nil)
	twoChannelAutoTypeLayout.AddWidget(cw.ViewButton, 0, core.Qt__AlignLeft)
	twoChannelAutoTypeLayout.AddWidget(cw.DeleteButton, 0, core.Qt__AlignTop)

	spacer := widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	twoChannelAutoTypeLayout.AddSpacerItem(spacer)
	twoChannelAutoTypeLayout.AddWidget(cw.ViewButton, 0, core.Qt__AlignLeft)
	twoChannelAutoTypeLayout.AddWidget(cw.RestoreButton, 0, core.Qt__AlignTop)
	vBoxLayout.AddLayout(twoChannelAutoTypeLayout, 0)

	/*	cw.BackgroundColorPickerBtn.ConnectClicked(func(checked bool) {
		colorDialog := widgets.NewQColorDialog2(gui.NewQColor3(255, 255, 255, 255), nil)
		colorDialog.Exec()
		if colorDialog != nil {
			cw.BackgroundColorPickerBtn.SetStyleSheet("background-color: " + colorDialog.CurrentColor().Name())
		}
	})*/

	return cw
}

func setHistoryTableWidget(tableWidget *widgets.QTableWidget, vBoxLayout *widgets.QVBoxLayout) {

	tableWidget.SetColumnCount(4)
	tableWidget.SetFixedWidth(600)
	// Set the header labels
	headerLabels := []string{"Version", "Title", "User Name", "Size"}
	tableWidget.SetHorizontalHeaderLabels(headerLabels)
	vBoxLayout.AddWidget(tableWidget, 0, core.Qt__AlignLeft)
}

func (historyTabSheet *HistoryTabSheet) SetTableRowData2(kpEntry gokeepasslib.Entry) {
	if kpEntry.Histories != nil {
		for _, history := range kpEntry.Histories {
			for i, entry := range history.Entries {
				historyTabSheet.TableWidget.SetRowCount(i + 1)
				version := entry.Times.LastModificationTime.Time
				formattedTime := version.Format("2006-01-02 15:04:05")
				username := entry.Values[4].Value.Content
				size := "1 KB"
				historyTabSheet.TableWidget.SetItem(i, 0, widgets.NewQTableWidgetItem2(formattedTime, 0))
				historyTabSheet.TableWidget.SetItem(i, 1, widgets.NewQTableWidgetItem2(entry.GetTitle(), 0))
				historyTabSheet.TableWidget.SetItem(i, 2, widgets.NewQTableWidgetItem2(username, 0))
				historyTabSheet.TableWidget.SetItem(i, 3, widgets.NewQTableWidgetItem2(string(size), 0))
			}
		}
	}
	/*for column := 0; column < tableWidget.ColumnCount(); column++ {
		tableWidget.SetItem(newRow, column, widgets.NewQTableWidgetItem2(rowData[column], 0))
	}
	tableWidget.SelectRow(newRow)*/
}

func getHistoryTableButtonVLayout(tableWidget *widgets.QTableWidget) *widgets.QVBoxLayout {
	tableButtonVLayout := widgets.NewQVBoxLayout2(nil)
	addButton := widgets.NewQPushButton2("Add", nil)
	editButton := widgets.NewQPushButton2("Edit", nil)
	moveButton := widgets.NewQPushButton2("Move", nil)
	upButton := widgets.NewQPushButton2("Up", nil)
	downButton := widgets.NewQPushButton2("Down", nil)
	tableButtonVLayout.AddWidget(addButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(editButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(moveButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(upButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(downButton, 0, core.Qt__AlignLeft)

	tableWidget.ConnectItemClicked(func(item *widgets.QTableWidgetItem) {
		// Get the row index of the clicked item
		//row := item.Row()

		// Select the entire row
		//tableWidget.SetRangeSelected(widgets.NewQTableWidgetSelectionRange2(row, 0, row, tableWidget.ColumnCount()-1), true)
		tableWidget.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	})

	addButton.ConnectClicked(func(checked bool) {
		row := tableWidget.CurrentRow()
		tableWidget.InsertRow(row + 1)

		//tableWidget.SetRowCount(tableWidget.RowCount() + 1)
		if row >= 0 {
			fieldName := tableWidget.Item(row, 0).Text()
			fieldValue := tableWidget.Item(row, 1).Text()

			// Add your edit logic here...

			// For demonstration, we just print the values
			fmt.Printf("Editing row %d - Field Name: %s, Field Value: %s\n", row, fieldName, fieldValue)
		}
	})

	editButton.ConnectClicked(func(checked bool) {
		row := tableWidget.CurrentRow()
		if row >= 0 {
			fieldName := tableWidget.Item(row, 0).Text()
			fieldValue := tableWidget.Item(row, 1).Text()

			// Add your edit logic here...

			// For demonstration, we just print the values
			fmt.Printf("Editing row %d - Field Name: %s, Field Value: %s\n", row, fieldName, fieldValue)
		}
	})

	// Connect the button clicked signals to the slots
	moveButton.ConnectClicked(func(checked bool) {
		// Get the selected row index
		row := tableWidget.CurrentRow()
		if row >= 0 {
			tableWidget.RemoveRow(row)
			if tableWidget.RowCount() >= row+1 {
				tableWidget.SelectRow(row)
			}
		}
	})

	upButton.ConnectClicked(func(checked bool) {
		row := tableWidget.CurrentRow()

		if row > 0 {
			fieldName := tableWidget.Item(row, 0).Text()
			fieldValue := tableWidget.Item(row, 1).Text()
			tableWidget.RemoveRow(row)
			tableWidget.InsertRow(row - 1)
			tableWidget.SetItem(row-1, 0, widgets.NewQTableWidgetItem2(fieldName, 0))
			tableWidget.SetItem(row-1, 1, widgets.NewQTableWidgetItem2(fieldValue, 0))
			tableWidget.SelectRow(row - 1)
			/*tableWidget.InsertRow(row - 1)
			for column := 0; column < tableWidget.ColumnCount(); column++ {
				tableWidget.SetItem(row-1, column, tableWidget.Item(row, column))
			}*/
		}
	})

	downButton.ConnectClicked(func(checked bool) {
		row := tableWidget.CurrentRow()
		if row < tableWidget.RowCount()-1 {
			fieldName := tableWidget.Item(row, 0).Text()
			fieldValue := tableWidget.Item(row, 1).Text()
			tableWidget.RemoveRow(row)
			tableWidget.InsertRow(row + 1)
			tableWidget.SetItem(row+1, 0, widgets.NewQTableWidgetItem2(fieldName, 0))
			tableWidget.SetItem(row+1, 1, widgets.NewQTableWidgetItem2(fieldValue, 0))
			tableWidget.SelectRow(row + 1)
		}

		/*if row < tableWidget.RowCount()-1 {
			tableWidget.RemoveRow(row)
			tableWidget.InsertRow(row + 1)
			for column := 0; column < tableWidget.ColumnCount(); column++ {
				tableWidget.SetItem(row+1, column, tableWidget.Item(row, column))
			}
		}*/
	})

	return tableButtonVLayout
}
