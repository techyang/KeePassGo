package entity

import (
	"fmt"
	"github.com/techyang/keepassgo/constants"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os"
	"os/exec"
	"runtime"
)

func DuplicationEntity(tableWidget *KeePassTable) {
	dialog := NewDuplicationOptionsDialog()

	dialog.ButtonBox.ConnectAccepted(func() {
		dialog.DoAccepted(tableWidget)
	})

	// Connect the button box's rejected signal
	dialog.ButtonBox.ConnectRejected(func() {
		fmt.Println("Cancel button clicked")
		dialog.Reject()
	})
	dialog.Exec()
}

func doLinkClicked(link string) {
	//link = "D:\\Program Files\\TotalCMDPortable\\App\\totalcmd\\Plugins\\wcx\\Total7zip\\7-zip.chm"
	/*if err := exec.Command("hh.exe", link).Start(); err != nil {
		fmt.Println("Error opening .chm file:", err)
	}*/
	//chmFilePath := "path/to/helpfile.chm"

	// Replace "your-topic" with the topic you want to display in the CHM file.
	topic := "Working In Spy++"

	// Use different commands based on the operating system.
	var command string
	var args []string
	switch runtime.GOOS {
	case "windows":
		command = "hh.exe"
		args = []string{link, fmt.Sprintf("-#%s", topic)}
	default:
		fmt.Println("Unsupported operating system.")
		os.Exit(1)
	}

	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error opening CHM file:", err)
		os.Exit(1)
	}
}

/*func SetEditOrViewEntryAction(tableWidget *KeePassTable, editOrViewEntryAction *widgets.QAction) {
	editOrViewEntryAction.ConnectTriggered(func(bool) {
		InitDetailWidget(tableWidget)
	})
}*/

func SetCopyUrlAction(tableWidget *KeePassTable, copyUrlAction *widgets.QAction) {
	copyUrlAction.ConnectTriggered(func(bool) {
		tableWidget.CopyUrl()
	})
}

func SetOpenUrlAction(tableWidget *KeePassTable, openUrlAction *widgets.QAction) {
	openUrlAction.ConnectTriggered(func(bool) {
		tableWidget.OpenWithBrowser(constants.Browser_Default)
		//OpenTableItemUrl(tableWidget, constants.Browser_Default)
	})
}

/*func OpenTableItemUrl(tableWidget *KeePassTable, browser constants.Browser) {
	OpenTableItemUrlWithOption(tableWidget, browser, "")
}

func OpenTableItemUrlWithOption(tableWidget *KeePassTable, browser constants.Browser, option string) {
	selectedRow := tableWidget.CurrentRow()

	// Retrieve the item at the first column of the selected row
	item := tableWidget.Item(selectedRow, 3)

	// Get the text of the item
	if item != nil {
		url := item.Text()
		fmt.Println("Text of the first item in the selected row:", url)

		if len(option) > 0 {
			if constants.Browser_InternetExplorer == browser {
				//gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput(url))
				cmd := exec.Command("cmd", "/c", "start", "iexplore", "-private", url)
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			} else {
				open.RunWith(url, string(browser))
			}
		} else {
			if "default" == browser {
				//gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput(url))
				open.Run(url)
			} else {
				open.RunWith(url, string(browser))
			}
		}
	}
}*/

func SetCopyPasswordAction(tableWidget *KeePassTable, copyPasswordAction *widgets.QAction) {
	copyPasswordAction.ConnectTriggered(func(bool) {
		CopyTableItemPassword(tableWidget)
	})
}

func CopyTableItemPassword(tableWidget *KeePassTable) {
	selectedRow := tableWidget.CurrentRow()

	// Retrieve the item at the first column of the selected row
	item := tableWidget.Item(selectedRow, 2)

	// Get the text of the item
	if item != nil {
		firstItemText := item.Text()
		fmt.Println("Text of the first item in the selected row:", firstItemText)
		clipboard := gui.QGuiApplication_Clipboard()
		if clipboard != nil {
			clipboard.SetText(firstItemText, gui.QClipboard__Clipboard)
		}
	}
}

func CopyTableItemUsername(tableWidget *KeePassTable) {
	selectedRow := tableWidget.CurrentRow()

	// Retrieve the item at the first column of the selected row
	item := tableWidget.Item(selectedRow, 1)

	// Get the text of the item
	if item != nil {
		firstItemText := item.Text()
		fmt.Println("Text of the first item in the selected row:", firstItemText)
		clipboard := gui.QGuiApplication_Clipboard()
		if clipboard != nil {
			clipboard.SetText(firstItemText, gui.QClipboard__Clipboard)
		}
	}
}

