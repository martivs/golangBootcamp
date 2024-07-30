package mypack

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Flags2 struct {
	Lflag, Mflag, Wflag bool
}

func CheckFlags2(flgs Flags2) {

	if flag.NFlag() == 0 {
		log.Fatal("Specify flag, please (-h help)")
	}

	if flag.NFlag() > 1 {
		log.Fatal("No more than one flag, please")
	}

	if flag.NArg() == 0 {
		log.Fatal("Specify txt file(s), please")
	}

	for _, arg := range flag.Args() {
		if !strings.HasSuffix(arg, ".txt") {
			log.Fatal("Only txt files, please")
		}
	}
}

func Worker(flgs Flags2, fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	switch {
	default:
		fmt.Println("Something wrong")
	case flgs.Lflag:
		fmt.Printf("%d %s\n", linesCounter(fileName), fileName)
	case flgs.Mflag:
		fmt.Printf("%d %s\n", charCounter(fileName), fileName)
	case flgs.Wflag:
		fmt.Printf("%d %s\n", wordsCounter(fileName), fileName)
	}
}

func charCounter(fileName string) int {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	charCount := 0
	rd := bufio.NewReader(file)
	for {
		_, _, err := rd.ReadRune()
		if err != nil {
			break
		}
		charCount++
	}

	return charCount
}

func wordsCounter(fileName string) int {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordsCount := 0
	sw := bufio.NewScanner(file)
	sw.Split(bufio.ScanWords)
	for sw.Scan() {
		wordsCount++
	}

	return wordsCount
}

func linesCounter(fileName string) int {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lineCount := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lineCount++
	}

	return lineCount
}
