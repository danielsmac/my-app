package main

import ( 
		"fmt"
)

func main()  {
	
	//exemplo básico
	if 7%2 == 0 {
		fmt.Println("7 é par")
	} else {
		fmt.Println("7 é ímpar")
	}

	//statement if sem o else
	if 8%4 == 0 {
		fmt.Println("8 é divisível por 4")
	}

	// usando operadores lógicos and e or
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 ou 7 são pares")
	}

	// statement pode preceder condicionais; qualquer variável declarada neste statement estará disponível na atual e subsequentes ramos
	if num := -1; num < 0 {
		fmt.Println(num, "é negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 dígito")
	} else {
		fmt.Println(num, "tem múltiplos dígitos")
	}
}