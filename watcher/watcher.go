package watcher

import (
	"log"
	"os"
	"path/filepath"
	"pxm/merger"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

func Watch(folder string, base string, out string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					if merger.Merge(folder, base, out) {
						color.Green("Merging changes detected in '%s'", filepath.Base(event.Name))
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(folder)
	if err != nil {
		color.Red("Folder '%s' not found", folder)
		color.Yellow("Create the folder '%s' and try again.", folder)
		color.Yellow("Or specify a different folder using the '--folder' flag.")
		os.Exit(1)
	}

	color.Cyan("Listening for changes...")

	<-make(chan struct{})
}
