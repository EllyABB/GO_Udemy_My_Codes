package main

import "fmt"

func main(){

    for i:=1;i<30;i++{
        if i%4==0{
             fmt.Printf("%d es divisible por 4  \n",i)
        } else if i%3==0{
             fmt.Printf("%d es divisible por 3  \n",i)
        } else {
             fmt.Printf("%d no es divisible ni por 4 ni por 3 \n",i)
        }
    }
}