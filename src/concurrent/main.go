package main

import (
	"fmt"
	"sync"
	"math"
)

func main() {
	names := []string{"Phil", "Noodles", "Barbaro"}
	wg := &sync.WaitGroup{}//could do var wg sync.WaitGroup (see AltWay2)
	//var wg *sync.WaitGroup = &sync.WaitGroup{}//also works
	wg.Add(len(names))//when it's a point to a struct, can do wg or *wg
	for _, name := range names {
		go printName(name, wg)
	}//if run sequential each call blocks processor for almost 5 seconds
	wg.Wait()//make it wait for all goroutines to complete
}
//can track run time by doing "time go run main.go" on Unix systems
func printName(n string, wg *sync.WaitGroup) {
	result := 0.0
	for i := 0; i < 100000000; i++ {
		result += math.Pi * math.Sin(float64(len(n)))
	}
	fmt.Println("Name: ", n)//prints arg to console
	wg.Done()
}
//go programs must be made aware of created goroutines or will finish b4



//AltWay1 to do main method
//	names := []string{"Phil", "Noodles", "Barbaro"}
//	w := sync.WaitGroup{}
//	wg := &w
//	wg.Add(len(names))
//	for _, name := range names {
//		goprintName(name, wg)
//	}//if run sequential each call blocks processor for almost 5 seconds
//	wg.Wait()//make it wait for all goroutines to complete

//AltWay2 to do main method
//	names := []string{"Phil", "Noodles", "Barbaro"}
//	var wg sync.WaitGroup
//	wg.Add(len(names))
//	for _, name := range names {
//		go printName(name, wg)
//	}//if run sequential each call blocks processor for almost 5 seconds
//	wg.Wait()//make it wait for all goroutines to complete