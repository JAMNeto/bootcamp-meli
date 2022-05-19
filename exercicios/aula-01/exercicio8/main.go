package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	counter := 0
	fmt.Printf("A idade de Benjamin é %d anos\n", employees["Benjamin"])
	for _, val := range employees {
		if val > 21 {
			counter += 1
		}
	}
	fmt.Printf("A empresa possui %d funcionários com mais de 21 anos\n", counter)
	employees["Frederico"] = 25
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println(employees)
}
