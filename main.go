package main

import (
	"github.com/techyang/keepassgo/functions"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	///start
	widgets.NewQApplication(len(os.Args), os.Args)
	window := functions.InitMainWindow()
	window.Show()
	widgets.QApplication_Exec()
}
