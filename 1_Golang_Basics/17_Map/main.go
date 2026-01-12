package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"Sush": 65,
		"Sai":  35,
	}

	fmt.Println(ages["sangam"], len(ages))

	// make(map[K]V)

	var scores map[string]int // nil map

	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)

	scores["math"] = 90

	fmt.Println(scores, scores["math"])

	users := map[string]string{
		"u1": "Sush",
		"u2": "Sai",
		"u3": "Rahul",
	}

	fmt.Println(users)

	delete(users, "u2")
	delete(users, "u100") // no error

	fmt.Println(users)

}
