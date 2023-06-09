package main

import "fmt"

type Vinny int

var x Vinny 


func main() {
    fmt.Println(x)
    fmt.Printf("%T\n",x)
    x=42
    fmt.Println(x)
    
}