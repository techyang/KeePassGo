package kpwidgets

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

// InitToolsMenu init tool menu content
func InitToolsMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	toolsMenu := menuBar.AddMenu2("Tools")

	// Create actions for the file menu
	generatePasswordAction := toolsMenu.AddAction("Generate Password...")
	generatePasswordAction.SetIcon(gui.NewQIcon5("Ext/Images_App_HighRes/Nuvola_Derived/B48x48_Key_New.png"))
	generatePasswordAction.ConnectTriggered(func(checked bool) {

	})

	generatePasswordListAction := toolsMenu.AddAction("Generate Password List...")
	generatePasswordListAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_KGPG_Gen.png"))
	generatePasswordListAction.ConnectTriggered(func(checked bool) {

	})

	toolsMenu.AddSeparator()

	tanWizardAction := toolsMenu.AddAction("TAN Wizard...")
	tanWizardAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_Wizard.png"))
	tanWizardAction.ConnectTriggered(func(checked bool) {

	})

	databaseToolsMenu := toolsMenu.AddMenu2("Database Tools")
	databaseMaintenanceAction := databaseToolsMenu.AddAction("Database Maintenance...")
	databaseToolsMenu.AddSeparator()
	databaseToolsMenu.AddAction("Delete duplicate Entries")
	databaseToolsMenu.AddAction("Delete Empty Groups")
	databaseToolsMenu.AddAction("Delete Unused Custom Icons")
	// Connect the actions and tool buttons to their respective triggered events
	/*databaseToolsAction.ConnectTriggered(func(checked bool) {
		functions.OpenDatabase(TreeWidget, TableWidget)
	})*/

	databaseMaintenanceAction.ConnectTriggered(func(checked bool) {

	})
	//toolsMenu.AddAction("")

	toolsMenu.AddSeparator()
	triggersAction := toolsMenu.AddAction("Triggers...")
	triggersAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_Make_KDevelop.png"))
	triggersAction.ConnectTriggered(func(checked bool) {

	})

	pluginsAction := toolsMenu.AddAction("Plugins...")
	pluginsAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_BlockDevice.png"))
	pluginsAction.ConnectTriggered(func(checked bool) {

	})

	toolsMenu.AddSeparator()
	optionsAction := toolsMenu.AddAction("Options...")
	optionsAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_KCMSystem.png"))
	optionsAction.ConnectTriggered(func(checked bool) {

	})

}
