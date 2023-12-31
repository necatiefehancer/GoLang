package functions

import "fmt"

func Add(x int, y int) (int, error) {

	return x + y, nil

}

func SayHello1() {
	fmt.Println("HELLO")
}

func SayHello2(message string) {
	fmt.Printf("HELLO %v\n", message)
}

func Change1(x int) {
	fmt.Println(x)
	x = 96
}

func Change1ByReference(x *int) {
	fmt.Println(*x)
	*x = 96
}

func Swap(m1, m2 string) (string, string) {
	return m2, m1
}

func VariadicSum(terms ...int) (int, int) {

	///slice alırlar

	var sum int = 0

	for _, term := range terms {
		sum += term
	}

	return len(terms), sum

}

func NamedVariableFunctions(sum int) (x, y int) {
	x = sum * sum
	y = sum * sum * sum
	return
}

func NamedVariableFunctionsVariadic(terms ...int) (sum int, lenterms int) {

	for _, v := range terms {
		sum += v
	}

	lenterms = len(terms)

	return

}
