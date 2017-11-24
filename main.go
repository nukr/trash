package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	home := os.Getenv("HOME")
	for _, arg := range os.Args[1:] {
		filename := filepath.Base(arg)
		new := home + "/.Trash/" + filename
		err := mv(arg, new)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func mv(old, new string) error {
	err := os.Rename(old, new)
	if linkerr, ok := err.(*os.LinkError); ok {
		if linkerr.Err == syscall.EEXIST {
			return mv(old, new+time.Now().Format(time.RFC3339))
		}
		return linkerr
	}
	if err != nil {
		return err
	}
	return nil
}
