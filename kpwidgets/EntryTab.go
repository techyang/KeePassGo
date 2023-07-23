package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type EntryTab struct {
	FormLayout         *widgets.QFormLayout
	Title              *widgets.QLineEdit
	LastNameLineEdit   *widgets.QLineEdit
	UserNameEdit       *widgets.QLineEdit
	PasswordEdit       *widgets.QLineEdit
	RepeatPasswordEdit *widgets.QLineEdit
	ProgressBar        *widgets.QProgressBar
	URLEdit            *widgets.QLineEdit
	NotesEdit          *widgets.QTextEdit
	DateTimeEdit       *widgets.QDateTimeEdit
}

func (entry *EntryTab) InitEntryTab2(entryTab *widgets.QWidget) {
	// Create the entry tab struct
	//entryTabWidget := widgets.NewQWidget(nil, 0)

	// Create the form layout
	entry.FormLayout = widgets.NewQFormLayout(entryTab)

	// Create and add widgets to the form layout
	entry.Title = widgets.NewQLineEdit(nil)
	entry.LastNameLineEdit = widgets.NewQLineEdit(nil)

	nameLayout := widgets.NewQHBoxLayout2(nil)
	nameLayout.AddWidget(entry.Title, 0, core.Qt__AlignLeft)
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

func ReAddTableItem(entry *EntryTab, tableWidget *widgets.QTableWidget) {
	tableWidget.SetRowCount(tableWidget.RowCount() + 1)
	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(tableWidget.RowCount()-1, 0, widgets.NewQTableWidgetItem2(entry.Title.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 1, widgets.NewQTableWidgetItem2(entry.UserNameEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 2, widgets.NewQTableWidgetItem2(entry.PasswordEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 3, widgets.NewQTableWidgetItem2(entry.URLEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 4, widgets.NewQTableWidgetItem2(entry.NotesEdit.ToPlainText(), 0))

}
