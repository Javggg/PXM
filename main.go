package main

// import (
// 	"flag"
// 	"os"
// 	"path/filepath"
// 	"pxm/watcher"

// 	"github.com/fatih/color"
// )

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

// func main() {
// 	out, err := os.ReadFile("map.xml")
// 	if err != nil {
// 		panic(err)
// 	}

// 	var latestApertureTag strings.Builder
// 	var latestClosureTag strings.Builder
// 	isWritingAperture := false
// 	isWritingClosure := false
// 	lastApertureIndex := -99
// 	lastClosureIndex := 99

// 	for i := range out {
// 		if i+1 == len(out) {
// 			break
// 		}

// 		currentChar := string(out[i])
// 		nextChar := string(out[i+1])

// 		// writers
// 		// aperture
// 		if isWritingAperture {
// 			latestApertureTag.WriteString(currentChar)
// 		}

// 		// closure
// 		if isWritingClosure && currentChar != "/" {
// 			latestClosureTag.WriteString(currentChar)
// 		}

// 		// checkers
// 		// aperture
// 		if currentChar == "<" && nextChar != "!" && nextChar != "/" {
// 			latestApertureTag.Reset()
// 			isWritingAperture = true
// 		}

// 		// closure
// 		if currentChar == "<" && nextChar == "/" {
// 			latestClosureTag.Reset()
// 			isWritingClosure = true
// 			lastClosureIndex = i
// 		}

// 		// reset
// 		if nextChar == ">" {
// 			if isWritingAperture {
// 				lastApertureIndex = i + 1
// 			}

// 			isWritingAperture = false
// 			isWritingClosure = false
// 		}

// 		println(lastApertureIndex, lastClosureIndex)
// 	}

// }
