package main

import "fmt"

func zeroval(ival int)  {
	ival = 0
}

// * representa um ponteiro
func zeroptr(iptr *int)  {
	*iptr = 0
}

func main()  {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// sintax &i libera endereço na memória, no caso pro ponteiro
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}