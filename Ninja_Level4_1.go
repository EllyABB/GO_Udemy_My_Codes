package main

import "fmt"


func main() {
    x := [5]int{1,1,1,1,1}
    
    for i,v := range x{
        fmt.Println(i,v)
        fmt.Printf("Tipo = %T\n",v)
    }
    fmt.Printf("Tipo = %T\n",x)

    
  
}