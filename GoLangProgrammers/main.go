package main

import (
	"fmt"
)

const currentlyWorking = ""

var workOut = map[string]func(){
	"Variables":    variables,
	"Constants":    constants,
	"TypeCasting":  typeCasting,
	"Operators":    operators,
	"Conditionals": conditionals,
	"Loops":        loops,
}

func main() {
	if len(currentlyWorking) != 0 {
		fmt.Println("\n\tWorking with ", currentlyWorking)
		workOut[currentlyWorking]()
	} else {
		for index, value := range workOut {
			fmt.Println("\n\tWorking with ", index)
			value()
		}
	}
}
