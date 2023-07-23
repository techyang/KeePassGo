package kpwidgets

import "github.com/therecipe/qt/widgets"

type KeePassDialog struct {
	TabWidget     *widgets.QTabWidget
	EntryTab      *EntryTab
	AdvancedTab   *AdvancedTab
	PropertiesTab *widgets.QWidget
	AutoTypeTab   *widgets.QWidget
	HistoryTab    *widgets.QWidget
}

func NewKeePassDialog(parent widgets.QWidget_ITF) *KeePassDialog {
	tabWidget := widgets.NewQTabWidget(parent)

	entryTab := NewEntryTab()
	advancedTab := NewAdvanceTab()

	propertiesTab := widgets.NewQWidget(nil, 0)
	autoTypeTab := widgets.NewQWidget(nil, 0)
	historyTab := widgets.NewQWidget(nil, 0)

	tabWidget.AddTab(entryTab.Widget, "Entry")
	tabWidget.AddTab(advancedTab.Widget, "Advanced")
	tabWidget.AddTab(propertiesTab, "Properties")
	tabWidget.AddTab(autoTypeTab, "Auto-Type")
	tabWidget.AddTab(historyTab, "History")

	return &KeePassDialog{
		TabWidget:     tabWidget,
		EntryTab:      entryTab,
		AdvancedTab:   advancedTab,
		PropertiesTab: propertiesTab,
		AutoTypeTab:   autoTypeTab,
		HistoryTab:    historyTab,
	}
}

func (tw *KeePassDialog) Resize(width, height int) {
	tw.TabWidget.Resize2(width, height)
}
