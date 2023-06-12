	package main

	import(
		"testing"
	)
	func TestMySum (t *testing.T){       // It stars now with capital letter. The name is whatever, but this name is a good practice
		x:=mySum(2,3)
		if x!=5{
			t.Error("Problemas! Esperado", 5, " obtenido: ", x)
		}

	}  