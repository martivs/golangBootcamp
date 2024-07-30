package mypack

import (
	// "bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Flags1 struct {
	Dflag, Fflag, Slflag bool
	ExtStr               string
}

func CheckFlags1(flgs Flags1) {
	if flgs.ExtStr != "" && !flgs.Fflag {
		fmt.Println(flgs.ExtStr)
		log.Fatal("The -ext flag works only with the -f flag")
	}
	if flag.NArg() == 0 {
		log.Fatal("PATH!")
	}
}

func Report1(flgs Flags1) {

	filepath.Walk(flag.Arg(0), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		switch {
		default:
			fmt.Println(path)
		case flgs.Dflag:
			if info.IsDir() {
				fmt.Println(path)
			}
		case flgs.Fflag && flgs.ExtStr == "":
			if info.Mode().IsRegular() {
				fmt.Println(path)
			}
		case flgs.Fflag && flgs.ExtStr != "":
			if info.Mode().IsRegular() && strings.HasSuffix(path, "."+flgs.ExtStr) {
				fmt.Println(path)
			}
		case flgs.Slflag:
			if info.Mode()&os.ModeSymlink != 0 {
				link, err := os.Readlink(path)
				if err != nil {
					fmt.Printf("%s -> [broken]\n", path)
				} else {
					fmt.Printf("%s -> %s\n", path, link)
				}
			}
		}

		return nil
	})
}
