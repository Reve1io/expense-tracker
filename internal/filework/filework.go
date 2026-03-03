package filework

import (
	"encoding/json"
	"expense-tracker/internal/model"
	"fmt"
	"os"
)

func ReadFile(filename string) []model.Expense {

	var expense []model.Expense

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	err = json.Unmarshal(data, &expense)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}

	return expense
}

func OpenFile(filename string) *os.File {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print("error:", err)
	}

	return file
}
