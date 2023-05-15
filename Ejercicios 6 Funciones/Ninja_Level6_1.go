package main

import "fmt"


func main() {
    x := []int{1,1,1,1,1}
    
    
    fmt.Println(foo(x...))    
  
}


func foo (i ...int){
    sum :=0
    for _,v := range i{
        sum+=v
    }
}