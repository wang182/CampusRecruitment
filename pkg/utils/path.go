package utils

import (
	"path"
	"path/filepath"
	"strings"
)

func FilenameExtTrim(s string) string {
	baseName := path.Base(s)
	return strings.TrimRight(baseName, filepath.Ext(baseName))
}
