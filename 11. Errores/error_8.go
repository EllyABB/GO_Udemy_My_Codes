package main

import(
    "os"
    
    "fmt"
)

func main(){

    defer foo() // Ahora si sale! corren las func defer!!
    
    _, err:=os.Open("nofile.txt")
    if err!=nil{
        //fmt.Println("Acá hay un error: ",err)
        //log.Println("log form: ",err) 
        //log.Fatalln("Puedo escribir algo aca?",err)

        //log.Panicln("Puedo escribir algo aca?",err,err.Error()) 

        //panic("y aca?",err) // En estos casos no puedo mandar un mensaje!
        panic(err) 
    }
    
    fmt.Print("No salí, veamos si las funciones diferidas corren")   
}


func foo(){
    fmt.Println("funcion diferida")
}