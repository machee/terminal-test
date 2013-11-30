package main

import (
	"fmt"
//	"github.com/kr/pty"
	"github.com/james4k/terminal"
	"os"
	"os/exec"
	"time"
)

func main() {
	//cmd := exec.Command("/usr/bin/dtach", "-A", "asdf", os.Getenv("SHELL"), "-i")
	cmd := exec.Command(os.Getenv("SHELL"), "-i")

	//shell, err := pty.Start(cmd)
	var state terminal.State
	_, shell, err := terminal.Start(&state, cmd)
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second)

	shell.Write([]byte("touch blah\nexit\n"))
	cmd.Wait()
}
