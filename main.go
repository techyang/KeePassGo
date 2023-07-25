package main

import (
	"github.com/techyang/keepassgo/kpwidgets"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	window := kpwidgets.InitMainWindow()
	window.Show()
	widgets.QApplication_Exec()
}
