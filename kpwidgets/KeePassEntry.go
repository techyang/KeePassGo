package kpwidgets

import "time"

type KeePassEntry struct {
	Title        string
	UserName     string
	Password     string
	URL          string
	Notes        string
	Expires      time.Time
	ExpiresSeted bool
}

func NewKeePassEntry() *KeePassEntry {
	entry := &KeePassEntry{}
	return entry
}
