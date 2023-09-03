package kpwidgets

import (
	"github.com/therecipe/qt/widgets"
)

func InitToolsMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	toolsMenu := menuBar.AddMenu2("Tools")

	// Create actions for the file menu
	generatePasswordAction := toolsMenu.AddAction("Generate Password...")
	generatePasswordAction.ConnectTriggered(func(checked bool) {

	})

	generatePasswordListAction := toolsMenu.AddAction("Generate Password List...")
	generatePasswordListAction.ConnectTriggered(func(checked bool) {

	})

	toolsMenu.AddSeparator()

	tanWizardAction := toolsMenu.AddAction("TAN Wizard...")
	tanWizardAction.ConnectTriggered(func(checked bool) {

	})

	databaseToolsAction := toolsMenu.AddAction("Database Tools")
	databaseToolsAction.ConnectTriggered(func(checked bool) {

	})

	toolsMenu.AddSeparator()
	triggersAction := toolsMenu.AddAction("Triggers...")
	triggersAction.ConnectTriggered(func(checked bool) {

	})

	pluginsAction := toolsMenu.AddAction("Plugins...")
	pluginsAction.ConnectTriggered(func(checked bool) {

	})

	toolsMenu.AddSeparator()
	optionsAction := toolsMenu.AddAction("Options...")
	optionsAction.ConnectTriggered(func(checked bool) {

	})

}
