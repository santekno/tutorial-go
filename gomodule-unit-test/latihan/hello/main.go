package main

import (
	"fmt"

	"rsc.io/quote"
)
func main()  {
	fmt.Println(HelloWorld())
}

func HelloWorld() string {
	return quote.Hello()
}