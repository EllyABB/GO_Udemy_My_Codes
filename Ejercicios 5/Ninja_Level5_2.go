package main

import "fmt"

type person struct{
    first string
    last string
    flavor []string
}

func main() {
    p1 := person {
        "Albert",
        "Einstein",
        []string{"No_idea","No_idea_2"},
    }
    p2 := person {
        first: "Elly",
        last: "Bayona",
        flavor: []string{"Vainilla!","Chocolate"},
    }
   
    m := map[string]person{
        p1.last: p1,
        p2.last: p2,
    }
    
    
    for _,v:=range m {
        fmt.Println(v.first)
        fmt.Println(v.last)
        for i,v2:= range v.flavor {
            fmt.Println(i,v2)
        }
     fmt.Println("----")

    
    }
}
