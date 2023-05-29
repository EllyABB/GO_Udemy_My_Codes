package main

import "fmt"

func main(){
	
	c:=make(chan int)
	
	c<-42 
	c<-43
	
	fmt.Println(c)
	fmt.Printf("%T",c)

}
