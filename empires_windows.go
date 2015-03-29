package main

import "os/exec"

func LaunchEmpires() (err error) {

	if err = exec.Command("cmd", "/c", "START", "steam://rungameid/17740").Start(); err != nil {

		return
	}

	return
}

func LaunchEmpiresConnect(ip, password string) (err error) {

	if err = exec.Command("cmd", "/c", "START", "steam://connect/"+ip+"/"+password).Start(); err != nil {

		return
	}

	return
}
