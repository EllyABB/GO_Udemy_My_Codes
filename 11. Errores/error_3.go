package main 

import (
	"fmt"
	"os"
	"io/iountil"
)

// Tengo que solucionar el problema con io/until

func main(){

	f, err :=	os.Open("names.txt")


	if err != nil{
		fmt.Println(err)
		return
	}

	defer f.Close()


	bs, err := iountil.ReadAll(f)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(strings(bs))


	
}