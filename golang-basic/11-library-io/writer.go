package main

import (
	"log"
	"os"
)

func main() {
	dst, err := os.Create("create.txt")
	if err != nil {
		log.Fatal("program broke")
	}
	defer dst.Close()

	bs := []byte("hello world, Santekno")

	_, err = dst.Write(bs)
	if err != nil {
		log.Fatal("program broke")
	}
}
