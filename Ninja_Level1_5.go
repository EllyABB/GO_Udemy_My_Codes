package main

import "fmt"

type Vinny int

var x Vinny 
var y int

func main() {
    fmt.Println("The zero value of x=",x)
    fmt.Printf("The type of x = %T\n",x)
    x=55
    fmt.Println("\nThe new value of x=",x)
    fmt.Printf("We have the same type= %T.\n",x)
    
    
    y=int(x)
    fmt.Println("\nThe value of y=",y)
    fmt.Printf("The type of y= %T.\n",y)
    
    
    
}