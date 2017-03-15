package main

import (
	"fmt"
	"github.com/aofry/go-tee/proxy"
)

func main() {
	fmt.Println("Starting")

	proxy.New()

	fmt.Println("Ending")
}
