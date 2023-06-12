package main

import (
	"fmt"
	"GO_Udemy_My_Codes/13.Test/Exercises/two/quote"
	"GO_Udemy_My_Codes/13.Test/Exercises/two/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}