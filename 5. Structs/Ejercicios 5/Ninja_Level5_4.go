package main

import "fmt"


func main() {
    p1 := struct{
            first string
            last string
            flavor []string
            }{
        "Albert",
        "Einstein",
        []string{"No_idea","No_idea_2"},
    }
    
    for _,v:=range p1.flavor {
          fmt.Println(v)
        }
     fmt.Println("----")

}