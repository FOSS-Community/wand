package wand

import (
	"fmt"
	"os/exec"
)

func Runcommand(c string, a ...string) {
	cmd := exec.Command(c, a...)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}
