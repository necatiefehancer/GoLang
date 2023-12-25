package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	StatementBreakContiuneGoto()
}

func StatementBreakContiuneGoto() {
	///break

	// bulunduğu scope kırar ordan çıkarır programı
	var i int = 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("i values of :", i)
		i++
	}
	fmt.Println("Break döngüyü bulunduğu scope`ı kırdı ve çıktı")

}

func SwitchStatementControl() {

	var a int = 13
	///case sağlasın sağlamasın fallthrough olursa alltaki case de çalıştırılır
	switch a {
	case 9:
		fmt.Println("numeric ==9")
		// fallthrough
	case 10:
		fmt.Println("numeric ==10")
	default:
		fmt.Println("not declared types")
	}

	var score int
	fmt.Println("Please log a score ")
	fmt.Scanf("%v", &score)
	fmt.Printf("%v", score)
	switch {
	case score <= 59:
		fmt.Println("Your Notes F")
	case score <= 69:
		fmt.Println("Your Notes C")
	case score <= 79:
		fmt.Println("Your Notes D")
	case score <= 89:
		fmt.Println("Your Notes B")
	case score <= 100:
		fmt.Println("Your notes A")
	default:
		fmt.Println("Not defined value")

	}

}

func İfStateMentControl() {

	var (
		a   int = 10
		b   int = 10
		foo int = 1
	)

	if a > b {
		fmt.Println("a>b")
	} else if a == b {
		fmt.Println("a=b")
	} else {
		fmt.Println("b>a")
	}

	if foo == 1 {
		fmt.Println("bar")
	}

	//if içinde initilaze kurulum yapıp sorgulama

	if efe := 9; efe == 3 {
		fmt.Println("efe ==3")
	} else {
		fmt.Println("efe!=3")
	}
	//   fmt.Println(efe) hata verir o değişken if içinde tanımlandı scope kapsam hatası
}

func DateAndTimePackages() {

	///istenilen tarihi time tipinde oluşturma
	t := time.Date(2004, time.September, 6, 22, 18, 50, 12, time.UTC)
	fmt.Printf("%T", t)
	fmt.Printf("%s", t)

	///şu anki saat tarih ve zamanı now methodu ile çekme
	//var ile değişken tipi bildirimi ile time tanımlama
	var now time.Time = time.Now()
	fmt.Println(now)

	// time nesnesi içinden istenilen verileri çekme

	fmt.Println(t.Day())
	fmt.Println(t.Weekday())
	fmt.Println(t.Month())
	fmt.Println(t.Year())
	fmt.Println(t.Location())

	///tarihi değiştirip yeni tarih dönderme

	var newTime time.Time = t.AddDate(1, 1, 1)
	fmt.Println(newTime)

	//formatlı gösterim
	// senin tarihini istediğin format şeklinde layouıt alarak geri döndürür
	var longFormat string = "Monday , september 6 , 2004"
	fmt.Println(newTime.Format(longFormat))

	var shortFormat string = "1/2/06"
	fmt.Println(t.Format(shortFormat))
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
