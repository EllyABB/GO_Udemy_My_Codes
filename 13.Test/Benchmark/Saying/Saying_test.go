package Saying

import(
	"testing"
	"fmt"
)

func TestGreet(t *testing.T){
	s:= Greet("Jammes")

	if s != "Welcome my dear Jammes"{
		t.Error("got", s,"expected","Welcome my dear Jammes")
	}
}

func ExampleGreet(){
	fmt.Println(Greet("Jammes"))
	//Output
	//Welcome my dear Jammes
}


// And the way benchmarking works is it's going to run this code 
// a lot of times until it gets a statistically
// accurate measurement of how long it took to run it.


func BenchmarkGreet(b *testing.B){
	for i :=0 ; i <b.N; i++{	//It'll determine how many times it runs.
		Greet("Jammes")
	} 
}