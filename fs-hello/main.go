package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	currentDirectory := "."
	fileSystem := os.DirFS(currentDirectory)
	fmt.Printf("Listing files in: %s\n", fileSystem)
	fs.WalkDir(fileSystem, ".", func(filename string, entry fs.DirEntry, err error) error {
		info, _ := entry.Info()
		fileType := entry.Type()
		fmt.Printf(
			"filename=%s, mode=%s, name=%s, type=%s\n",
			filename,
			info.Mode(),
			info.Name(),
			fileType.String(),
		)

		return nil
	})
}
