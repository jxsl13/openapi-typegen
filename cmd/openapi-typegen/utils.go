package main

import (
	"os"
	"path/filepath"
)

func AppName() string {
	appNameWithExt := filepath.Base(os.Args[0])
	ext := filepath.Ext(appNameWithExt)
	appNameWithoutExt := appNameWithExt[:len(appNameWithExt)-len(ext)]
	return appNameWithoutExt
}
