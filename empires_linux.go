package main

import "os/exec"

func LaunchEmpires() (err error) {

	if err = exec.Command("steam", "-applaunch", "17740").Start(); err != nil {

		return
	}

	return
}
