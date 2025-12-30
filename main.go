package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"pxm/modules/maproot"
)

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
	targetFolder := "target"
	files, err := os.ReadDir(targetFolder)

	if err != nil {
		panic(err)
	}

	var xmlFiles []maproot.Map

	for _, entry := range files {
		if entry.IsDir() {
			continue
		}

		file, err := os.Open(targetFolder + "/" + entry.Name())
		if err != nil {
			panic(err)
		}

		defer file.Close()

		xmlMap := decodeXml(file)

		xmlFiles = append(xmlFiles, xmlMap)

	}

	for i, xml := range xmlFiles {
		if i == 0 {
			continue
		}

		xmlFiles[0].Merge(xml)
	}

	out, err := xml.MarshalIndent(xmlFiles[0], "  ", "    ") // prefix=" ", indent="    "
	if err != nil {
		panic(err)
	}

	fmt.Print(string(out))
}
