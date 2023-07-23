package kpwidgets

import "github.com/therecipe/qt/widgets"

type MyTabWidget struct {
	TabWidget     *widgets.QTabWidget
	EntryTab      *widgets.QWidget
	AdvancedTab   *widgets.QWidget
	PropertiesTab *widgets.QWidget
	AutoTypeTab   *widgets.QWidget
	HistoryTab    *widgets.QWidget
}

func NewMyTabWidget(parent widgets.QWidget_ITF) *MyTabWidget {
	tabWidget := widgets.NewQTabWidget(parent)

	entryTab := widgets.NewQWidget(nil, 0)
	advancedTab := widgets.NewQWidget(nil, 0)
	propertiesTab := widgets.NewQWidget(nil, 0)
	autoTypeTab := widgets.NewQWidget(nil, 0)
	historyTab := widgets.NewQWidget(nil, 0)

	tabWidget.AddTab(entryTab, "Entry")
	tabWidget.AddTab(advancedTab, "Advanced")
	tabWidget.AddTab(propertiesTab, "Properties")
	tabWidget.AddTab(autoTypeTab, "Auto-Type")
	tabWidget.AddTab(historyTab, "History")

	return &MyTabWidget{
		TabWidget:     tabWidget,
		EntryTab:      entryTab,
		AdvancedTab:   advancedTab,
		PropertiesTab: propertiesTab,
		AutoTypeTab:   autoTypeTab,
		HistoryTab:    historyTab,
	}
}

func (tw *MyTabWidget) Resize(width, height int) {
	tw.TabWidget.Resize2(width, height)
}
