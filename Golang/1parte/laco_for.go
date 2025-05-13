package main

import ("fmt")

func main()  {

	// tipo básico com uma condição simples
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// clássico inicial/condicional/depois do loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// outro mode de faser o tipo básico (do tipo N vezes), usando range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// laço for sem condicional repetindo um loop até o break ou fechamento da função
	for {
		fmt.Println("loop")
		break
	}

	// pode-se também continuar até a próxima iteração do loop
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	
}