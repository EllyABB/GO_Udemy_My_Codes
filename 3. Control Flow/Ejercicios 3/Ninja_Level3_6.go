package main

import "fmt"

func main(){

    for i:=10;i<100;i++{
        if i%4==0{
             fmt.Println("i es divisible por 4 = ",i)
        }
    }
}