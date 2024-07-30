package main

import (
	"flag"
	mp "src/mypack"
)

func main() {

	var flgs mp.Flags1
	flag.BoolVar(&flgs.Dflag, "d", false, "Display directories")
	flag.BoolVar(&flgs.Fflag, "f", false, "Display files")
	flag.BoolVar(&flgs.Slflag, "sl", false, "Display symlinks")
	flag.StringVar(&flgs.ExtStr, "ext", "", "Display files by extentions. Only with flag -f")
	flag.Parse()

	mp.CheckFlags1(flgs)
	mp.Report1(flgs)
}
