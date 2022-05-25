package main

import (
	"fmt"

	controllers "github.com/crosstech/xblocks-golang/example/controllers/_controllers"
)

func main() {

	fmt.Println("Waiting for requests...")

	controllers.Run()
}
