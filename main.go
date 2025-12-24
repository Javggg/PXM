package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"pxm/modules/maproot"
)

func decodeXml(file *os.File) {
	var m maproot.Map
	decoder := xml.NewDecoder(file)

	if err := decoder.Decode(&m); err != nil {
		panic(err)
	}

	fmt.Printf("Parsed map from %s: %+v\n", file.Name(), m)

	output, err := xml.MarshalIndent(m, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(xml.Header) + string(output))
}

func main() {
	targetFolder := "target"
	files, err := os.ReadDir(targetFolder)

	if err != nil {
		panic(err)
	}

	for _, entry := range files {
		if entry.IsDir() {
			continue
		}

		file, err := os.Open(targetFolder + "/" + entry.Name())
		if err != nil {
			panic(err)
		}

		defer file.Close()

		decodeXml(file)
	}
}
