package kpwidgets

import (
	log "github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
)

type AdvancedTabSheet struct {
	*widgets.QWidget
	StringFieldsGroupBox         *widgets.QGroupBox
	StringFieldsTable            *widgets.QTableWidget
	StringFieldsButtonsLayout    *widgets.QVBoxLayout
	FileAttachmentsGroupBox      *widgets.QGroupBox
	FileAttachmentsTable         *widgets.QTableWidget
	FileAttachmentsButtonsLayout *widgets.QVBoxLayout
}

func NewAdvancedTabSheet() *AdvancedTabSheet {
	cw := &AdvancedTabSheet{
		QWidget: widgets.NewQWidget(nil, 0),
	}

	// Create the layout for the central widget
	layout := widgets.NewQVBoxLayout2(cw)

	// Create the upper group box for "String fields"
	cw.StringFieldsGroupBox = widgets.NewQGroupBox2("String fields", nil)
	layout.AddWidget(cw.StringFieldsGroupBox, 0, core.Qt__AlignTop)
	cw.StringFieldsGroupBox.Resize2(700, 600)
	// Create the layout for the upper group box
	stringFieldsLayout := widgets.NewQHBoxLayout2(nil)
	cw.StringFieldsGroupBox.SetLayout(stringFieldsLayout)

	// Create the table for "String fields"
	cw.StringFieldsTable = widgets.NewQTableWidget(nil)
	cw.StringFieldsTable.SetColumnCount(2)
	cw.StringFieldsTable.SetRowCount(0)
	cw.StringFieldsTable.SetHorizontalHeaderLabels([]string{"Field Name", "Field Value"})
	cw.StringFieldsTable.SetFixedWidth(500)
	stringFieldsLayout.AddWidget(cw.StringFieldsTable, 0, core.Qt__AlignLeft)

	// Create the buttons for "String fields"
	cw.StringFieldsButtonsLayout = widgets.NewQVBoxLayout()
	addButton := widgets.NewQPushButton2("Add", nil)
	editButton := widgets.NewQPushButton2("Edit", nil)
	deleteButton := widgets.NewQPushButton2("Delete", nil)
	moveButton := widgets.NewQPushButton2("Move", nil)
	cw.StringFieldsButtonsLayout.AddWidget(addButton, 0, core.Qt__AlignTop)
	cw.StringFieldsButtonsLayout.AddWidget(editButton, 0, core.Qt__AlignTop)
	cw.StringFieldsButtonsLayout.AddWidget(deleteButton, 0, core.Qt__AlignTop)
	cw.StringFieldsButtonsLayout.AddWidget(moveButton, 0, core.Qt__AlignTop)
	stringFieldsLayout.AddLayout(cw.StringFieldsButtonsLayout, 0)

	// Create the lower group box for "File attachments"
	cw.FileAttachmentsGroupBox = widgets.NewQGroupBox2("File attachments", nil)
	layout.AddWidget(cw.FileAttachmentsGroupBox, 0, core.Qt__AlignTop)

	// Create the layout for the lower group box
	fileAttachmentsLayout := widgets.NewQHBoxLayout2(nil)
	cw.FileAttachmentsGroupBox.SetLayout(fileAttachmentsLayout)

	// Create the table for "File attachments"
	cw.FileAttachmentsTable = widgets.NewQTableWidget(nil)
	cw.FileAttachmentsTable.SetColumnCount(2)
	cw.FileAttachmentsTable.SetRowCount(0)
	cw.FileAttachmentsTable.SetHorizontalHeaderLabels([]string{"Attachments", "Size"})
	cw.FileAttachmentsTable.SetFixedWidth(500)
	fileAttachmentsLayout.AddWidget(cw.FileAttachmentsTable, 0, core.Qt__AlignLeft)

	// Create the buttons for "File attachments"
	cw.FileAttachmentsButtonsLayout = widgets.NewQVBoxLayout()
	attachButton := widgets.NewQPushButton2("Attach", nil)
	deleteAttachmentButton := widgets.NewQPushButton2("Delete", nil)
	openAttachmentButton := widgets.NewQPushButton2("Open", nil)
	saveAttachmentButton := widgets.NewQPushButton2("Save", nil)
	cw.FileAttachmentsButtonsLayout.AddWidget(attachButton, 0, core.Qt__AlignTop)
	cw.FileAttachmentsButtonsLayout.AddWidget(deleteAttachmentButton, 0, core.Qt__AlignTop)
	cw.FileAttachmentsButtonsLayout.AddWidget(openAttachmentButton, 0, core.Qt__AlignTop)
	cw.FileAttachmentsButtonsLayout.AddWidget(saveAttachmentButton, 0, core.Qt__AlignTop)
	fileAttachmentsLayout.AddLayout(cw.FileAttachmentsButtonsLayout, 0)

	return cw
}

func (advancedTabSheet *AdvancedTabSheet) SetTableRowData(kpEntry gokeepasslib.Entry) {
	if kpEntry.Values != nil {
		j := 0
		for _, value := range kpEntry.Values {
			log.Info("key: " + value.Key)
			if value.Key != "Notes" && value.Key != "URL" && value.Key != "Password" && value.Key != "Title" && value.Key != "UserName" {
				advancedTabSheet.StringFieldsTable.SetRowCount(j + 1)
				advancedTabSheet.StringFieldsTable.SetItem(j, 0, widgets.NewQTableWidgetItem2(value.Key, 0))
				advancedTabSheet.StringFieldsTable.SetItem(j, 1, widgets.NewQTableWidgetItem2(value.Value.Content, 0))
				j++
			}
		}
	}
	/*for column := 0; column < tableWidget.ColumnCount(); column++ {
		tableWidget.SetItem(newRow, column, widgets.NewQTableWidgetItem2(rowData[column], 0))
	}
	tableWidget.SelectRow(newRow)*/
}
