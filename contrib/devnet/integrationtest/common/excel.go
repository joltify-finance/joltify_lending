package common

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func WriteDataToExcel(sheetName string, data [][]string, fileName string) {
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

	_, err = f.GetSheetIndex("Sheet1")
	if err == nil {
		err := f.DeleteSheet("Sheet1")
		if err != nil {
			fmt.Printf(">>>>delete sheet error %v\n", err)
		}
	}

	_, err = f.NewSheet(sheetName)
	if err != nil {
		fmt.Printf("shee already exist %v\n", err)
	}

	switch sheetName {
	case "pool_info":
		f.SetColWidth(sheetName, "A", "A", 20)
		f.SetColWidth(sheetName, "B", "B", 80)
		f.SetColWidth(sheetName, "C", "D", 30)
		f.SetColWidth(sheetName, "E", "G", 10)
		f.SetColWidth(sheetName, "H", "H", 30)
	case "depositor_info":
		f.SetColWidth(sheetName, "A", "A", 50)
		f.SetColWidth(sheetName, "B", "C", 30)
		f.SetColWidth(sheetName, "D", "D", 10)
		f.SetColWidth(sheetName, "E", "G", 30)

	case "borrow_info":
		f.SetColWidth(sheetName, "A", "A", 50)
		f.SetColWidth(sheetName, "B", "C", 30)
		f.SetColWidth(sheetName, "D", "E", 10)
		f.SetColWidth(sheetName, "F", "H", 30)
	default:
		panic("unspported sheet name")
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
