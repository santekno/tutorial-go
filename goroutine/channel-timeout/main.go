package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		randoSec := rand.Int()%10 + 1
		fmt.Println("lama proses : ", randoSec)
		time.Sleep(time.Duration(randoSec) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 seconds")
			break loop
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go sendData(messages)
	retreiveData(messages)
}
