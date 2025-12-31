package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"pxm/modules/maproot"
)

const TargetFolder = "target"
const BaseXML = "base.xml"

func decodeXml(file *os.File) maproot.Map {
	var m maproot.Map
	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&m); err != nil {
		panic(err)
	}

	fmt.Printf("Parsed map from %s: %+v\n", file.Name(), m)

	return m
}

func main() {
	files, err := os.ReadDir(TargetFolder)

	if err != nil {
		panic(err)
	}

	var xmlFiles []maproot.Map
	var baseIndex int = -1

	for i, entry := range files {
		if entry.IsDir() {
			continue
		}

		if entry.Name() == BaseXML {
			baseIndex = i
		}

		file, err := os.Open(TargetFolder + "/" + entry.Name())
		if err != nil {
			panic(err)
		}

		defer file.Close()

		xmlMap := decodeXml(file)

		xmlFiles = append(xmlFiles, xmlMap)

	}

	if baseIndex == -1 {
		panic("Base XML not found")
	}

	for i, xml := range xmlFiles {
		if i == baseIndex {
			continue
		}

		xmlFiles[baseIndex].Merge(xml)
	}

	out, err := xml.MarshalIndent(xmlFiles[baseIndex], "  ", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Print(string(out))
}
