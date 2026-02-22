package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type gote struct {
	title, desc string
}

func (i gote) Title() string       { return i.title }
func (i gote) Description() string { return i.desc }
func (i gote) FilterValue() string { return i.title }

func listGotes() []list.Item {
	gotes := make([]list.Item, 0)

	entries, err := os.ReadDir(notesDir)
	if err != nil {
		log.Fatal("Error reading your gotes")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			fileInfo, err := entry.Info()
			if err != nil {
				continue
			}
			fileModTime := fileInfo.ModTime().Format("2006-01-02 15:04:05")
			fileName := entry.Name()
			gotes = append(gotes, gote{title: fileName, desc: fmt.Sprintf("Modified: %s", fileModTime)})
		}
	}

	return gotes
}
