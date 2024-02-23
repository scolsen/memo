package main

import (
        "io"
	"fmt"
	"os"
	"strings"
)

const memofile = ".memo"

var memoloc = "/tmp/" + memofile

func main() {
   args := os.Args[1:] // ignore the program name
   // If stdin is a TTY and the user provided no argument, there is possibly
   // data fed through stdin by other means, e.g. via a pipe: command | memo ...
   // Attempt to fetch an input argument from stdin directly.
   if !isTTY(os.Stdin) {
     // No argument should be passed if we're not operating in TTY mode.
     if len(args) > 0 {
       os.Exit(1)
     }
     writeInputFromStdin()
     os.Exit(0)
   }

   // Operation in direct invocation TTY mode.
   switch len(args) {
     case 0:
       read()
     default:
       write(args)
   } 
}

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

// writeInputFromStdin attempts to write data passed on stdin to the memo file.
func writeInputFromStdin() {
  buf, err := io.ReadAll(os.Stdin)
  if err != nil {
    os.Exit(1)
  }
  write([]string{string(buf)})
}

// isTTY returns true if the given file is a terminal emulator.
func isTTY(file *os.File) bool {
  fi, err := file.Stat()
  if err != nil {
    return false
  }

  return fi.Mode() & os.ModeCharDevice != 0
}

