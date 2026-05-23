package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type PwdLoader struct {
	socket string
}

func NewPwdLoader(socket string) *PwdLoader {
	return &PwdLoader{socket: socket}
}

func (l *PwdLoader) LoadFromSocket() string {
	c := l.dialSocket()

	if c == nil {
		return ""
	}

	defer c.Close()

	return l.readPwd(c)
}

func (l *PwdLoader) dialSocket() net.Conn {
	conn, err := net.Dial("unix", l.socket)

	if err != nil {
		fmt.Fprintln(os.Stderr, "connect error:", err)
		return nil
	}
	return conn
}

func (l *PwdLoader) readPwd(c net.Conn) string {
	fmt.Fprintln(c, "pwd-retrieve")

	c.SetReadDeadline(time.Now().Add(1 * time.Second))

	buf := make([]byte, 1024)
	n, err := c.Read(buf)

	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		return ""
	}

	return l.normalizeHomeDir(string(buf[:n]))
}

func (l *PwdLoader) normalizeHomeDir(path string) string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to load user home dir")
		return ""
	}

	return strings.ReplaceAll(path, "~", homedir)
}
