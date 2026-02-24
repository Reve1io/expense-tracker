package main

import (
	"fmt"
	"os"
	"strconv"

	"expense-tracker/internal/processes"
)

func main() {

	if len(os.Args) == 0 {
		fmt.Print("Not enough arguments!")
		return
	}

	if len(os.Args) <= 2 {
		fmt.Println("Not enough arguments! Example command: expense help")
		return
	}

	command := os.Args[2]
	category := os.Args[3]
	amount, _ := strconv.ParseFloat(os.Args[4], 64)
	description := os.Args[5]

	switch command {
	case "add":
		processes.Add(amount, category, description)
	case "list":
		fmt.Println("Command get list all expense")
	case "help":
		fmt.Println("Command helper")
	default:
		fmt.Println("Enter: expense help")
	}

}
