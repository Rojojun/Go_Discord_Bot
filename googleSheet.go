package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"honnef.co/go/tools/config"
)

func writeToSheet((data []interface{}) error {
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
        return fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "your-spreadsheet-id" // 스프레드시트 ID
    rangeData := "Sheet1!A1"    		   // 데이터 입력 범위

	vr := &sheets.ValueRange{
        Values: [][]interface{}{data},
    }

    _, err = srv.Spreadsheets.Values.Append(spreadsheetId, rangeData, vr).ValueInputOption("RAW").Do()
    if err != nil {
        return fmt.Errorf("unable to write data to sheet: %v", err)
    }

    return nil
}