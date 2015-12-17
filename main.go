package main

import (
	"fmt"
	"github.com/dlsniper/test/mymodule"
	"time"
)

func main() {
	fmt.Printf("BEGINNING\n")

	sleep := 1 * time.Second; _ = sleep

	toto := 42

	mymodule.Mdr()

	Haha()

	fmt.Printf("HELLO %d\n", toto)

	//time.Sleep(sleep)
}
