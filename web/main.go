package main

import (
	"fmt"
	"mux-rest-api/infra"
)

func main() {

	fmt.Println("started application...")
	infra.LoadVars()

}
