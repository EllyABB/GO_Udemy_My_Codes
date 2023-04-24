package main

import "fmt"

func main(){
    favSport:="Roller"
    switch favSport{
        case "Futbol","Futbol Americano":
            fmt.Println("Tu deporte favorito usa un balón")
        case "Roller", "Bicicleta":
            fmt.Println("Tu deporte favorito usa ruedas!")
        
        default:
            fmt.Println("No sé nada de tu deporte")
        }
    
}