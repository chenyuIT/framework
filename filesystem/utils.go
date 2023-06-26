package filesystem

import (
	"path/filepath"
	"strings"

	"github.com/chenyuIT/framework/contracts/filesystem"
	"github.com/chenyuIT/framework/support/file"
)

func fullPathOfFile(filePath string, source filesystem.File, name string) (string, error) {
	extension := filepath.Ext(name)
	if extension == "" {
		var err error
		extension, err = file.Extension(source.File(), true)
		if err != nil {
			return "", err
		}

		return filepath.Join(filePath, strings.TrimSuffix(strings.TrimPrefix(filepath.Base(name), string(filepath.Separator)), string(filepath.Separator))+"."+extension), nil
	} else {
		return filepath.Join(filePath, strings.TrimPrefix(filepath.Base(name), string(filepath.Separator))), nil
	}
}
