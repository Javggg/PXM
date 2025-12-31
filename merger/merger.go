package merger

import (
	"encoding/xml"
	"fmt"
	"os"
	"pxm/modules/maproot"
)

var BaseXML = "base.xml"
var OutName = "map.xml"

func decodeXml(file *os.File) maproot.Map {
	var m maproot.Map
	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&m); err != nil {
		panic(err)
	}

	fmt.Printf("Parsed map from %s: %+v\n", file.Name(), m)

	return m
}

func Merge(path string) {
	files, err := os.ReadDir(path)

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

		file, err := os.Open(path + "/" + entry.Name())
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

	err = os.WriteFile(OutName, []byte(out), 0644)
	if err != nil {
		panic(err)
	}
}
