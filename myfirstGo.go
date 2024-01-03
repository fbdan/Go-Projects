// 03.01.2024
// GO Language
// Bogdan F

package main

import (
	"fmt"
)

func main() {
	// Input numele necesar
	fmt.Print("Introduceti numele: ")

	// Citeste input de la user
	var name string
	fmt.Scanln(&name)

	// Bun venit
	greeting := fmt.Sprintf("Salut, %s! Acesta este primul tau program in Go !", name)
	fmt.Println(greeting)
}
