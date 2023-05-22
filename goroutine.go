package main

import (
    "fmt"
    "runtime"
    "sync"
)


func main(){
    
    fmt.Println("NumCPU",runtime.NumCPU())
    fmt.Println("Goroutines",runtime.NumGoroutine())
    counter :=0
    
    const gs = 100
    var wt sync.WaitGroup   
    wt.Add(gs)      
    
    var mu sync.Mutex //// Acá solucionamos el lío

    
    
    for i := 0 ;  i < gs;  i++ {
        go func(){
            mu.Lock()   // Bloqueado
            v:= counter    
            
            runtime.Gosched() 
            v++
            counter =v
            
            mu.Unlock() // DEsbloqueado
            wt.Done()  
            
        }()
        //fmt.Println("Counter",counter)    
        fmt.Println("Goroutines",runtime.NumGoroutine()) 

    }
    
    wt.Wait() 
    fmt.Println("Goroutines",runtime.NumGoroutine()) 
    fmt.Println("Counter",counter)
    
}