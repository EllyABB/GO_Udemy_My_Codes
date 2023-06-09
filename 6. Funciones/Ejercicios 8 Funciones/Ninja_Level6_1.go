package main

import "fmt"

type retorno struct{
    entero int
    letra string
}
func main() {
    fmt.Println(foo())
    fmt.Println(bar())
    
    z,y:=bar2()
    fmt.Println(bar2())
    fmt.Println(z,y)
}

func foo() int{
    return 4
}

func bar() retorno {
    a:= retorno{
        entero: 2,
        letra: "a",
    }
    return a

}

func bar2() (int,string) {
    return 3,"a"
}