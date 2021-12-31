package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func run() error {
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("> ")

		text, err := r.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read line: %+v", err)
		}
		text = strings.TrimSuffix(text, "\n")

		cmd := exec.Command(text)
		cmd.Stdout = os.Stdout
		cmd.Env = os.Environ()
		err = cmd.Run()
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
