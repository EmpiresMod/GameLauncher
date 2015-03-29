package main

import (
	"fmt"
	"os/exec"
)

func LaunchEmpires() (err error) {

	if err = exec.Command("cmd", "/c", "START", "steam://rungameid/17740").Start(); err != nil {

		return
	}

	return
}

func LaunchEmpiresConnect(ip, password string) (err error) {

	ipPassword := fmt.Sprintf("%s/%s", ip, password)
	if err = exec.Command("cmd", "/c", "START", "steam://connect/"+ipPassword).Start(); err != nil {

		return
	}

	return
}