func SetTableRowData(tableWidget *KeePassTable, newRow int, rowData []string) {
	for column := 0; column < tableWidget.ColumnCount(); column++ {
		tableWidget.SetItem(newRow, column, widgets.NewQTableWidgetItem2(rowData[column], 0))
	}
	tableWidget.SelectRow(newRow)
}

// Function to retrieve the data of a specific row in a QTableWidget and store it in an array
func GetRowData(tableWidget *KeePassTable, row int) []string {
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

func InitDetailWidget(tableWidget *KeePassTable) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	keePassTabWidget := NewKeePassTabWidget(dialog)
	keePassTabWidget.Resize(600, 400)

	keePassEntry, entry := GetKeePassEntry(tableWidget.ObjectName(), tableWidget.CurrentRow())

	keePassTabWidget.EntryTab.InitEntryTab2(keePassEntry)
	//keePassTabWidget.HistoryTab.SetTableRowData2(entry)
	keePassTabWidget.AdvancedTab.SetTableRowData(entry)
	hBoxLayout := initBottomButton(keePassTabWidget, tableWidget, dialog)

	vBoxLayout := widgets.NewQVBoxLayout2(dialog)
	vBoxLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(keePassTabWidget.TabWidget, 0, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(hBoxLayout, 0)

	dialog.Resize2(600, 400)
	dialog.Exec()

	return dialog
}

func NewDetailWidget(tableWidget *KeePassTable) *widgets.QDialog {
	// Create and add tabs to the tab widget
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Open Dialog")

	imageLabel := initKeePassImage()

	// Create the tab widget
	keePassTabWidget := NewKeePassTabWidget(dialog)
	keePassTabWidget.Resize(600, 400)

	//keePassEntry, entry := GetKeePassEntry(tableWidget.ObjectName(), tableWidget.CurrentRow())

	//keePassTabWidget.EntryTab.InitEntryTab2(keePassEntry)
	//keePassTabWidget.HistoryTab.SetTableRowData2(entry)
	//keePassTabWidget.AdvancedTab.SetTableRowData(entry)
	hBoxLayout := initBottomButton(keePassTabWidget, tableWidget, dialog)

	vBoxLayout := widgets.NewQVBoxLayout2(dialog)
	vBoxLayout.AddWidget(imageLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(keePassTabWidget.TabWidget, 0, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(hBoxLayout, 0)

	dialog.Resize2(600, 400)
	dialog.Exec()

	return dialog
}

func initKeePassImage() *widgets.QLabel {
	imageLabel := widgets.NewQLabel(nil, 0)
	imagePixmap := gui.NewQPixmap3("src\\Hello\\img\\keepass.png", "", core.Qt__AutoColor)
	imageLabel.SetPixmap(imagePixmap)
	return imageLabel
}

func initBottomButton(keePassDialog *KeePassTabWidget, tableWidget *KeePassTable, dialog *widgets.QDialog) *widgets.QHBoxLayout {
	entryTab := keePassDialog.EntryTab
	//advancedTab := keePassDialog.AdvancedTab

	hBoxLayout := widgets.NewQHBoxLayout2(nil)
	toolButton := widgets.NewQPushButton2("Tool", nil)

	toolButton.SetIcon(gui.NewQIcon5("/Ext/Images_Client_16/C59_Package_Development.png")) // Replace with the path to your icon file
	toolButton.SetIconSize(core.NewQSize2(32, 32))                                         // Set the size of the icon
	//toolButton.SetText("Tool")

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

		doOkButtonClicked(keePassDialog, entryTab, tableWidget, dialog)
	})

	cancelButton.ConnectClicked(func(bool) {
		// Code to handle cancelButton click event
		fmt.Println("Cancel Button clicked")
		dialog.Close()
	})
	return hBoxLayout
}

func doOkButtonClicked(keePassDialog *KeePassTabWidget, entryTab *EntryTabSheet, tableWidget *KeePassTable, dialog *widgets.QDialog) {
	msgBox := widgets.NewQMessageBox(nil)
	msgBox.SetWindowTitle("提示信息")
	//msgBox.SetText(keePassDialog.EntryTabSheet.UserNameEdit.Text())
	msgBox.SetInformativeText(keePassDialog.EntryTab.UserNameEdit.Text())
	msgBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
	msgBox.SetIcon(widgets.QMessageBox__Information)
	msgBox.Exec()

	// Code to handle cancelButton click event
	//tabWidget.get
	fmt.Println("okButton clicked")
	ReAddTableItem(entryTab, tableWidget)

	file, _ := os.Open("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials("111111")
	_ = gokeepasslib.NewDecoder(file).Decode(db)

	db.UnlockProtectedEntries()

	fmt.Println("Password entry saved successfully.")

	dialog.Close()
}
