// Edit file with editor defined in environment variable $EDITOR.
// Copyright (c) 2017 Ruda Moura

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"syscall"
)

var (
	progname = os.Args[0]
	editor   = os.Getenv("EDITOR")
	err      error
)

func findEditor(editor string) (string, error) {
	for _, p := range strings.Split(os.Getenv("PATH"), ":") {
		fmt.Printf("%s\n", p)
		if _, err := os.Stat(path.Join(p, editor)); err == nil {
			return path.Join(p, editor), nil
		}
	}
	err := errors.New("Editor not found")
	return "", err
}

func main() {
	if editor == "" {
		editor = "vi"
	}
	if !path.IsAbs(editor) {
		if editor, err = findEditor(editor); err != nil {
			fmt.Printf("%s: %s: %s\n", progname, err, editor)
			os.Exit(1)
		}
	}
	os.Args[0] = path.Base(editor)
	err := syscall.Exec(editor, os.Args, os.Environ())
	fmt.Printf("%s: Exec error: %s '%s'\n", progname, err, editor)
	os.Exit(1)
}
