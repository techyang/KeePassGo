package functions

import (
	"fmt"
	"github.com/sqweek/dialog"
	"github.com/techyang/keepassgo/constants"
	"github.com/techyang/keepassgo/entity"
	"github.com/tobischo/gokeepasslib/v3"
	log "log/slog"
	"os"
	"strings"
)

func NewDatabase() {
	file, err := dialog.File().Title("Create New Password Database").SetStartFile("NewDatabase.kdbx").Filter("KeePass KDBX Files(*.kdbx)", "kdbx").Save()

	if len(file) > 0 {

		if !strings.HasSuffix(file, constants.KEEPASS_DB_EXT) {
			file += constants.KEEPASS_DB_EXT
		}

		fmt.Println("Error:", err)
		fmt.Print(file)
		masterPassword := constants.KEEPASS_DB_DEFAULT_PASSWORD

		file, err := os.Create(file)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// create the new database
		db := gokeepasslib.NewDatabase(
			gokeepasslib.WithDatabaseKDBXVersion4(),
		)
		db.Content.Meta.DatabaseName = "KDBX4"
		db.Credentials = gokeepasslib.NewPasswordCredentials(masterPassword)

		// Lock entries using stream cipher
		db.LockProtectedEntries()

		// and encode it into the file
		keepassEncoder := gokeepasslib.NewEncoder(file)
		if err := keepassEncoder.Encode(db); err != nil {
			panic(err)
		}
	}
}

func OpenDatabase(keePassTree *entity.KeePassTree, keePassTable *entity.KeePassTable) {
	exts := []string{"kdbx"} // exts := []string{"jpg", "png", "gif", "kdbx"}
	file, err := dialog.File().Title("Open").Filter("*.kdbx", exts...).Load()

	if err == nil {
		if len(file) > 0 {
			keePassTree.Clear()
			keePassTree.LoadKeePassTree(file, keePassTable)
		}
	} else {
		log.Error("open database error", err.Error())
	}
}
