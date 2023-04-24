package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true


func main() {
    s:=fmt.Sprintf("Tiene %v años, su nombre es %v. Esto es %v.\n",x,y,z)
    fmt.Println("s=",s)
}