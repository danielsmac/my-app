package main

import (
		"fmt"
		"maps"
)

func main()  {
	
	// para criar um map vazio, usa o builtin make: make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	// pegar um valor de uma chave com nome [key]

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// se a chave não existe, o valor zero é retornado como valor type
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	//para remover key/value
	delete(m, "k2")
	fmt.Println("map:", m)

	// para remover todas key/value
	clear(m)
	fmt.Println("map;", m)

	// _ (identificador de valor vazio)
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	// pode-se iniciar um novo map na mesma linha com esta sintax
	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map", n)

	// o pacote map possui um número de utilitário úteis para maps

	n2 := map[string]int{"foo":1, "bar":2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}