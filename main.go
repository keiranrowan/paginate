package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	length := 22
	files := make([]string, 0)
	
	if len(args) > 0 {
		for k, v := range args {
			if strings.HasPrefix(args[k], "-") {
				var err error
				length, err = strconv.Atoi(strings.TrimLeft(args[k], "-"))
				if (err != nil) {
					panic("Unable to interpret line count")
				}
			} else {
				files = append(files, v)
			}
		}
	}
	
	// TTY buffer
	tty, err := os.Open("/dev/tty")
	if err != nil {
		panic("Could not open /dev/tty")
	}
	r := bufio.NewReader(tty)

	// STDIN or file buffer
	if len(files) > 0 {
		for _, v := range files {
			fp, err := os.Open(v)
			if (err != nil) {
				panic(err)
			}
			defer fp.Close()
			
			s := bufio.NewScanner(fp)

			process(length, s, r)
		}
	} else {
		s := bufio.NewScanner(os.Stdin)
		process(length, s, r)
	}
}

func process(l int, s *bufio.Scanner, r *bufio.Reader) {
	s.Split(bufio.ScanLines)

	var buf strings.Builder
		
	clear()
	
	i := 1
	for s.Scan() {
		buf.WriteString(s.Text() + "\n")
		i++
		
		if (i >= l) {
			output(&buf)
			
			wait(r)
			clear()
			i = 1
		}
	}
	output(&buf)
	wait(r)
}

func output(buf *strings.Builder) {
	fmt.Print(buf.String())
	buf.Reset()
}

func wait(r *bufio.Reader) {
	str, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if (strings.HasPrefix(str, "q")) {
		os.Exit(0);
	}
	if (strings.HasPrefix(str, "!")) {
		cmd := exec.Command("sh", "-c", strings.TrimLeft(str, "!"))
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		fmt.Println(out.String())
		wait(r)
	}
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
