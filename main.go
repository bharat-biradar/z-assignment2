package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type GenericTable struct {
	name    string
	columns table.Row
	rows    []table.Row
}

func main() {

	file_path := flag.String("path", "", "Path to file")
	flag.Parse()

	if *file_path == "" {
		log.Fatal("No file path provided")
	}

	// Open the CSV file
	file, err := os.Open(*file_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read in the CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV data:", err)
		return
	}

	// Parse the data into a table struct
	table := newTable()
	for i, record := range records {
		if i == 0 {
			for _, v := range record {
				table.columns = append(table.columns, v)
			}
			continue
		}
		row := make([]interface{}, 0)
		for _, v := range record {
			row = append(row, v)
		}
		table.rows = append(table.rows, row)
	}

	table.printTable()

}

func newTable() GenericTable {
	var t GenericTable
	t.columns = make(table.Row, 0)
	t.rows = make([]table.Row, 0)
	return t
}

func (t GenericTable) printTable() {
	table_formatter := table.NewWriter()
	table_formatter.SetOutputMirror(os.Stdout)
	table_formatter.SetTitle(t.name)
	table_formatter.AppendHeader(t.columns)
	table_formatter.AppendRows(t.rows)
	table_formatter.SetStyle(table.StyleLight)
	table_formatter.Style().Options.SeparateRows = true
	table_formatter.Style().Options.SeparateHeader = true
	table_formatter.Style().Box.MiddleHorizontal = "-"
	table_formatter.Style().Box.MiddleSeparator = "+"
	table_formatter.Style().Box.MiddleVertical = "|"
	table_formatter.Style().Options.DrawBorder = true
	table_formatter.Render()
}
