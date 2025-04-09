package file

import (
	"path/filepath"
	"strings"
)

// ChangeFileExtension changes the extension of a file path to a new extension.
// This operation is dangerous and should be used with caution.
func ChangeFileExtension(filePath string, newExtension string) string {
	fileDir := filepath.Dir(filePath)
	tmpFilename := filepath.Base(filePath)
	fileNameWithoutExt := strings.TrimSuffix(tmpFilename, filepath.Ext(tmpFilename))

	if newExtension[0] != '.' {
		newExtension = "." + newExtension
	}
	return filepath.Join(fileDir, fileNameWithoutExt+newExtension)
}
