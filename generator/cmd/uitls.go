package cmd

import (
	"os"
	"path/filepath"
)

type Filter func(path string) bool

func Ls(path string, filter Filter) []string {
	var files []string
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filter(path) {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func MustOverrideFile(path string, bytes []byte) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(bytes)
}
