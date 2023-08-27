package kpwidgets

import (
	"github.com/sqweek/dialog"
	log "log/slog"
)

func OpenDatabase() {
	exts := []string{"kdbx"} // exts := []string{"jpg", "png", "gif", "kdbx"}
	file, err := dialog.File().Title("Open").Filter("*.kdbx", exts...).Load()

	if err == nil {
		if len(file) > 0 {
			TreeWidget.Clear()
			TreeWidget.LoadKeePassTree(file, TableWidget)
		}
	} else {
		log.Error("open database error", err.Error())
	}
}
