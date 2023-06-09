package main

import (
    "fmt"
)

func main() {
    m := map[string][]string{
        "James":           []string{`Shaken, not stirred`, `Martinis`, `Women`},
        "Miss Moneypenny": []string{`James Bond`, `Literature`, `Computer Science`},
        "no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
        
    }
    
    m["Elly"]=[]string{"Patinar","Familia","Estudiar"}
    
    for k,v := range m{
        fmt.Println(k,v)
    }
    
    delete(m,"Elly")
    
    
    fmt.Println("
    for k,v := range m{
        fmt.Println(k,v)
    }
    
}