
package main


import(
	"fmt"
	"GO_Udemy_My_Codes/13.Test/Examples_Doc/acdc"
)
//Para que funcione inicializo el mod de la forma:
//go mod init GO_Udemy_My_Codes/13.Test/Examples_Doc

func main(){
	fmt.Println(acdc.Sum(2,3))	
	fmt.Println(acdc.Sum(2,3,4))	
}

