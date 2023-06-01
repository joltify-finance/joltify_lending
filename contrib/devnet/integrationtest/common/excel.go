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
	for idx := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Printf(">>>>set coordinates cell name error %v\n", err)
			return
		}
		err = f.SetSheetRow(sheetName, cell, &data[idx])
		if err != nil {
			fmt.Printf(">>>>set sheet row error %v\n", err)
			return
		}
	}
	// Set value of a cell.
	// Save spreadsheet by the given path.
	if err := f.SaveAs(fileName); err != nil {
		fmt.Printf(">>>>save excel file error %v\n", err)
	}
}
