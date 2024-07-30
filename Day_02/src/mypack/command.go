package mypack

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Echo() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		log.Fatal(sc.Err())
	} else {
		fmt.Println(sc.Text())
	}
}

func Runner() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		var cmd exec.Cmd
		if len(os.Args) == 2 {
			cmd = *exec.Command(os.Args[1], sc.Text())
		} else {
			cmd = *exec.Command(os.Args[1], append(os.Args[2:], sc.Text())...)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
