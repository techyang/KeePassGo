//go:build windows
// +build windows

//go:generate cmd /c "echo -ldflags -H=windowsgui > script.syso"

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
	// show the main windows
	window.Show()
	widgets.QApplication_Exec()
}
