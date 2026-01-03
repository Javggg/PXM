package merger

import (
	"encoding/xml"
	"os"
	"pxm/modules/root"
)

func decodeXml(file *os.File) root.Map {
	var m root.Map
	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&m); err != nil {
		panic(err)
	}

	return m
}

func Merge(folder string, base string, out string) {
	files, err := os.ReadDir(folder)

	if err != nil {
		panic(err)
	}

	var xmlFiles []root.Map
	var baseIndex int = -1

	for i, entry := range files {
		if entry.IsDir() {
			continue
		}

		if entry.Name() == base {
			baseIndex = i
		}

		file, err := os.Open(folder + "/" + entry.Name())
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

	final, err := xml.MarshalIndent(xmlFiles[baseIndex], "  ", "    ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(out, []byte(final), 0644)
	if err != nil {
		panic(err)
	}
}
