package entity

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type IconLabel struct {
	*widgets.QWidget
	iconLabel *widgets.QLabel
	textLabel *widgets.QLabel
}

func NewIconLabel(iconPath, text string) *IconLabel {
	icon := gui.NewQIcon5(iconPath)
	iconLabel := widgets.NewQLabel2("", nil, 0)
	iconLabel.SetPixmap(icon.Pixmap(core.NewQSize2(32, 32), gui.QIcon__Normal, 0))

	textLabel := widgets.NewQLabel2(text, nil, 0)

	layout := widgets.NewQHBoxLayout2(nil)
	layout.AddWidget(iconLabel, 0, core.Qt__AlignLeft)
	layout.AddWidget(textLabel, 0, core.Qt__AlignLeft)

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	return &IconLabel{
		QWidget:   widget,
		iconLabel: iconLabel,
		textLabel: textLabel,
	}
}
