package main

import (
	"flag"
	"fmt"
	"log"
	mlb "src/myLib"
	"strings"
	"time"
)

func main() {

	//flags
	var firstFileName, secondFileName string
	flag.StringVar(&firstFileName, "old", "", "file name")
	flag.StringVar(&secondFileName, "new", "", "file name")
	flag.Parse()

	if !strings.HasSuffix(firstFileName, ".txt") || !strings.HasSuffix(secondFileName, ".txt") {
		log.Fatalln("Specify txt files, please")
	}

	mlb.WriteFile(firstFileName)
	for i := 30; i >= 0; i-- {
		fmt.Printf("\r%d sec.", i)
		time.Sleep(time.Second)
	}
	mlb.WriteFile(secondFileName)

	fmt.Println("\rLIST OF CHANGES:\n_______________")
	mlb.CheckAddedStrings(firstFileName, secondFileName)
	mlb.CheckRemovedStrings(firstFileName, secondFileName)

}
