package main

import "fmt"

func main() {
	var langs []string
	
	langs = append(langs, "Go")
	langs = append(langs, "Ruby")
	langs = append(langs, "JavaScript")
	fmt.Println(langs)
}