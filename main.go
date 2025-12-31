package main

import (
	"os"
	"pxm/watcher"
)

var TargetFolder = "target"

func main() {
	if len(os.Args) >= 2 {
		TargetFolder = os.Args[1]
	}

	watcher.Watch(TargetFolder)
}
