package kpwidgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type PasswordDelegate struct {
	widgets.QStyledItemDelegate
}

func NewPasswordDelegate() *PasswordDelegate {
	return &PasswordDelegate{}
}

func (delegate *PasswordDelegate) DisplayText(value *core.QVariant, locale *core.QLocale) string {
	// If the value is a string, return "***"
	if value.Type() == core.QVariant__String {
		return "***"
	}

	// Otherwise, call the base class method to display the default text
	return delegate.QStyledItemDelegate.DisplayText(value, locale)
}

func InitMainContent(window *widgets.QMainWindow) {
	// Create a QVBoxLayout and a QWidget as the container

	splitter := widgets.NewQSplitter2(core.Qt__Horizontal, nil)

	// Create the left and right widgets
	//leftWidget := widgets.NewQWidget(nil, 0)
	hBoxLayout := widgets.NewQVBoxLayout()
	//	hBoxLayout2 := widgets.NewQVBoxLayout()

	container := widgets.NewQWidget(nil, 0)
	layout := widgets.NewQHBoxLayout2(container)

	// Create the QHBoxLayout for the buttons
	//	hBoxLayout := widgets.NewQVBoxLayout()

	// Create the buttons
	//button1 := widgets.NewQPushButton2("Button 1", nil)
	//button2 := widgets.NewQPushButton2("Button 2", nil)

	// Create a QTreeWidget
	tableWidget := NewKeePassTable()
	treeWidget := NewKeePassTree(tableWidget)

	//restTable(tableWidget)

	// Add the buttons to the QHBoxLayout
	//hBoxLayout.AddWidget(button1, 0, 0)
	hBoxLayout.AddWidget(treeWidget, 0, 0)

	button3 := widgets.NewQPushButton2("Button 3", nil)
	// Add the QHBoxLayout to the QVBoxLayout
	layout.AddLayout(hBoxLayout, 0)
	layout.AddWidget(button3, 0, 0)

	splitter.AddWidget(treeWidget)
	splitter.AddWidget(tableWidget)
	// Set the container as the central widget of the main window
	window.SetCentralWidget(splitter)

}

func InitToolbar(window *widgets.QMainWindow) {
	toolBar := widgets.NewQToolBar("Toolbar", window)

	// Add tool buttons to the toolbar
	newToolButton := widgets.NewQToolButton(nil)
	newToolButton.SetText("New")
	//newToolIcon := window.Style().StandardIcon(widgets.QStyle__SP_FileIcon, nil, nil)
	//newToolButton.SetIcon(newToolIcon)
	//.
	newToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_FileNew.png"))
	newToolButton.SetToolTip("New")
	iconSize := newToolButton.IconSize()

	buttonWidth := iconSize.Width() + 10
	buttonHeight := iconSize.Height() + 8
	newToolButton.SetFixedSize2(buttonWidth, buttonHeight)

	newToolButton.AdjustSize()

	// Add tool buttons to the toolbar
	openToolButton := widgets.NewQToolButton(nil)
	openToolButton.SetToolTip("Open")
	openToolButton.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Folder_Yellow_Open.png"))
	openToolButton.SetFixedSize2(buttonWidth, buttonHeight)

	openToolButton.AdjustSize()

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

	/*newAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-new"), "New", nil)
	openAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-open"), "Open", nil)
	saveAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-save"), "Save", nil)
	B16x16_FileSaveAs
	// Add the actions to the toolbar
	toolBar.AddActions([]*widgets.QAction{newAction, openAction, saveAction})*/

	toolBar.AddWidget(newToolButton)
	toolBar.AddWidget(openToolButton)
	toolBar.AddWidget(saveAsToolButton)
	toolBar.AddWidget(addEntityToolButton)
	toolBar.AddWidget(copyUserNameToolButton)
	toolBar.AddWidget(copyPasswordToolButton)

	// Add the toolbar to the main window
	window.AddToolBar2(toolBar)
}
