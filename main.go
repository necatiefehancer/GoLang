package main

import (
	"fmt"
	"os"
)

func main() {

}

func letterUpperKeywords() {

	fmt.Println("hello planet")
	fmt.Println("www.heraSystem.com")

	// public private kapsamı kapısına çıkan isimlendirmeler

	var projectName string = "GoBackendServer" //dış dosyalardan buna erişilmesin private yapısı
	var ProjectName string = "GOBackendServer" // dış dosyalardan buna erişilsin public yapısı

	fmt.Println(projectName, ProjectName)
}

func areaVariables() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	var uName string = os.Getenv("USERNAME")
	var uDomain string = os.Getenv("USERDOMAIN")
	var processorArthitecture string = os.Getenv("PROCESSOR_ARCHITECTURE")
	var processorIdentifier string = os.Getenv("PROCESSOR_IDENTIFIER")
	var processorLevel string = os.Getenv("PROCESSOR_LEVEL")
	var goRoot string = os.Getenv("GOROOT")
	var goPath string = os.Getenv("GOPATH")
	var homePath string = os.Getenv("HOMEPATH")

	fmt.Println("username" + uName)
	fmt.Println("Domain" + uDomain)
	fmt.Println("İşlemci Mimarisi " + processorArthitecture)
	fmt.Println(processorIdentifier)
	fmt.Println("İşlemci seviyesi " + processorLevel)
	fmt.Println("GO ROOT ALANI" + goRoot)
	fmt.Println("GO PATH ALANI" + goPath)
	fmt.Println("HOME PATH ALANI" + homePath)
}

func Variables() {
	// const text string = "efe"

	// var (
	// 	degisken1 string = "efe"
	// 	degisken2 int    = 23
	// )

	// var name string
	// name = "necatiefe"
	// fmt.Println(name)

	// var a, b, c int
	// var a, b, c int = 3, 4, 5

	// fmt.Println(a, b, c)

	// var message = "HelloPlanet"
	// var a, b, c = 3, true, 4.8
	// fmt.Println(message, a, b, c)

	// var k, o string = "abc", "xyz"

	// var p = 42
	// var s, b = "xyz", true

	// n := 28
	// name, age := "necati efe", 20
	// fmt.Println(name, age)

	// a := "Text"
	// b := 'T'
	// c := `Text`

	// fmt.Println(a, b, c)

	// var myFloat32 float32 = 58.
	// myComplex := complex(5, 6)
	// println(myFloat32)
	// println(myComplex)

}

func Operators() {
	var (
		a     int = 10
		b     int = 20
		total int
	)
	total = a + b
	total = total - 5

	total *= 20

	total++

	total--
	fmt.Println(total)
}
