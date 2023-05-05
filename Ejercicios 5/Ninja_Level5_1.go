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
   
    fmt.Println(p1,p2)
    fmt.Println(p1.first,p1.last,p1.flavor)
    fmt.Println(p2.first,p2.last,p2.flavor)
    fmt.Printf("%T,\t%T\n",p1,p2)
    
    
    for i,v:=range p1.flavor {
        fmt.Println(i,v)
    
    }
}
