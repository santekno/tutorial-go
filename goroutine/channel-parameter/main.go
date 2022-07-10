package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	var messages = make(chan string)

	for _, each := range []string{"america", "marvel", "panther"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
