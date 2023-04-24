package main

import "fmt"


var x int = 34

func main(){
   
    fmt.Printf("Formato decimal = %d. Formato binario = %b. Formato Hex = %#X\n",x,x,x)
    
    y:= (x<<1)
    fmt.Println("\nDespues del shift:")
    fmt.Printf("Formato decimal = %d. Formato binario = %b. Formato Hex = %#X\n",y,y,y)
    
    
}