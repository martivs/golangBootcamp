package main

import (
	"flag"
	"log"
	"strings"

	mlb "src/myLib"
)

func main() {

	//flags
	var fileName string
	flag.StringVar(&fileName, "f", "", "file name")
	flag.Parse()

	var dbr mlb.DBReader

	switch {
	default:
		log.Fatalln("Specify JSON or XML file, please")
	case strings.HasSuffix(fileName, ".xml"):
		dbr = &mlb.XmlReader{}
	case strings.HasSuffix(fileName, ".json"):
		dbr = &mlb.JsonReader{}
	}

	dbr.Read(fileName)
	dbr.Print()
}
