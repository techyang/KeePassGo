package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	// Create the main window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Multi-Language Menu Demo")

	// Create the menu bar
	menuBar := window.MenuBar()

	// Create the file menu
	fileMenu := menuBar.AddMenu2(core.QCoreApplication_Translate("@file", "File", "", -1))

	// Add menu items to the file menu
	newAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-new"), core.QCoreApplication_Translate("@file", "New", "", -1), nil)
	openAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-open"), core.QCoreApplication_Translate("@file", "Open", "", -1), nil)
	saveAction := widgets.NewQAction3(gui.QIcon_FromTheme("document-save"), core.QCoreApplication_Translate("@file", "Save", "", -1), nil)
	//exitAction := widgets.NewQAction3(gui.QIcon_FromTheme("application-exit"), core.QCoreApplication_Translate("@file", "Exit", "", -1), nil)

	/*actions := make([]interface{}, len([]*widgets.QAction{newAction, openAction, saveAction}))
	for i, action := range []*widgets.QAction{newAction, openAction, saveAction} {
		actions[i] = action
	}*/
	actions := []interface{}{newAction, openAction, saveAction}
	fileMenu.AddActions(convertToQActions(actions))
	//fileMenu.AddActions(actions)
	fileMenu.AddSeparator()
	//fileMenu.AddAction(exitAction)

	// Connect the exit action to the application's Quit slot
	/*exitAction.ConnectTriggered(func(checked bool) {
		widgets.QApplication_Quit()
	})*/

	// Load the translation file for the desired language
	//translator := core.NewQTranslator(nil)
	//translator.Load("path/to/translation_file.qm", "")
	//widgets.QApplication_Instance().InstallTranslator(translator)

	// Show the main window
	window.Show()

	widgets.QApplication_Exec()
}

func convertToQActions(actions []interface{}) []*widgets.QAction {
	qActions := make([]*widgets.QAction, len(actions))
	for i, a := range actions {
		qActions[i] = a.(*widgets.QAction)
	}
	return qActions
}
