package functions

import (
	"fmt"
	"github.com/therecipe/qt/widgets"
)

func DoNewAction(window *widgets.QMainWindow) {
	msgBox := widgets.NewQMessageBox(nil)
	msgBox.SetWindowTitle("Message Box")
	msgBox.SetText("This is a message box.")
	msgBox.SetInformativeText("This is a message box with text information.")
	msgBox.SetStandardButtons(widgets.QMessageBox__Ok | widgets.QMessageBox__Cancel)
	msgBox.SetDefaultButton2(widgets.QMessageBox__Ok)

	// Connect the clicked signal of the buttons
	msgBox.ConnectButtonClicked(func(button *widgets.QAbstractButton) {
		if button.Text() == "OK" {
			newFileBox := widgets.NewQFileDialog2(window, "新建", "", "*.txt")
			newFileBox.Show()
			newFileBox.ConnectFileSelected(func(file string) {
				fmt.Print(file)
			})
		} else if button.Text() == "&Cancel" {
			// Handle the logic for the Cancel button
			// ...
		}

		// Close the message box
		msgBox.Close()
	})

	// Show the message box
	msgBox.Exec()
}
