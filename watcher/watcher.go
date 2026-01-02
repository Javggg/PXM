package watcher

import (
	"log"
	"pxm/merger"

	"github.com/fsnotify/fsnotify"
)

func Watch(path string) {
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
				// log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					merger.Merge(path)
					// log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
