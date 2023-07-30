package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
)

type EntryTabSheet struct {
	*widgets.QWidget
	Title              *widgets.QLineEdit
	LastNameLineEdit   *widgets.QLineEdit
	UserNameEdit       *widgets.QLineEdit
	PasswordEdit       *widgets.QLineEdit
	RepeatPasswordEdit *widgets.QLineEdit
	ProgressBar        *widgets.QProgressBar
	URLEdit            *widgets.QLineEdit
	NotesEdit          *widgets.QTextEdit
	ExpiresCheckBox    *widgets.QCheckBox
	DateTimeEdit       *widgets.QDateTimeEdit
}

func NewEntryTabSheet() *EntryTabSheet {
	entry := &EntryTabSheet{
		QWidget:            widgets.NewQWidget(nil, 0),
		Title:              widgets.NewQLineEdit(nil),
		LastNameLineEdit:   widgets.NewQLineEdit(nil),
		UserNameEdit:       widgets.NewQLineEdit(nil),
		PasswordEdit:       widgets.NewQLineEdit(nil),
		RepeatPasswordEdit: widgets.NewQLineEdit(nil),
		ProgressBar:        widgets.NewQProgressBar(nil),
		URLEdit:            widgets.NewQLineEdit(nil),
		NotesEdit:          widgets.NewQTextEdit(nil),
		ExpiresCheckBox:    widgets.NewQCheckBox2("Expires:", nil),
		DateTimeEdit:       widgets.NewQDateTimeEdit(nil),
	}
	iconLabel := widgets.NewQLabel2("Icon:", nil, 0)
	iconButton := widgets.NewQPushButton(nil)
	iconButton.SetIcon(gui.NewQIcon5("Ext/Images_Client_HighRes/C00_Password.png")) // Replace with the path to your icon file

	nameLayout := widgets.NewQHBoxLayout2(nil)
	nameLayout.AddWidget(entry.Title, 0, core.Qt__AlignLeft)
	nameLayout.AddWidget(iconLabel, 0, core.Qt__AlignLeft)
	nameLayout.AddWidget(iconButton, 0, core.Qt__AlignLeft)
	entry.Title.SetFixedWidth(500)
	label2 := widgets.NewQLabel2("Title:", nil, 0)
	formLayout := widgets.NewQFormLayout(nil)
	formLayout.AddRow2(label2, nameLayout)

	formLayout.AddRow3("User name:", entry.UserNameEdit)
	formLayout.AddRow3("Password:", entry.PasswordEdit)
	formLayout.AddRow3("Repeat:", entry.RepeatPasswordEdit)

	entry.ProgressBar.SetRange(0, 100)
	formLayout.AddRow3("Quality:", entry.ProgressBar)

	formLayout.AddRow3("URL:", entry.URLEdit)

	entry.NotesEdit.Resize2(300, 200)
	formLayout.AddRow3("Notes:", entry.NotesEdit)

	//entry.FormLayout.AddRow3("Expires:", entry.DateTimeEdit)
	entry.DateTimeEdit.SetCalendarPopup(true)
	entry.DateTimeEdit.SetFixedWidth(400)
	entry.DateTimeEdit.SetDisplayFormat("yyyy-MM-dd HH:mm:ss")

	button := widgets.NewQPushButton2("Get DateTime", nil)
	button.ConnectClicked(func(checked bool) {
		selectedDateTime := entry.DateTimeEdit.DateTime().ToString("2006-01-02 15:04:05")
		entry.DateTimeEdit.SetDateTime(core.QDateTime_CurrentDateTime())
		fmt.Println("Selected DateTime:", selectedDateTime)
	})

	//entry.FormLayout.AddRow3("Change datetime:", button)
	// Create a checkbox to replace the label

	iconButton2 := widgets.NewQPushButton(nil)
	iconButton2.SetIcon(gui.NewQIcon5("Ext/Images_Client_HighRes/C00_Password.png")) // Replace with the path to your icon file

	vLayout := widgets.NewQVBoxLayout2(entry)
	nameLayout2 := widgets.NewQHBoxLayout2(nil)
	nameLayout2.AddWidget(entry.ExpiresCheckBox, 0, core.Qt__AlignLeft)
	nameLayout2.AddWidget(entry.DateTimeEdit, 0, core.Qt__AlignLeft)
	nameLayout2.AddWidget(iconButton2, 0, core.Qt__AlignLeft)
	nameLayout2.AddSpacing(100)
	vLayout.AddLayout(formLayout, 0)
	vLayout.AddLayout(nameLayout2, 0)

	// Create the context menu
	contextMenu := widgets.NewQMenu(nil)

	// Add actions to the context menu
	action1 := contextMenu.AddAction("Option 1")
	action2 := contextMenu.AddAction("Option 2")
	action1.Text()
	action2.Text()
	// Connect the button's clicked signal to a slot that shows the context menu
	iconButton2.ConnectCustomContextMenuRequested(func(pos *core.QPoint) {
		contextMenu.Exec2(iconButton2.MapFromGlobal(pos), nil)
		//contextMenu.Popup(pos, nil)
	})

	/*iconButton2.ConnectClicked(func(checked bool) {
		fmt.Println("OK button clicked")

		contextMenu.Popup()
	})*/

	// Add the checkbox to the form layout
	return entry
}

