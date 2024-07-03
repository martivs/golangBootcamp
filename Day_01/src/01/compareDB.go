package main

import (
	"flag"
	"log"
	"strings"

	mlb "src/myLib"
)

func main() {

	//flags
	var oldFileName, newFileName string
	flag.StringVar(&oldFileName, "old", "", "file name")
	flag.StringVar(&newFileName, "new", "", "file name")
	flag.Parse()

	var xOld, xNew mlb.XmlReader
	var jOld, jNew mlb.JsonReader

	switch {
	default:
		log.Fatalln("Specify old and new JSON or XML files, please")
	case strings.HasSuffix(oldFileName, ".xml") && strings.HasSuffix(newFileName, ".xml"):
		xOld.Read(oldFileName)
		jOld = xOld.XmlToJson()
		xNew.Read(newFileName)
		jNew = xNew.XmlToJson()
		mlb.CompareJson(&jOld, &jNew)
	case strings.HasSuffix(oldFileName, ".json") && strings.HasSuffix(newFileName, ".json"):
		jOld.Read(oldFileName)
		jNew.Read(newFileName)
		mlb.CompareJson(&jOld, &jNew)
	case strings.HasSuffix(oldFileName, ".xml") && strings.HasSuffix(newFileName, ".json"):
		xOld.Read(oldFileName)
		jOld = xOld.XmlToJson()
		jNew.Read(newFileName)
		mlb.CompareJson(&jOld, &jNew)
	case strings.HasSuffix(oldFileName, ".json") && strings.HasSuffix(newFileName, ".xml"):
		jOld.Read(oldFileName)
		xNew.Read(newFileName)
		jNew = xNew.XmlToJson()
		mlb.CompareJson(&jOld, &jNew)
	}
}
