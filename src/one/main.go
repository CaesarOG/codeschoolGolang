package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

var (
	hour int
)

func main() {
	hour = time.Now().Hour()
	greeting, err := getGreeting(hour)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	j := 0 // no init, condition or post statement in for loop!
	for {
		if j >= 5 {
			break
		}
		fmt.Printf("%s %d\n", greeting, j)
		j++
	}
}

func getGreeting(hour int) (string, error) {
	var message string
	if hour < 7 {
		err := errors.New("Too early for greetings!")
		return message, err
	} else if hour < 12 {
		message = "good morning"
	} else if hour < 18 {
		message = "good afternoon"
	} else {
		message = "good evening"
	}
	return message, nil
}
