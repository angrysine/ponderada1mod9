package main

import (
	"os/exec"
)

func main()  {

	go broker()
	go Pub()
	Sub()
	 

	
}

func broker() {
	cmd := exec.Command("mosquitto","-c","mosquitto.conf")
	cmd.Output()

	
}

