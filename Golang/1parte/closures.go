package main

import "fmt"

//esta função retorna outra funcão que foi definida como anonima no corpo da função intSeq.
func intSeq() func() int  {
	
	i := 0
	return func() int {
		i++
		return i
	}
}

func main()  {
	
	//vamos chamar intSeq assinando como resultado (uma função) para nextInt. a função captura o próprio valor de i, que será atualizada cada vez que chamamos nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// para confirmar o estado unico desta particular função, criamos outra variável.
	newInts := intSeq()
	fmt.Println(newInts())
}