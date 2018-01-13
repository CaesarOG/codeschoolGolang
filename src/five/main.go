package main

import (
	"fmt"
	"five/model"
)

func main() {
	jumperList :=model.GetList()
	for _, jumper := range jumperList {
		fmt.Println(jumper.Jump())
	}
}