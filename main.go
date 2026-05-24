package main

import (
	"os/exec"
)

const SOCKET = "/tmp/hyperland-daemon.sock"

func main() {
	pwd := NewPwdLoader(SOCKET).LoadFromSocket()
	cmd := exec.Command("alacritty", "--working-directory", pwd)

	if err := cmd.Start(); err != nil {
		panic(err)
	}
}
