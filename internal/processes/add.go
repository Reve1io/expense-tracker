package processes

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"expense-tracker/internal/filework"
	"expense-tracker/internal/model"
)

func Add(amount float64, category string, description string) {

	timeExpense := time.Now()
	formattedDate := timeExpense.Format("2006-01-02 15:04:05")

	expense := model.Expense{
		ID:       2,
		Date:     formattedDate,
		Category: category,
		Amount:   amount,
		Note:     description,
	}

	expenses := []model.Expense{}

	filename := "expenses.json"

	expenseJson := filework.ReadFile(filename)

	if expenseJson != nil {
		fmt.Println("the file exists, but it is empty")
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Print("error:", err)
	}

	defer file.Close()

	expenses = append(expenses, expense)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(expenses)
	if err != nil {
		fmt.Print("error:", err)
	}
}
