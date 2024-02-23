package main

import (
	"fmt"
	"os"
	"strings"
)

const memofile = ".memo"

var memoloc = "/tmp/" + memofile

func write(args []string) {
	content := strings.Join(args, " ")
	err := os.WriteFile(memoloc, []byte(content), 0644)
	if err != nil {
		os.Exit(1)
	}
}

func read() {
	buf, err := os.ReadFile(memoloc)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			fmt.Println("Nothing memorized!")
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}

	fmt.Printf("%s", buf)
}

func main() {
	switch len(os.Args) {
	case 1:
		read()
	default:
		args := os.Args[1:]
		write(args)
	}
}