func (entryTab *EntryTabSheet) InitEntryTab(tableWidget *KeePassTable) {
	keePassEntry := GetKeePassEntry(tableWidget.ObjectName(), tableWidget.CurrentRow())
	entryTab.Title.SetText(keePassEntry.Title)
	entryTab.UserNameEdit.SetText(keePassEntry.UserName)
	entryTab.URLEdit.SetText(keePassEntry.URL)
	entryTab.NotesEdit.SetText(keePassEntry.Notes)
	entryTab.PasswordEdit.SetText(keePassEntry.Password)
	entryTab.RepeatPasswordEdit.SetText(keePassEntry.Password)

	dateTime := keePassEntry.Expires

	qDate := core.NewQDate3(dateTime.Year(), int(dateTime.Month()), dateTime.Day())
	qTime := core.NewQTime3(dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0)
	qDateTime := core.NewQDateTime4(qDate, qTime, core.Qt__LocalTime, 0)
	//dateTimeEdit.SetDateTime(qDateTime)
	entryTab.DateTimeEdit.SetDateTime(qDateTime)
	entryTab.ExpiresCheckBox.SetChecked(keePassEntry.ExpiresSeted)

}

func ReAddTableItem(entry *EntryTabSheet, tableWidget *KeePassTable) {
	tableWidget.SetRowCount(tableWidget.RowCount() + 1)
	// Create and set QTableWidgetItem for each cell
	tableWidget.SetItem(tableWidget.RowCount()-1, 0, widgets.NewQTableWidgetItem2(entry.Title.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 1, widgets.NewQTableWidgetItem2(entry.UserNameEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 2, widgets.NewQTableWidgetItem2(entry.PasswordEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 3, widgets.NewQTableWidgetItem2(entry.URLEdit.Text(), 0))
	tableWidget.SetItem(tableWidget.RowCount()-1, 4, widgets.NewQTableWidgetItem2(entry.NotesEdit.ToPlainText(), 0))
}

func GetKeePassEntry(groupUUID string, itemIndex int) *KeePassEntry {
	//treeWidget.SetHeaderLabels([]string{"yangwl"})
	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()
	rootGroups := db.Content.Root.Groups

	group := findGroupByUUID(rootGroups, groupUUID)

	entry := group.Entries[itemIndex]

	fmt.Println("entry.GetPassword():", entry.GetPassword())
	fmt.Println("entry.GetTitle():", entry.GetTitle())
	username := entry.Get("UserName").Value.Content
	url := entry.Get("URL").Value.Content
	notes := entry.Get("Notes").Value.Content

	expiryTime := entry.Times.ExpiryTime.Time
	expiresSeted := entry.Times.Expires.Bool
	// If "ExpiryTime" attribute is not found, try "Expires" attribute
	//expiryTime, ok = entry.Get("Expires")
	fmt.Println("Expiry date not found in entry:", expiryTime)

	return &KeePassEntry{
		Title:        entry.GetTitle(),
		UserName:     username,
		Password:     entry.GetPassword(),
		URL:          url,
		Notes:        notes,
		Expires:      expiryTime,
		ExpiresSeted: expiresSeted,
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
