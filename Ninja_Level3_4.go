package main

import "fmt"

func main(){
    edad := 28
    year :=2023
    i:=0
    for {
        fmt.Println(year-i)
        i++
        if i>edad+1{
            break
        }
    }
}