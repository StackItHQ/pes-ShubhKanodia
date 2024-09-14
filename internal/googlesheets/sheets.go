package googlesheets

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// wrapper for sheet service
type SheetService struct {
	service *sheets.Service
}

//schema is {name, age, dept, sem}

func NewSheetService(ctx context.Context, credentialsFile string) (*SheetService, error) {
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	return &SheetService{service: srv}, nil
}

func (s *SheetService) WriteToSheet(spreadsheetID, range_ string, values [][]interface{}) error {
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := s.service.Spreadsheets.Values.Update(spreadsheetID, range_, valueRange).
		ValueInputOption("RAW").Do()

	if err != nil {
		return fmt.Errorf("unable to write data to sheet: %v", err)
	}
	fmt.Print("Written heheheha!")
	return nil
}

// to write to sheet
func (s *SheetService) WriteStudentData(spreadsheetID string, students [][]string) error {
	values := make([][]interface{}, len(students))
	for i, student := range students {
		values[i] = make([]interface{}, len(student))
		for j, value := range student {
			values[i][j] = value
		}
	}

	return s.WriteToSheet(spreadsheetID, "Sheet1!A1", values)
}
