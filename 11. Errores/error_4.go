package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("nofile.txt")

	if err != nil {
		fmt.Println("Acá hay un error: ", err)
		log.Println("otra vez: ", err)
	}

	fmt.Println("Salí")
}
