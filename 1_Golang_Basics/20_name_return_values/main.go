package main

import "fmt"

func divide(a int, b int) (sush int, amar int) {
	sush = a / b
	amar = a + b

	return
}

func main() {

	q, r := divide(10, 10)
	fmt.Println(q, r)
}
