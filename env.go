package main

import (
	"fmt"
	"os"
)

func main() {
	environ := os.Environ()
	for i := range environ {
		fmt.Println(environ[i])
	}
	fmt.Println("------------------------------------------------------------\n")
	logname := os.Getenv("LOGNAME")
	fmt.Printf("logname is %s\n", logname)
}
