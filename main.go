package main

import (
	"fmt"
	"os"
)

func main() {

	// fmt.Println("hello planet")
	// fmt.Println("www.heraSystem.com")

	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	var uName string = os.Getenv("USERNAME")
	var uDomain string = os.Getenv("USERDOMAIN")
	var processorArthitecture string = os.Getenv("PROCESSOR_ARCHITECTURE")
	var processorIdentifier string = os.Getenv("PROCESSOR_IDENTIFIER")
	var processorLevel string = os.Getenv("PROCESSOR_LEVEL")

	fmt.Println(uName)
	fmt.Println(uDomain)
	fmt.Println(processorArthitecture)
	fmt.Println(processorIdentifier)
	fmt.Println(processorLevel)
}
