
package dog

import(
	"testing"
	"fmt"
)

func TestYears(t *testing.T){
	s:=Years(2)
	if s!=14{
		t.Error("obtenido",s,"want", 14)
	}
}


func TestYearsTwo(t *testing.T){
	s:=YearsTwo(2)
	if s!=14{
		t.Error("obtenido",s,"want", 14)
	}
}


func ExampleYear(){
	fmt.Println(Years(2))
	// Output:
	// 14
}

func ExampleYearTwo(){
	fmt.Println(YearsTwo(2))
	// Output:
	// 14
}


func BenchmarkYear(b *testing.B){
	b.ResetTimer()

	for i := 0; i<b.N; i++{
		Years(2)
	}
}

func BenchmarkYearTwo(b *testing.B){
	b.ResetTimer()

	for i := 0; i<b.N; i++{
		YearsTwo(2)
	}
}