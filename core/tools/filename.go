package tools

import "path/filepath"

func RemoveExtensionFromFielname(filename string) string {
	return filename[:len(filename)-len(filepath.Ext(filename))]
}
