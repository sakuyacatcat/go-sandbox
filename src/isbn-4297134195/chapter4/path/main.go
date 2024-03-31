package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := `C:/path/to/file.txt`

	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Clean(`C:/path/to/..\file.txt`))
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.IsAbs(path))
	fmt.Println(filepath.IsAbs(`.\file.txt`))
	fmt.Println(filepath.Join(`C:/path`, `to/file.txt`))

	abs, err := filepath.Abs(`../file.txt`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(abs)

	abs, err = filepath.Rel(`C:\path`, `C:\path\to\file.txt`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(abs)

	fmt.Println(filepath.VolumeName(path))
}
