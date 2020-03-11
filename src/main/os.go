package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	os.Hostname()
	fmt.Println(os.Hostname())
	fmt.Println(os.Environ())

	f, e := exec.LookPath("main")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(f)

}
