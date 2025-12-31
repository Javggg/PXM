package main

import (
	"os"
	"pxm/watcher"
)

var TargetFolder = "src"

func main() {
	if len(os.Args) >= 2 {
		TargetFolder = os.Args[1]
		// TODO: check if it's the same folder as the executable
	}

	watcher.Watch(TargetFolder)
}
