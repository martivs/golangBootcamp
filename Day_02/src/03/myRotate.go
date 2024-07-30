package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	mp "src/mypack"
	"sync"
)

func main() {

	archDir := flag.String("a", "", "Archive directory")
	flag.Parse()

	if len(os.Args) < 2 {
		log.Fatal("Specify log files, please")
	}

	switch {
	case *archDir == "":
		err := mp.Translator(os.Args[1], "")
		if err != nil {
			fmt.Println(err)
		}
	case *archDir != "":
		err := os.MkdirAll(*archDir, 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
		var wg sync.WaitGroup
		for _, logFileName := range flag.Args() {
			wg.Add(1)
			go func(logFileName string) {
				defer wg.Done()
				err := mp.Translator(logFileName, *archDir)
				if err != nil {
					fmt.Println(err)
				}
			}(logFileName)
		}
		wg.Wait()
	}
}
