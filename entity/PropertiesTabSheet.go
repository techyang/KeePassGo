package entity

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type PropertiesTabSheet struct {
	*widgets.QWidget
	ForegroundColorCheckbox  *widgets.QCheckBox
	ForegroundColorPickerBtn *widgets.QPushButton
	BackgroundColorCheckbox  *widgets.QCheckBox
	BackgroundColorPickerBtn *widgets.QPushButton
	TagsLabel                *widgets.QLabel
	TagsInput                *widgets.QLineEdit
	OverrideURLLabel         *widgets.QLabel
	OverrideURLInput         *widgets.QLineEdit
	PluginDataLabel          *widgets.QLabel
	PluginDataInput          *widgets.QTextEdit
	DeleteButton             *widgets.QPushButton
	UUIDLabel                *widgets.QLabel
	UUIDInput                *widgets.QLineEdit
}

func NewPropertiesTabSheet() *PropertiesTabSheet {
	cw := &PropertiesTabSheet{
		QWidget:                  widgets.NewQWidget(nil, 0),
		ForegroundColorCheckbox:  widgets.NewQCheckBox2("Custom foreground color", nil),
		ForegroundColorPickerBtn: widgets.NewQPushButton2("Pick Color", nil),
		BackgroundColorCheckbox:  widgets.NewQCheckBox2("Custom background color", nil),
		BackgroundColorPickerBtn: widgets.NewQPushButton2("Pick Color", nil),
		TagsLabel:                widgets.NewQLabel2("Tags", nil, 0),
		TagsInput:                widgets.NewQLineEdit(nil),
		OverrideURLLabel:         widgets.NewQLabel2("Override URL(e.g. to use a specific browser)", nil, 0),
		OverrideURLInput:         widgets.NewQLineEdit(nil),
		PluginDataLabel:          widgets.NewQLabel2("Plugin data:", nil, 0),
		PluginDataInput:          widgets.NewQTextEdit(nil),
		DeleteButton:             widgets.NewQPushButton2("Delete", nil),
		UUIDLabel:                widgets.NewQLabel2("UUID", nil, 0),
		UUIDInput:                widgets.NewQLineEdit(nil),
	}

	// Layout
	vBoxLayout := widgets.NewQVBoxLayout2(cw)
	nameLayout := widgets.NewQHBoxLayout2(nil)
	nameLayout.AddWidget(cw.ForegroundColorCheckbox, 0, core.Qt__AlignLeft)
	nameLayout.AddWidget(cw.ForegroundColorPickerBtn, 1, core.Qt__AlignLeft)
	nameLayout.AddSpacing(40)
	vBoxLayout.AddLayout(nameLayout, 0)

	nameLayout2 := widgets.NewQHBoxLayout2(nil)
	nameLayout2.AddWidget(cw.BackgroundColorCheckbox, 0, core.Qt__AlignLeft)
	nameLayout2.AddWidget(cw.BackgroundColorPickerBtn, 1, core.Qt__AlignLeft)
	nameLayout2.AddSpacing(40)
	vBoxLayout.AddLayout(nameLayout2, 0)

	vBoxLayout.AddWidget(cw.TagsLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.TagsInput, 0, core.Qt__AlignLeft)

	vBoxLayout.AddWidget(cw.OverrideURLLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.OverrideURLInput, 0, core.Qt__AlignLeft)

	vBoxLayout.AddWidget(cw.PluginDataLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.PluginDataInput, 0, core.Qt__AlignLeft)

	pluginDataLayout := widgets.NewQHBoxLayout2(nil)
	pluginDataLayout.AddWidget(cw.PluginDataInput, 0, core.Qt__AlignLeft)
	pluginDataLayout.AddWidget(cw.DeleteButton, 0, core.Qt__AlignTop)
	vBoxLayout.AddLayout(pluginDataLayout, 0)

	nameLayout3 := widgets.NewQHBoxLayout2(nil)
	nameLayout3.AddWidget(cw.UUIDLabel, 0, core.Qt__AlignLeft)
	nameLayout3.AddWidget(cw.UUIDInput, 1, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(nameLayout3, 0)

	cw.PluginDataInput.SetFixedWidth(500)
	cw.OverrideURLInput.SetFixedWidth(620)
	cw.TagsInput.SetFixedWidth(620)
	cw.UUIDInput.SetFixedWidth(600)
	// Connect the "Pick Color" buttons clicked signal to open the color selection dialogs
	cw.ForegroundColorPickerBtn.ConnectClicked(func(checked bool) {
		colorDialog := widgets.NewQColorDialog2(gui.NewQColor3(255, 255, 255, 255), nil)
		colorDialog.Exec()
		if colorDialog != nil {
			cw.ForegroundColorPickerBtn.SetStyleSheet("background-color: " + colorDialog.CurrentColor().Name())
		}
	})

	cw.BackgroundColorPickerBtn.ConnectClicked(func(checked bool) {
		colorDialog := widgets.NewQColorDialog2(gui.NewQColor3(255, 255, 255, 255), nil)
		colorDialog.Exec()
		if colorDialog != nil {
			cw.BackgroundColorPickerBtn.SetStyleSheet("background-color: " + colorDialog.CurrentColor().Name())
		}
	})

	return cw
}
