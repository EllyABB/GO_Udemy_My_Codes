
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")   
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)		// It seems like everything is put on f because this line
							// without this one, everything works like fmt.Println

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("err happened", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
	defer f2.Close()


	log.Println("check the log.txt file in the directory")  // this is also put on the txt
}