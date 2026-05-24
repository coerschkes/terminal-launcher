package main

import (
	"os"
	"os/exec"
)

const SOCKET = "/tmp/hyperland-daemon.sock"

func main() {
	pwd := NewPwdLoader(SOCKET).LoadFromSocket()
	args := append([]string{"--working-directory", pwd}, os.Args[1:]...)

	cmd := exec.Command("alacritty", args...)

	if err := cmd.Start(); err != nil {
		panic(err)
	}
}
