package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Languages string

const (
	JAVASCRİPT Languages = "Javascript D"
	C          Languages = "C"
	PYTHON     Languages = "Python"
	JAVA       Languages = "Java"
)

func PrintLanguage(lng Languages) {
	fmt.Println(lng)
}

func main() {
	ConsoleInputs()
}

func ConsoleInputs() {

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println("Your Values ", str)

	fmt.Print("Enter number")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("your values %v", f)
	}

}

func ConsoleOutputs() {
	var (
		text              string = "Necati Efe Hancer"
		stringLength, err        = fmt.Println(text)
		aNumber           int    = 28
		isBool            bool   = true
	)
	fmt.Println(stringLength, err)
	fmt.Printf("Value of number %v \n", aNumber)
	fmt.Printf("Value of boolean %v\n", isBool)
	fmt.Printf("Value of converted float %.2f\n", float64(aNumber))
	fmt.Printf("Data Types %T %T %T %T %T %T", text, stringLength, err, aNumber, isBool, 15.8)
}

func convertionVariables() {

	var (
		deger1 string  = "8"
		deger2 int     = 10
		deger3 float64 = 2.8
	)

	convertedNumber1, err := strconv.Atoi(deger1)

	fmt.Println(deger1, deger2, deger3)

	fmt.Println(convertedNumber1, err)

	var convertedNumber2 string = strconv.Itoa(deger2)

	fmt.Println(convertedNumber2 + "sivas")

	var i int = 80
	var j float64 = float64(i)
	var k string = "75"
	fmt.Println(i, j, k)
}

func ConstantState() {
	PrintLanguage(JAVASCRİPT)
	PrintLanguage(JAVA)
	PrintLanguage(PYTHON)
	PrintLanguage(C)
}

func BlcakIdentifier() {
	var _, x, _ int = 0, 9, 8
	fmt.Println(x)
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
