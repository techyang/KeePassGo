package kpwidgets

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/techyang/keepassgo/constants"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/tobischo/gokeepasslib/v3"
	"os/exec"
	"runtime"
)

type KeePassTable struct {
	*CommonTableWidget
}

func NewKeePassTable() *KeePassTable {
	// Create a QTableWidget
	// Create the table widget
	headerLabels := []string{"Title", "User Name", "Password", "URL", "Notes"}
	tableWidget := NewCommonTableWidget(headerLabels)
	tableWidget.SetColumnCount(5)

	// Set the header labels

	//tableWidget.SetHorizontalHeaderLabels(headerLabels)

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
		CommonTableWidget: tableWidget,
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
	urlsMenu.AddSeparator()
	openWithIEAction := urlsMenu.AddAction("Open With Internet Explorer")
	openWithIEInPirvateAction := urlsMenu.AddAction("Open With Internet Explorer (Private)")
	urlsMenu.AddAction("Open With Edge").ConnectTriggered(func(bool) {
		keePassTable.openWithBrowser(constants.Browser_Edge)
	})
	urlsMenu.AddAction("Open With Google Chrome").ConnectTriggered(func(bool) {
		keePassTable.openWithBrowser(constants.Browser_Chrome)
	})
	urlsMenu.AddAction("Open With Google Chrome (Private)").ConnectTriggered(func(bool) {
		keePassTable.openWithBrowserInPrivate(constants.Browser_Chrome)
	})

	urlsMenu.AddAction("Open With 360安全浏览器").ConnectTriggered(func(bool) {
		open.Run("C:\\Program Files (x86)\\Sybase\\PowerDesigner 16\\bpm.chm")
		keePassTable.openWithBrowser(constants.Browser_360SE)
	})
	contextMenu.AddSeparator()
	openWithIEAction.ConnectTriggered(func(bool) {
		keePassTable.openWithBrowser(constants.Browser_InternetExplorer)
	})

	openWithIEInPirvateAction.ConnectTriggered(func(bool) {
		keePassTable.openWithBrowserInPrivate(constants.Browser_InternetExplorer)
	})

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
	setDuplicateAction2(keePassTable, duplicateAction)

	selectAllAction.ConnectTriggered(func(bool) {
		keePassTable.SelectAll()
	})

	keePassTable.ConnectCustomContextMenuRequested(func(pos *core.QPoint) {
		contextMenu.Exec2(keePassTable.MapToGlobal(pos), nil)
	})

	// Connect the itemSelectionChanged signal of the table widget
	keePassTable.ConnectItemSelectionChanged(func() {
		/*selectedItems := keePassTable.SelectedItems()
		if len(selectedItems) > 0 {
			//copyToolButton.SetEnabled(true) // Enable the button when items are selected
		} else {
			//w.copyToolButton.SetEnabled(false) // Disable the button when no items are selected
		}*/
	})

	// Connect the triggered signal of the menu actions
	addItemAction.ConnectTriggered(func(bool) {
		NewDetailWidget(keePassTable)
	})

	deleteItemAction.ConnectTriggered(func(bool) {

		row := keePassTable.CurrentRow()
		if row > 0 {
			keePassTable.RemoveRow(row)
		}
	})
}

func (keePassTable *KeePassTable) openWithBrowser(browser constants.Browser) {
	selectedRow := tableWidget.CurrentRow()
	item := tableWidget.Item(selectedRow, 3)

	// Get the text of the item
	if item != nil {
		url := item.Text()
		fmt.Println("Text of the first item in the selected row:", url)
		if browser == constants.Browser_Default {
			open.Run(url)
			return
		}
		open.RunWith(url, string(browser))
	}
}

func (keePassTable *KeePassTable) openWithBrowserInPrivate(browser constants.Browser) {
	selectedRow := tableWidget.CurrentRow()
	item := tableWidget.Item(selectedRow, 3)

	// Get the text of the item
	if item != nil {
		url := item.Text()
		openUrlInPrivate(browser, url)
	}
}

func openUrlInPrivate(browser constants.Browser, url string) {
	switch browser {
	case constants.Browser_Chrome:
		openChromeInPrivate(url)
	case constants.Browser_Firefox:
		openFirefoxInPrivate(url)
	case constants.Browser_InternetExplorer:
		openIEInPrivate(url)
	}
}

func openChromeInPrivate(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "chrome", "--incognito", url)
	case "darwin":
		cmd = exec.Command("open", "-a", "Google Chrome", "--args", "--incognito", url)
	case "linux":
		cmd = exec.Command("google-chrome", "--incognito", url)
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening browser:", err)
		// Fallback to using the open package
		err := open.Run(url)
		if err != nil {
			fmt.Println("Error opening URL:", err)
		}
	}
}

func openFirefoxInPrivate(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "firefox", "--private-window", url)
	case "darwin":
		cmd = exec.Command("open", "-a", "Firefox", "--args", "-private-window", url)
	case "linux":
		cmd = exec.Command("firefox", "--private-window", url)
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening browser:", err)
		// Fallback to using the open package
		err := open.Run(url)
		if err != nil {
			fmt.Println("Error opening URL:", err)
		}
	}
}

func openIEInPrivate(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "iexplore", "-private", url)
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Error opening browser:", err)
		// Fallback to using the open package
		err := open.Run(url)
		if err != nil {
			fmt.Println("Error opening URL:", err)
		}
	}
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
	iconPath := "Ext/Images_App_HighRes/Nuvola/B48x48_KGPG_Key1.png"
	for i, entry := range group.Entries {
		username := entry.Get("UserName").Value.Content
		url := entry.Get("URL").Value.Content
		note := entry.Get("Notes").Value.Content
		tableWidget.SetRowCount(i + 1)
		iconLabel := NewIconLabel(iconPath, "Item "+entry.GetTitle())
		tableWidget.SetCellWidget(i, 0, iconLabel)

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
