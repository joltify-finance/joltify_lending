package common

import (
	"fmt"

	excelize "github.com/xuri/excelize/v2"
)

func WritePoolToExcel(sheetName string, data [][]string, fileName string) {
	var f *excelize.File
	var err error
	f, err = excelize.OpenFile(fileName)
	if err != nil {
		// Create a new sheet.
		f = excelize.NewFile()
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	_, err = f.NewSheet(sheetName)
	if err != nil {
		fmt.Printf("shee already exist %v\n", err)
	}
	for idx, rowData := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.SetSheetRow(sheetName, cell, rowData)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// Set value of a cell.
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}
}
