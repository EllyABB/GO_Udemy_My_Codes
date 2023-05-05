package main

import "fmt"

const(
    y = 2023 + iota
    y1 = 2023 + iota
    
    y2 = 2023 + iota
    y3 = 2023 + iota
    y4 = 2023 + iota
    
)

func main(){

    
    fmt.Println(y,y1,y2,y3,y4)
}