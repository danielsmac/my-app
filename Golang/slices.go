package main

import (
	"fmt"
	"slices"
)

func main()  {
	// slices sao tipificados unicamente pelos elementos que eles contém (não pelo números de elemntos). um slice nao iniciado é = a nil e tem tamanho 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// para criar um slice com tamanho non-zero, usa-se o builtin make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices podem ser copiados.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	// imprime slice e exclui posição s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// imprime slice e inclui a posição s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// podemos declarar e iniciar o slice com uma linha
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// o pacote slice contém várias funções úteis
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2")
	}

	//slices podem se composto dentro de uma estrutura de dados multi-dimensional
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}