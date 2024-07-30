package main

import (
	"os"
	mp "src/mypack"
)

func main() {

	switch {
	case len(os.Args) == 1:
		mp.Echo()
	case len(os.Args) > 1:
		mp.Runner()
	}

}
