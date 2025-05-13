package main

import (
	"fmt"
	"time"
)

func main()  {
	
	// switch básico
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Pode-se usar várias expressões separadas por vírgula no mesmo case. Pode-se usar como opcional default case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("É fim de semana")
	default:
		fmt.Println("É dia de semana")
	}

	// switch sem uma expressão é um meio alternativo para expressões do tipo if/else. case pode ser non-constant
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Antes do meio-dia")
	default:
		fmt.Println("Depois do meio-dia")
	}

	// Switch pode comparar tipo em vez de valores. pode-se usar isso para desobrir qual tipo do valor da interface.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("É booleano")
		case int:
			fmt.Println("É um inteiro")
		default:
			fmt.Println("Não sabe-se o tipo %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}