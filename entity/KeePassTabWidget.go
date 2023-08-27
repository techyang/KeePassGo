package entity

import (
	"github.com/therecipe/qt/widgets"
)

type KeePassTabWidget struct {
	TabWidget     *widgets.QTabWidget
	EntryTab      *EntryTabSheet
	AdvancedTab   *AdvancedTabSheet
	PropertiesTab *PropertiesTabSheet
	AutoTypeTab   *AutoTypeTabSheet
	HistoryTab    *HistoryTabSheet
}

func NewKeePassTabWidget(parent widgets.QWidget_ITF) *KeePassTabWidget {
	tabWidget := widgets.NewQTabWidget(parent)

	entryTab := NewEntryTabSheet()
	advancedTab := NewAdvancedTabSheet()

	propertiesTab := NewPropertiesTabSheet()
	autoTypeTab := NewAutoTypeTabSheet()
	historyTab := NewHistoryTabSheet()

	tabWidget.AddTab(entryTab, "Entry")
	tabWidget.AddTab(advancedTab, "Advanced")
	tabWidget.AddTab(propertiesTab, "Properties")
	tabWidget.AddTab(autoTypeTab, "Auto-Type")
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
