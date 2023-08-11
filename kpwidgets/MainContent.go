package kpwidgets

import (
	"github.com/therecipe/qt/core"
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
	toolBar := NewKeePassToolBar(window)
	// Add the toolbar to the main window
	window.AddToolBar2(toolBar)
}
