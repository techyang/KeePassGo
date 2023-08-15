package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type DuplicateEntryDialog struct {
	*widgets.QDialog
	AppendCopyCheck *widgets.QCheckBox
	RepeatUserName  *widgets.QCheckBox
	CopyHistory     *widgets.QCheckBox
	ButtonBox       *widgets.QDialogButtonBox
}

func NewDuplicationOptionsDialog() *DuplicateEntryDialog {
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Duplication Options")

	optionsDialog := &DuplicateEntryDialog{
		QDialog: dialog,
	}

	// Create the widgets
	optionsDialog.AppendCopyCheck = widgets.NewQCheckBox2("Append \"-Copy\" to entry titles", nil)
	optionsDialog.RepeatUserName = widgets.NewQCheckBox2("Enable auto-type for this entry", nil)
	optionsDialog.CopyHistory = widgets.NewQCheckBox2("Copy history", nil)

	// Create the layout
	layout := widgets.NewQVBoxLayout2(dialog)
	layout.AddWidget(optionsDialog.AppendCopyCheck, 0, core.Qt__AlignLeft)
	layout.AddWidget(optionsDialog.RepeatUserName, 0, core.Qt__AlignLeft)

	// Add the help label and link
	helpLabel1 := widgets.NewQLabel2("If this option is enabled, the copies will reference", nil, 0)
	helpLabel1.SetWordWrap(true)
	layout.AddWidget(helpLabel1, 0, core.Qt__AlignLeft)

	helpLabel2 := widgets.NewQLabel2("Help: Field References", nil, 0)
	helpLabel2.SetTextInteractionFlags(core.Qt__LinksAccessibleByMouse)
	helpLabel2.ConnectLinkActivated(func(link string) {
		doLinkClicked(link)
	})
	layout.AddWidget(helpLabel2, 0, core.Qt__AlignLeft)

	// Add the separator
	separator := widgets.NewQFrame(nil, 0)
	separator.SetFrameShape(widgets.QFrame__HLine)
	separator.SetLineWidth(20)
	layout.AddWidget(separator, 0, core.Qt__AlignLeft)

	// Add the CopyHistory checkbox
	layout.AddWidget(optionsDialog.CopyHistory, 0, core.Qt__AlignLeft)

	// Add the separator (use QLabel to simulate a horizontal line)
	separatorLabel := widgets.NewQLabel2("<hr />", nil, 0)
	separatorLabel.SetTextFormat(core.Qt__RichText)
	separatorLabel.AdjustSize()
	separatorLabel.SetFixedWidth(400)
	layout.AddWidget(separatorLabel, 0, core.Qt__AlignLeft)

	// Add the button box
	optionsDialog.ButtonBox = widgets.NewQDialogButtonBox(dialog)
	okButton := optionsDialog.ButtonBox.AddButton3(widgets.QDialogButtonBox__Ok)
	cancelButton := optionsDialog.ButtonBox.AddButton3(widgets.QDialogButtonBox__Cancel)

	// Set the button text
	okButton.SetText("OK")
	cancelButton.SetText("Cancel")

	// Connect the button box's accepted signal

	layout.AddWidget(optionsDialog.ButtonBox, 0, core.Qt__AlignRight)

	return optionsDialog
}

func doLinkClicked2(link string) {
	// Handle link clicked event
	fmt.Println("Link clicked:", link)
}

// click Ok Button action
func (dialog *DuplicateEntryDialog) DoAccepted(tableWidget *KeePassTable) {
	row := tableWidget.CurrentRow()
	if row >= 0 {
		values := getRowData(tableWidget, row)

		if dialog.AppendCopyCheck.IsChecked() {
			values[0] += "-Copy"
		}

		tableWidget.InsertRow(row + 1)
		for col := 0; col < tableWidget.ColumnCount(); col++ {
			tableWidget.SetItem(row+1, col, widgets.NewQTableWidgetItem2(values[col], 0))
		}
		tableWidget.SelectRow(row + 1)
	}
	dialog.Accept()
}
