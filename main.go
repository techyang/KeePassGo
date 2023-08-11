package main

import (
	"github.com/techyang/keepassgo/kpwidgets"
	"github.com/therecipe/qt/widgets"
	log "log/slog"
	"os"
)

func main() {
	///start
	//functions.SetupLogger()

	log.Info("begin to start app ,please wait...")
	widgets.NewQApplication(len(os.Args), os.Args)
	window := kpwidgets.InitMainWindow()
	window.Show()
	widgets.QApplication_Exec()
}
