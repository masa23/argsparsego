package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/masa23/argsparsego"
)

func main() {
	args := "hoge \"fuga piyo\""

	parsed, err := argsparsego.Parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := exec.Command("echo", parsed...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
