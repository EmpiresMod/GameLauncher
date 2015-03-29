package main

import "os/exec"

func LaunchEmpires() (err error) {

	if err = exec.Command("steam", "-applaunch", "17740").Start(); err != nil {

		return
	}

	return
}

func LaunchEmpiresConnect(ip, password string) (err error) {

	if err = exec.Command("steam", "-applaunch", "17740", "+connect", ip, "password", password).Start(); err != nil {

		return
	}

	return
}
