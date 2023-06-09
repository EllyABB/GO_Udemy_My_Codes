package main

import(
    "os"
    "log"
    "fmt"
)

func main(){

    defer foo() // No sale! no corren las func defer!!
    
    _, err:=os.Open("nofile.txt")
    if err!=nil{
        //fmt.Println("Acá hay un error: ",err)
        //log.Println("log form: ",err) 
        
        log.Fatalln("Puedo escribir algo aca?",err)
    }
    
    fmt.Print("No salí, veamos si las funciones diferidas corren")   
}


func foo(){
    fmt.Println("funcion diferida")
}