package main

import(
    "os"
    "log"
    "fmt"
)

func main(){

    defer foo() // Ahora si sale! corren las func defer!!
    
    _, err:=os.Open("nofile.txt")
    if err!=nil{
        //fmt.Println("Acá hay un error: ",err)
        //log.Println("log form: ",err) 
        //log.Fatalln("Puedo escribir algo aca?",err)

        log.Panicln("Puedo escribir algo aca?",err,err.Error()) // equivalente a un frt.println seguido de un panic

        // Poner el método .Error() es lo mismo que err

    }
    
    fmt.Print("No salí, veamos si las funciones diferidas corren")   
}


func foo(){
    fmt.Println("funcion diferida")
}