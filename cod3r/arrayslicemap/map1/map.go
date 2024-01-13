package main

import "fmt"

func main() {
	// var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[69697697697] = "Pedro"
	aprovados[84384834834] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[84384834834])
	delete(aprovados, 84384834834)
	fmt.Println(aprovados[84384834834])
}
