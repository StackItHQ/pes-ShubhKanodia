package main

import (
	"context"
	"fmt"
	"log"

	"github.com/StackItHQ/pes-ShubhKanodia/internal/googlesheets"
)

func main() {
	ctx := context.Background()
	sheetService, err := googlesheets.NewSheetService(ctx, "credentials.json")
	if err != nil {
		log.Fatalf("Unable to create sheet service :( %v", err)
	}

	spreadsheetID := "1oD5RkhAjiqjeuvF-Lr1E0va5t0XXeOxNSo0T8OeOIrg"

	students := [][]string{
		{"Name", "Age", "Dept", "Sem"},
		{"A", "20", "Computer Science", "3"},
		{"B", "22", "Electrical Engineering", "5"},
		{"C", "21", "Mechanical Engineering", "4"},
	}

	err = sheetService.WriteStudentData(spreadsheetID, students)
	if err != nil {
		log.Fatalf("Unable to write to sheet: %v", err)
	}

	fmt.Println("Data successfully written to sheet heheheha!")
}
