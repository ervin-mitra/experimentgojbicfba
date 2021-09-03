package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())

	// variable types
	var i int = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	numb := 69
	fmt.Println((numb))

	b := true
	fmt.Println((b))

	c := complex(6, 9)
	fmt.Println((c))

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	// pointer data types
	var firstName *string = new(string)
	*firstName = "pointers"
	fmt.Println(*firstName)

	midName := "mid"

	ptr := &midName
	fmt.Println(ptr, *ptr)

	midName = "topson"
	fmt.Println(*ptr, ptr)

	// constants

}
