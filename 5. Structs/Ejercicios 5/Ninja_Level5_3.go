package main

import "fmt"

type vehicle struct{
    doors int
    color string
}

type truck struct{
    vehicle
    fourWheel bool
}

type sedan struct{
    vehicle
    luxury bool
}

func main() {
    t1 := truck {
        vehicle: vehicle{
            doors: 2,
            color: "red",
        },
        fourWheel: true,
    }
   
    s1 := sedan {
        vehicle: vehicle{
            doors: 4,
            color: "blue",
        },
        luxury: true,
    }
   
    fmt.Println(t1,s1)
    fmt.Println(t1.color,s1.doors)
    
    

}