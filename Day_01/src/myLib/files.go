package myLib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WriteFile(fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if !info.IsDir() {
			_, err = file.WriteString(path + "\n")
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}

		return nil
	})
}

func CheckRemovedStrings(firstFileName, secondFileName string) {

	//first file
	firstFile, err := os.Open(firstFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer firstFile.Close()

	//second file
	secondFile, err := os.Open(secondFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer secondFile.Close()

	// check removed strings
	firstScanner := bufio.NewScanner(firstFile)
	var str string
	for firstScanner.Scan() {
		str = firstScanner.Text()
		var flag = false
		secondFile.Seek(0, 0)
		secondScanner := bufio.NewScanner(secondFile)
		for secondScanner.Scan() {
			if str == secondScanner.Text() {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("REMOVED\t", str)
		}
		err = secondScanner.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	err = firstScanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func CheckAddedStrings(firstFileName, secondFileName string) {

	//first file
	firstFile, err := os.Open(firstFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer firstFile.Close()

	//second file
	secondFile, err := os.Open(secondFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer secondFile.Close()

	// check removed strings
	secondScanner := bufio.NewScanner(secondFile)
	var str string
	for secondScanner.Scan() {
		str = secondScanner.Text()
		var flag = false
		firstFile.Seek(0, 0)
		firstScanner := bufio.NewScanner(firstFile)
		for firstScanner.Scan() {
			if str == firstScanner.Text() {
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("ADDED\t", str)
		}
		err = firstScanner.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
	err = secondScanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
