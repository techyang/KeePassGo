package kpwidgets

import "github.com/therecipe/qt/widgets"

type KeePassTabWidget struct {
	TabWidget     *widgets.QTabWidget
	EntryTab      *EntryTab
	AdvancedTab   *AdvancedTab
	PropertiesTab *PropertiesTab
	AutoTypeTab   *AutoTypeTab
	HistoryTab    *widgets.QWidget
}

func NewKeePassTabWidget(parent widgets.QWidget_ITF) *KeePassTabWidget {
	tabWidget := widgets.NewQTabWidget(parent)

	entryTab := NewEntryTab()
	advancedTab := NewAdvanceTab()

	propertiesTab := NewPropertiesTab()
	autoTypeTab := NewAutoTypeTab()
	historyTab := widgets.NewQWidget(nil, 0)

	tabWidget.AddTab(entryTab.Widget, "Entry")
	tabWidget.AddTab(advancedTab.Widget, "Advanced")
	tabWidget.AddTab(propertiesTab.Widget, "Properties")
	tabWidget.AddTab(autoTypeTab.Widget, "Auto-Type")
	tabWidget.AddTab(historyTab, "History")

	return &KeePassTabWidget{
		TabWidget:     tabWidget,
		EntryTab:      entryTab,
		AdvancedTab:   advancedTab,
		PropertiesTab: propertiesTab,
		AutoTypeTab:   autoTypeTab,
		HistoryTab:    historyTab,
	}
}

func (tw *KeePassTabWidget) Resize(width, height int) {
	tw.TabWidget.Resize2(width, height)
}
