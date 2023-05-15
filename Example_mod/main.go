package main

import(
	"github.com/EllyABB/Example_package"
	"fmt"
)

func main(){
	s  := puppy.Ladrar()
	s1 := puppy.Ladridos()

	fmt.Println(s)
	fmt.Println(s1)

	fmt.Println(puppy.La_que_llama_a_otra())
}