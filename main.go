package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/techyang/keepassgo/functions"
	"github.com/techyang/keepassgo/kpwidgets"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	///start
	functions.SetupLogger()
	log.Info("begin to start app ...")
	widgets.NewQApplication(len(os.Args), os.Args)
	window := kpwidgets.InitMainWindow()
	window.Show()
	widgets.QApplication_Exec()
}
