package main

import (
	"flag"
	"os"
	"path/filepath"
	"pxm/watcher"

	"github.com/fatih/color"
)

func main() {
	folderFlag := flag.String("folder", "src", "Folder containing the XML files")
	baseFileFlag := flag.String("base", "base.xml", "Name of the base XML file")
	outFlag := flag.String("out", "map.xml", "Name of the final XML file")
	flag.Parse()

	path, err := filepath.Abs(*folderFlag)
	if err != nil {
		color.Red("There was an error obtaining the route", err)
		os.Exit(1)
	}

	watcher.Watch(path, *baseFileFlag, *outFlag)
}
