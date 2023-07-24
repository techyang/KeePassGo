package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os/exec"
)

type AutoTypeTab struct {
	Widget                            *widgets.QWidget
	EnableAutoTypeCheckbox            *widgets.QCheckBox
	InheritDefaultTypeRadio           *widgets.QRadioButton
	OverrideDefaultSequenceRadio      *widgets.QRadioButton
	OverrideDefaultSequenceInput      *widgets.QLineEdit
	OverrideDefaultSequenceEditButton *widgets.QPushButton
	WindowSequenceTable               *widgets.QTableView
	TwoChannelAutoTypeCheckbox        *widgets.QCheckBox
	TableWidget                       *widgets.QTableWidget
}

func NewAutoTypeTab() *AutoTypeTab {
	cw := &AutoTypeTab{
		Widget:                       widgets.NewQWidget(nil, 0),
		EnableAutoTypeCheckbox:       widgets.NewQCheckBox2("Enable auto-type for this entry", nil),
		InheritDefaultTypeRadio:      widgets.NewQRadioButton2("Inherit default auto-type sequence from group", nil),
		OverrideDefaultSequenceRadio: widgets.NewQRadioButton2("Override default sequence", nil),
		OverrideDefaultSequenceInput: widgets.NewQLineEdit(nil),
		TableWidget:                  widgets.NewQTableWidget(nil),

		OverrideDefaultSequenceEditButton: widgets.NewQPushButton2("Edit", nil),
		TwoChannelAutoTypeCheckbox:        widgets.NewQCheckBox2("Two-channel auto-type obfuscation", nil),
	}

	// Layout
	vBoxLayout := widgets.NewQVBoxLayout2(cw.Widget)

	vBoxLayout.AddWidget(cw.EnableAutoTypeCheckbox, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.InheritDefaultTypeRadio, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.OverrideDefaultSequenceRadio, 0, core.Qt__AlignLeft)

	overrideDefaultSequenceLayout := widgets.NewQHBoxLayout2(nil)
	overrideDefaultSequenceLayout.AddSpacing(40)
	cw.OverrideDefaultSequenceInput.SetFixedWidth(400)
	overrideDefaultSequenceLayout.AddWidget(cw.OverrideDefaultSequenceInput, 0, core.Qt__AlignLeft)
	overrideDefaultSequenceLayout.AddWidget(cw.OverrideDefaultSequenceEditButton, 1, core.Qt__AlignLeft)
	vBoxLayout.AddLayout(overrideDefaultSequenceLayout, 0)

	setTableWidget(cw.TableWidget, vBoxLayout)
	/*vBoxLayout.AddWidget(cw.TagsLabel, 0, core.Qt__AlignLeft)
	vBoxLayout.AddWidget(cw.TagsInput, 0, core.Qt__AlignLeft)*/

	twoChannelAutoTypeLayout := widgets.NewQHBoxLayout2(nil)
	twoChannelAutoTypeLabel := widgets.NewQLabel2("link", nil, 0)
	twoChannelAutoTypeLabel.SetText("<a href=\"D:\\Program Files\\TotalCMDPortable\\App\\totalcmd\\Plugins\\wcx\\Total7zip\\7-zip.chm\">What is this?</a>")

	twoChannelAutoTypeLabel.SetTextInteractionFlags(core.Qt__LinksAccessibleByMouse)

	twoChannelAutoTypeLabel.ConnectLinkActivated(func(link string) {
		//link = "D:\\Program Files\\TotalCMDPortable\\App\\totalcmd\\Plugins\\wcx\\Total7zip\\7-zip.chm"
		if err := exec.Command("hh.exe", link).Start(); err != nil {
			fmt.Println("Error opening .chm file:", err)
		}
	})
	twoChannelAutoTypeLayout.AddWidget(cw.TwoChannelAutoTypeCheckbox, 0, core.Qt__AlignLeft)
	twoChannelAutoTypeLayout.AddWidget(twoChannelAutoTypeLabel, 0, core.Qt__AlignTop)
	vBoxLayout.AddLayout(twoChannelAutoTypeLayout, 0)

	// Connect the "Pick Color" buttons clicked signal to open the color selection dialogs
	cw.InheritDefaultTypeRadio.ConnectClicked(func(checked bool) {
		cw.OverrideDefaultSequenceInput.SetDisabled(checked)
		cw.OverrideDefaultSequenceEditButton.SetDisabled(checked)
	})

	cw.OverrideDefaultSequenceRadio.ConnectClicked(func(checked bool) {
		cw.OverrideDefaultSequenceInput.SetDisabled(!checked)
		cw.OverrideDefaultSequenceEditButton.SetDisabled(!checked)
	})

	/*	cw.BackgroundColorPickerBtn.ConnectClicked(func(checked bool) {
		colorDialog := widgets.NewQColorDialog2(gui.NewQColor3(255, 255, 255, 255), nil)
		colorDialog.Exec()
		if colorDialog != nil {
			cw.BackgroundColorPickerBtn.SetStyleSheet("background-color: " + colorDialog.CurrentColor().Name())
		}
	})*/

	return cw
}

func setTableWidget(tableWidget *widgets.QTableWidget, vBoxLayout *widgets.QVBoxLayout) {
	tableLable := widgets.NewQLabel2("Use custom sequences for specific window:", nil, 0)
	vBoxLayout.AddWidget(tableLable, 0, core.Qt__AlignLeft)

	tableWidget.SetColumnCount(2)
	tableWidget.SetFixedWidth(600)
	// Set the header labels
	headerLabels := []string{"Taget Window", "Sequence"}
	tableWidget.SetHorizontalHeaderLabels(headerLabels)

	tableButtonVLayout := getTableButtonVLayout()

	tableHLayout := widgets.NewQHBoxLayout2(nil)
	tableHLayout.AddWidget(tableWidget, 0, core.Qt__AlignLeft)
	tableHLayout.AddLayout(tableButtonVLayout, 0)
	vBoxLayout.AddLayout(tableHLayout, 0)
}

func getTableButtonVLayout() *widgets.QVBoxLayout {
	tableButtonVLayout := widgets.NewQVBoxLayout2(nil)
	addButton := widgets.NewQPushButton2("Add", nil)
	editButton := widgets.NewQPushButton2("Edit", nil)
	moveButton := widgets.NewQPushButton2("Move", nil)
	upButton := widgets.NewQPushButton2("Up", nil)
	downButton := widgets.NewQPushButton2("Down", nil)
	tableButtonVLayout.AddWidget(addButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(editButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(moveButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(upButton, 0, core.Qt__AlignLeft)
	tableButtonVLayout.AddWidget(downButton, 0, core.Qt__AlignLeft)
	return tableButtonVLayout
}
