package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type EntryTab struct {
	Widget             *widgets.QWidget
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

func NewEntryTab() *EntryTab {
	entry := &EntryTab{
		Widget:             widgets.NewQWidget(nil, 0),
		Title:              widgets.NewQLineEdit(nil),
		LastNameLineEdit:   widgets.NewQLineEdit(nil),
		UserNameEdit:       widgets.NewQLineEdit(nil),
		PasswordEdit:       widgets.NewQLineEdit(nil),
		RepeatPasswordEdit: widgets.NewQLineEdit(nil),
		ProgressBar:        widgets.NewQProgressBar(nil),
		URLEdit:            widgets.NewQLineEdit(nil),
		NotesEdit:          widgets.NewQTextEdit(nil),
		DateTimeEdit:       widgets.NewQDateTimeEdit(nil),
	}

	nameLayout := widgets.NewQHBoxLayout2(nil)
	nameLayout.AddWidget(entry.Title, 0, core.Qt__AlignLeft)
	nameLayout.AddWidget(entry.LastNameLineEdit, 0, core.Qt__AlignLeft)
	label2 := widgets.NewQLabel2("Title:", nil, 0)
	entry.FormLayout = widgets.NewQFormLayout(entry.Widget)
	entry.FormLayout.AddRow2(label2, nameLayout)

	entry.FormLayout.AddRow3("User name:", entry.UserNameEdit)
	entry.FormLayout.AddRow3("Password:", entry.PasswordEdit)
	entry.FormLayout.AddRow3("Repeat:", entry.RepeatPasswordEdit)

	entry.ProgressBar.SetRange(0, 100)
	entry.FormLayout.AddRow3("Quality:", entry.ProgressBar)

	entry.FormLayout.AddRow3("URL:", entry.URLEdit)

	entry.NotesEdit.Resize2(300, 200)
	entry.FormLayout.AddRow3("Notes:", entry.NotesEdit)

	entry.FormLayout.AddRow3("Expires:", entry.DateTimeEdit)

	button := widgets.NewQPushButton2("Get DateTime", nil)
	button.ConnectClicked(func(checked bool) {
		selectedDateTime := entry.DateTimeEdit.DateTime().ToString("2006-01-02 15:04:05")
		entry.DateTimeEdit.SetDateTime(core.QDateTime_CurrentDateTime())
		fmt.Println("Selected DateTime:", selectedDateTime)
	})

	entry.FormLayout.AddRow3("Change datetime:", button)

	return entry
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
