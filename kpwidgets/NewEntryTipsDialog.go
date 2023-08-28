package kpwidgets

import (
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type NewEntryTipsDialog struct {
	*widgets.QDialog
	ButtonBox *widgets.QDialogButtonBox
}

func NewNewEntryTipsDialog() *NewEntryTipsDialog {
	dialog := widgets.NewQDialog(nil, 0)
	dialog.SetWindowTitle("Duplication Options")

	optionsDialog := &NewEntryTipsDialog{
		QDialog: dialog,
	}

	// Add the icon and label to the layout
	//iconLabel := widgets.NewQLabel(nil, 0)
	//iconLabel.SetPixmap(gui.NewQPixmap("path/to/info-icon.png"))

	// Create the widgets
	// Add the separator
	separator := widgets.NewQFrame(nil, 0)
	separator.SetFrameShape(widgets.QFrame__HLine)
	separator.SetLineWidth(20)
	// Create the layout
	layout := widgets.NewQVBoxLayout2(dialog)

	// Add the help label and link
	helpLabel1 := widgets.NewQLabel2(kpResources.DatabaseFileIntro, nil, 0)
	//blankLabel := widgets.NewQLabel2("  ", nil, 0)

	helpLabel1.SetWordWrap(true)
	helpLabel1.AdjustSize()
	layout.AddWidget(helpLabel1, 0, core.Qt__AlignLeft)

	//layout.AddWidget(blankLabel, 0, core.Qt__AlignLeft)
	//layout.AddWidget(separator, 0, core.Qt__AlignLeft)

	helpLabel2 := widgets.NewQLabel2(kpResources.DatabaseFileRem, nil, 0)
	helpLabel2.SetWordWrap(true)
	helpLabel2.AdjustSize()
	helpLabel3 := widgets.NewQLabel2(kpResources.BackupDatabase, nil, 0)
	helpLabel3.SetWordWrap(true)
	helpLabel3.AdjustSize()
	layout.AddWidget(helpLabel2, 0, core.Qt__AlignLeft)

	//layout.AddWidget(blankLabel, 0, core.Qt__AlignLeft)
	// Add the CopyHistory checkbox
	layout.AddWidget(helpLabel3, 0, core.Qt__AlignLeft)

	// Add the button box
	optionsDialog.ButtonBox = widgets.NewQDialogButtonBox(dialog)
	okButton := optionsDialog.ButtonBox.AddButton3(widgets.QDialogButtonBox__Ok)
	cancelButton := optionsDialog.ButtonBox.AddButton3(widgets.QDialogButtonBox__Cancel)

	// Set the button text
	okButton.SetText("OK")
	cancelButton.SetText("Cancel")

	// Connect the button box's accepted signal

	layout.AddWidget(optionsDialog.ButtonBox, 0, core.Qt__AlignRight)
	optionsDialog.Resize2(400, 300)
	return optionsDialog
}

func (d *NewEntryTipsDialog) doAccepted() {
	// Do something when the OK button is clicked
	fmt.Println("OK button clicked")
}
