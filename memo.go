package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const memofile = ".memo"
var memoloc = "/tmp/" + memofile

func write(args []string) {
	content := strings.Join(args, " ")
	file, err := os.OpenFile(memoloc, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	if _, e := file.WriteString(content); e != nil {
		log.Fatal(e)
	}
}

func read() {
	buf, e := ioutil.ReadFile(memoloc)

	if os.IsNotExist(e) {
		fmt.Println("Nothing memorized!")
		os.Exit(0)
	}

	if e != nil {
		log.Fatal(e)
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
