package Util

import (
	"os"
)

func CreateFolder(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return nil
}
