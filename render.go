package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Comamoca/chuno"
	"github.com/fatih/color"
)

func renderCmd(file string, isDark bool) error {
	html, err := chuno.Build(file, isDark)
	if err != nil {
		log.Fatal("An error occurred while rendering the HTML.")
		return err
	}

	fileName := filepath.Base(file[:len(file)-len(filepath.Ext(file))])
	err = write(fileName+".html", html)
	if err != nil {
		log.Fatal("An error occurred while writing the file.")
		return err
	}

	color.Cyan("File writing completed✨")

	return nil
}

func write(path string, content []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()

	f.Write(content)

	return nil
}
