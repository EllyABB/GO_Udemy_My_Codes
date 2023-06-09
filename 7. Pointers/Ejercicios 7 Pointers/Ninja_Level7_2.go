package main
import "fmt"


type person struct{
    name string
}

func changeMe(p *person, other_name string){
    p.name=other_name
    
    //to dereference a struct, use (*value).field
    // (*p).name = other_name
}


func main(){
    p1 := person{"Elly"}
    fmt.Println(p1.name)
    
    // Sin los apuntadores esto no funcionaría!! 
    changeMe(&p1,"Anne")
    fmt.Println(p1.name)   
}