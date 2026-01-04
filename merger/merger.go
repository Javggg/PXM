package merger

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"pxm/modules/root"

	"github.com/fatih/color"
)

func decodeXml(file *os.File) (root.Map, error) {
	var m root.Map
	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&m); err != nil {
		color.Red(filepath.Base(file.Name()) + ": " + err.Error())
		return m, err
	}

	return m, nil
}

func Merge(folder string, base string, out string) bool {
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

		xmlMap, err := decodeXml(file)
		if err != nil {
			return false
		}

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

	return true
}
