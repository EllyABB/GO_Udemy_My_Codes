package main 

import "fmt"

func main(){

	var var1, var2, var3 string

	fmt.Println("Escribe tu nombre:")
	_,err := fmt.Scan(&var1)

	if err != nil{
		fmt.Println("entré")
		panic(err)
	}

	fmt.Println("Deporte fav: ")
	_,err = fmt.Scan(&var2)      // aqui solo re asignamos la variable err

	if err != nil{
		fmt.Println("entré")
		panic(err)
	}
	

	fmt.Println("Comida fav: ")
	_,err = fmt.Scan(&var3)

	if err != nil{
		fmt.Println("entré")
		panic(err)
	}


	fmt.Println("salí con: ", var1, var2, var3)
}