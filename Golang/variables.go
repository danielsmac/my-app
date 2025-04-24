package main

import ("fmt")

func main()  {
	
	// variável do tipo string
	var a = "initial"
	fmt.Println(a)

	// pode-se declarar várias variáveis ao mesmo tempo (tipo int)
	var b, c int = 1, 2
	fmt.Println(b, c)

	//pode-se inferir o tipo de inicializador na variável
	var d = true
	fmt.Println(d)

	// Variáveis sem um valor inicial é considerada zero-valued.
	var e int
	fmt.Println(e)

	// a sintax := é um facilitador para declarar e inicializar variáveis (e.g. var f string = "apple")
	f := "apple"
	fmt.Println(f)
}