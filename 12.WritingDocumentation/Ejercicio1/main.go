
package main

import (
    "fmt"
    "GO_Udemy_My_Codes/12.WritingDocumentation/Ejercicio1/dog"
)

// Be careful!!!!!
// go mod init entra en conflicto con el de afuera.
// Lo inicialice con:
// > go mod init GO_Udemy_My_Codes/12.WritingDocumentation/Ejercicio1/dog

func main(){
    a:=dog.Year(3)
    fmt.Println(a)
} 
