package main

import "fmt"

func main(){
    a:=24
    b:=32
    
    x1:=(a==b)
    x2:=(a<=b)
    x3:=(a>=b)
    x4:=(a!=b)
    x5:=(a<b)
    x6:=(a>b)
    
    fmt.Println("Son iguales?=",x1)
    fmt.Println("Menor o igual?=",x2)
    fmt.Println("Mayor o igual?=",x3)
    fmt.Println("Son diferentes?=",x4)
    fmt.Println("Menor?=",x5)
    fmt.Println("Mayor?=",x6)

}