package processes

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"expense-tracker/internal/model"
)

func Add(amount float64, category string, description string) {

	timeExpense := time.Now()
	formattedDate := timeExpense.Format("2006-01-02 15:04:05")

	expense := model.Expense{
		ID:       1,
		Date:     formattedDate,
		Category: category,
		Amount:   amount,
		Note:     description,
	}

	fmt.Print("Added new expense:", expense)

	file, err := os.OpenFile("expenses.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print("error:", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(expense)
	if err != nil {
		fmt.Print("error:", err)
	}
}
