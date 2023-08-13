package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	log "log/slog"
)

type KeePassToolBar struct {
	*widgets.QToolBar
}

func NewKeePassToolBar(window *widgets.QMainWindow) *KeePassToolBar {
	toolBar := widgets.NewQToolBar("Toolbar", window)
	// Add tool buttons to the toolbar
	newToolButton := widgets.NewQToolButton(nil)
	newToolButton.SetText("New")

	newToolButton.ConnectClicked(func(bool) {
		// Code to handle cancelButton click event
		fmt.Println("toolButton clicked")
		treeWidget.QTreeWidget.Clear()
		//dialog.Close()
	})
	//newToolIcon := window.Style().StandardIcon(widgets.QStyle__SP_FileIcon, nil, nil)
	//newToolButton.SetIcon(newToolIcon)
	//.
	newToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_FileNew.png"))
	newToolButton.SetToolTip("New")
	iconSize := newToolButton.IconSize()

	buttonWidth := iconSize.Width() + 10
	buttonHeight := iconSize.Height() + 8
	buttonWidth = 22
	buttonHeight = 26
	newToolButton.SetFixedSize2(buttonWidth, buttonHeight)

	newToolButton.AdjustSize()

	// Add tool buttons to the toolbar
	openToolButton := widgets.NewQToolButton(nil)
	openToolButton.SetToolTip("Open")
	openToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Folder_Yellow_Open.png"))
	openToolButton.SetFixedSize2(buttonWidth, buttonHeight)

	openToolButton.AdjustSize()

	openToolButton.ConnectClicked(func(bool) {
		// Code to handle cancelButton click event
		fmt.Println("toolButton clicked")

		treeWidget.loadKeePassTree("D:\\workspace_go\\gokeepasslib-master\\example-new-database2023.kdbx")
		//dialog.Close()
	})

	saveAsToolButton := widgets.NewQToolButton(nil)
	saveAsToolButton.SetToolTip("Save")
	saveAsToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_FileSave.png"))
	saveAsToolButton.SetFixedSize2(buttonWidth, buttonHeight)
	saveAsToolButton.AdjustSize()

	addEntityToolButton := widgets.NewQToolButton(nil)
	addEntityToolButton.SetToolTip("Add Entity")
	addEntityToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Import.png"))
	addEntityToolButton.SetFixedSize2(buttonWidth, buttonHeight)
	addEntityToolButton.AdjustSize()

	copyUserNameToolButton := widgets.NewQToolButton(nil)
	copyUserNameToolButton.SetToolTip("Copy UserName to ClipBoard ")
	copyUserNameToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Personal.png"))
	copyUserNameToolButton.SetFixedSize2(buttonWidth, buttonHeight)
	copyUserNameToolButton.AdjustSize()

	copyPasswordToolButton := widgets.NewQToolButton(nil)
	copyPasswordToolButton.SetToolTip("Copy Password to ClipBoard ")
	//icon := gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Info.png")
	copyPasswordToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Info.png"))
	//copyPasswordToolButton.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	copyPasswordToolButton.SetFixedSize2(buttonWidth, buttonHeight)
	copyPasswordToolButton.AdjustSize()

	openUrlsToolButton := widgets.NewQToolButton(nil)
	openUrlsToolButton.SetToolTip("Open URL(s)")
	openUrlsToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Browser.png"))
	openUrlsToolButton.SetFixedSize2(22, 22)
	openUrlsToolButton.AdjustSize()

	// Create a menu for the first dropdown
	openUrlsMenu := widgets.NewQMenu(nil)
	action11 := openUrlsMenu.AddAction("Open URLs")
	action11.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Key3.png"))
	openUrlsMenu.AddSeparator()
	action22 := openUrlsMenu.AddAction("Open URLs with Internet Explorer")
	action22.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_History_Clear.png"))

	action33 := openUrlsMenu.AddAction("Open URLs with Internet Explorer (Private)")
	action33.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_History_Clear.png"))
	// Set the menus to the tool button
	openUrlsToolButton.SetMenu(openUrlsMenu)
	openUrlsToolButton.SetPopupMode(widgets.QToolButton__InstantPopup)

	copyUrlsToClipBoardToolButton := widgets.NewQToolButton(nil)
	copyUrlsToClipBoardToolButton.SetToolTip("Copy URL(s) to ClipBoard")
	copyUrlsToClipBoardToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola_Derived/B16x16_EditCopyUrl.png"))
	copyUrlsToClipBoardToolButton.SetFixedSize2(22, 22)
	copyUrlsToClipBoardToolButton.AdjustSize()

	performAutoTypeToolButton := widgets.NewQToolButton(nil)
	performAutoTypeToolButton.SetToolTip("Perform Auto-Type")
	performAutoTypeToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KTouch.png"))
	performAutoTypeToolButton.SetFixedSize2(22, 22)
	performAutoTypeToolButton.AdjustSize()

	findToolButton := widgets.NewQToolButton(nil)
	findToolButton.SetToolTip("Find")
	findToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_XMag.png"))
	findToolButton.SetFixedSize2(22, 22)
	findToolButton.AdjustSize()
	/*newAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-new"), "New", nil)
	openAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-open"), "Open", nil)
	saveAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-save"), "Save", nil)
	B16x16_FileSaveAs
	// Add the actions to the toolbar
	toolBar.AddActions([]*widgets.QAction{newAction, openAction, saveAction})*/

	showEntriesButton := widgets.NewQToolButton(nil)
	showEntriesButton.SetToolTip("Show Entries")
	showEntriesButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Key3.png"))
	//showEntriesButton.SetFixedSize2(22, 22)
	//showEntriesButton.AdjustSize()

	// Create a menu for the first dropdown
	menu1 := widgets.NewQMenu(nil)
	action1 := menu1.AddAction("All")
	action1.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KGPG_Key3.png"))
	action2 := menu1.AddAction("Expired")
	action1.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_History_Clear.png"))
	// Set the menus to the tool button
	showEntriesButton.SetMenu(menu1)
	showEntriesButton.SetPopupMode(widgets.QToolButton__InstantPopup)

	// Connect the actions to slots
	action1.ConnectTriggered(func(checked bool) {
		println("Option 1 selected")
	})

	action2.ConnectTriggered(func(checked bool) {
		println("Option 2 selected")
	})

	lockWorkspaceToolButton := widgets.NewQToolButton(nil)
	lockWorkspaceToolButton.SetToolTip("Lock Workspace")
	lockWorkspaceToolButton.SetIcon(gui.NewQIcon5("Resources/images/B16x16_LockWorkspace.png"))
	lockWorkspaceToolButton.SetFixedSize2(22, 22)
	lockWorkspaceToolButton.AdjustSize()

	// Create a QComboBox
	comboBox := widgets.NewQComboBox(nil)
	//comboBox.AddItems([]string{"  "})
	comboBox.SetEditable(true)
	comboBox.SetToolTip("Type to search the database")
	comboBoxLineEdit := comboBox.LineEdit()
	comboBoxLineEdit.SetPlaceholderText("Search...")
	// Connect a slot to the line edit's clicked signal
	// Connect a slot to the line edit's focus in event

	// Connect the line edit's FocusInEvent signal to clear the placeholder text
	comboBoxLineEdit.ConnectFocusInEvent(func(event *gui.QFocusEvent) {
		log.Info("ConnectFocusInEvent")
		comboBoxLineEdit.Clear()
	})
	/*comboBoxLineEdit.ConnectEventFilter(func(obj *core.QObject, event *core.QEvent) bool {
		log.Info("comboBoxLineEdit.ConnectEventFilter", event.Type())
		if event.Type() == core.QEvent__FocusIn {
			// Clear the placeholder text when the line edit gains focus
			comboBoxLineEdit.Clear()
		}
		return false
	})*/
	// Add items to the combo box
	//comboBox.AddItem("Option 1")
	//comboBox.AddItem("Option 2",0)
	//comboBox.AddItem("Option 3",0)

	// Set the minimum width of the combo box to fit its contents
	comboBox.SetMinimumWidth(100)

	toolBar.AddWidget(newToolButton)
	toolBar.AddWidget(openToolButton)
	toolBar.AddWidget(saveAsToolButton)
	toolBar.AddSeparator()
	toolBar.AddWidget(addEntityToolButton)
	toolBar.AddSeparator()
	toolBar.AddWidget(copyUserNameToolButton)
	toolBar.AddWidget(copyPasswordToolButton)
	toolBar.AddWidget(openUrlsToolButton)
	toolBar.AddWidget(copyUrlsToClipBoardToolButton)
	toolBar.AddWidget(performAutoTypeToolButton)
	toolBar.AddSeparator()
	toolBar.AddWidget(findToolButton)
	toolBar.AddWidget(showEntriesButton)
	toolBar.AddSeparator()
	toolBar.AddWidget(lockWorkspaceToolButton)
	toolBar.AddSeparator()
	toolBar.AddWidget(comboBox)
	entry := &KeePassToolBar{
		QToolBar: toolBar,
	}
	return entry
}
