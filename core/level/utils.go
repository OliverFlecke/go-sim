package level

import (
	"io/fs"
	"path/filepath"
)

func GetMaps(root_directory string, callback func(level string)) {
	filepath.Walk(root_directory, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			callback(path)
		}

		return err
	})
}
