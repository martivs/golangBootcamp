package main

import (
	"flag"
	mp "src/mypack"
	"sync"
)

func main() {

	var flgs mp.Flags2
	flag.BoolVar(&flgs.Lflag, "l", false, "Counting lines in files")
	flag.BoolVar(&flgs.Mflag, "m", false, "Counting characters in files")
	flag.BoolVar(&flgs.Wflag, "w", false, "Counting words in file")
	flag.Parse()

	mp.CheckFlags2(flgs)

	var wg sync.WaitGroup
	for _, fileName := range flag.Args() {
		wg.Add(1)
		go mp.Worker(flgs, fileName, &wg)
	}

	wg.Wait()
}
