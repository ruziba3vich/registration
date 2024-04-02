package main

import "fmt"

func main() {
	// Ikkita maplarni birlashtiradigan dastur tuzing
	// Masalan: {"a": 1, "b": 2} | {"a": 2, "c": 3} => {"a": 2, "b": 2, "c": 3}
	m1 := map[string]int {"a": 1, "b": 2}
	m2 := map[string]int {"a": 2, "c": 3}

	for k, v := range m2 {
		m1[k] = v
	}

	fmt.Println(m1)
}
