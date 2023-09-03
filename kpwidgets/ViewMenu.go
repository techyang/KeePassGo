package kpwidgets

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func InitViewMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the file menu
	viewMenu := menuBar.AddMenu2("View")

	// Create actions for the file menu
	changeLanguageAction := viewMenu.AddAction("Change Language...")
	changeLanguageAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_Keyboard_Layout.png"))
	// Connect the actions and tool buttons to their respective triggered events
	changeLanguageAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	viewMenu.AddSeparator()

	// Create actions for the file menu
	showToolbarAction := viewMenu.AddAction("Show Toolbar")
	showToolbarAction.SetCheckable(true)
	showToolbarAction.SetChecked(true)
	// Connect the actions and tool buttons to their respective triggered events
	showToolbarAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	// Create actions for the file menu
	showEntryViewAction := viewMenu.AddAction("Show Entry View")
	showEntryViewAction.SetCheckable(true)
	showEntryViewAction.SetChecked(true)
	// Connect the actions and tool buttons to their respective triggered events
	showEntryViewAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	windowLayoutMenu := viewMenu.AddMenu2("Window Layout")
	// Create actions for the file menu
	statckedAction := windowLayoutMenu.AddAction("Stacked")
	// Connect the actions and tool buttons to their respective triggered events
	statckedAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	sideBySideAction := windowLayoutMenu.AddAction("Side by Side")
	// Connect the actions and tool buttons to their respective triggered events
	sideBySideAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	viewMenu.AddSeparator()
	alwaysOnTopAction := viewMenu.AddAction("Always on Top")
	alwaysOnTopAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})
	viewMenu.AddSeparator()

	configureColumnsAction := viewMenu.AddAction("Configure Columns...")
	configureColumnsAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B48x48_View_Detailed.png"))
	configureColumnsAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	sortByMenu := viewMenu.AddMenu2("Sort By")

	noSortAction := sortByMenu.AddAction("No Sort")
	noSortAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	sortByMenu.AddSeparator()
	sortByGroup := widgets.NewQActionGroup(sortByMenu)

	sortByTitleAction := sortByMenu.AddAction("Title")
	sortByTitleAction.SetActionGroup(sortByGroup)
	sortByTitleAction.SetCheckable(true)
	sortByTitleAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})
	sortByUserNameAction := sortByMenu.AddAction("User Name")
	sortByUserNameAction.SetActionGroup(sortByGroup)
	sortByUserNameAction.SetCheckable(true)
	sortByUserNameAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	sortByPasswordAction := sortByMenu.AddAction("Password")
	sortByPasswordAction.SetActionGroup(sortByGroup)
	sortByPasswordAction.SetCheckable(true)
	sortByPasswordAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	sortByUrlAction := sortByMenu.AddAction("Url")
	sortByUrlAction.SetActionGroup(sortByGroup)
	sortByUrlAction.SetCheckable(true)
	sortByUrlAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	sortByNotesAction := sortByMenu.AddAction("Notes")
	sortByNotesAction.SetActionGroup(sortByGroup)
	sortByNotesAction.SetCheckable(true)
	sortByNotesAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	sortByMenu.AddSeparator()
	sortByMenuGroup := widgets.NewQActionGroup(sortByMenu)
	sortByAscendingAction := sortByMenu.AddAction("Ascending")
	sortByAscendingAction.SetActionGroup(sortByMenuGroup)
	sortByAscendingAction.SetCheckable(true)
	sortByAscendingAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})
	sortByDescendingAction := sortByMenu.AddAction("Descending")
	sortByDescendingAction.SetActionGroup(sortByMenuGroup)
	sortByDescendingAction.SetCheckable(true)
	sortByDescendingAction.ConnectTriggered(func(checked bool) {
		DoNewAction(window)
	})

	tanViewOptionsMenu := viewMenu.AddMenu2("TAN View Options")
	useSimpleListViewAction := tanViewOptionsMenu.AddAction("Use Simple List View for TAN-Only Groups")
	useSimpleListViewAction.SetCheckable(true)
	useSimpleListViewAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	showTANIndicesAction := tanViewOptionsMenu.AddAction("Show TAN Indices in Entry Titles")
	showTANIndicesAction.SetCheckable(true)
	showTANIndicesAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	groupingInEntryListMenu := viewMenu.AddMenu2("Grouping in Entry List ")
	groupingInEntryListGroup := widgets.NewQActionGroup(sortByMenu)

	groupingInEntryListOnAction := groupingInEntryListMenu.AddAction("On")
	groupingInEntryListOnAction.SetCheckable(true)
	groupingInEntryListOnAction.SetActionGroup(groupingInEntryListGroup)
	groupingInEntryListOnAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	groupingInEntryListRecommendAction := groupingInEntryListMenu.AddAction("Auto(Recommend)")
	groupingInEntryListRecommendAction.SetCheckable(true)
	groupingInEntryListRecommendAction.SetActionGroup(groupingInEntryListGroup)
	groupingInEntryListRecommendAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	groupingInEntryListOffAction := groupingInEntryListMenu.AddAction("Off")
	groupingInEntryListOffAction.SetCheckable(true)
	groupingInEntryListOffAction.SetActionGroup(groupingInEntryListGroup)
	groupingInEntryListOffAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

	viewMenu.AddSeparator()
	showEntriesOfSubGroupsAction := viewMenu.AddAction("Show Entries of Subgroups")
	showEntriesOfSubGroupsAction.SetCheckable(true)
	showEntriesOfSubGroupsAction.ConnectTriggered(func(checked bool) {
		//DoNewAction(window)
	})

}
