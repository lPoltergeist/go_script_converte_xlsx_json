package main

import (
	"encoding/json"
	"os"

	"github.com/thedatashed/xlsxreader"
)

var (
	headers []string
	result  []map[string]any
	i       int
)

func main() {
	xl, err := xlsxreader.OpenFile("./nome do arquivo.xlsx")
	if err != nil {
		panic(err)
	}

	defer xl.Close()

	headers := map[string]string{}
	result := []map[string]string{}

	rowIndex := 0
	for row := range xl.ReadRows(xl.Sheets[0]) {

		if rowIndex == 0 {
			for _, cell := range row.Cells {
				headers[cell.Column] = cell.Value
			}
			rowIndex++
			continue
		}

		rowData := map[string]string{}
		for _, cell := range row.Cells {
			if header, ok := headers[cell.Column]; ok && header != "" {
				rowData[header] = cell.Value
			}
		}
		result = append(result, rowData)
		rowIndex++

	}
	file, err := os.Create("./nome de saida.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(result); err != nil {
		panic(err)
	}

}
