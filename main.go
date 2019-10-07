package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

// TODO: 機能実現優先の雑実装
func main() {
	/*
	 * Input -> Structured
	 */
	excelFileName := "テーブル定義書「商品」.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		os.Exit(-1)
	}

	var worksheets []*WorkSheet

	for sn, sheet := range xlFile.Sheets {
		ws := &WorkSheet{No: sn, Name: sheet.Name}
		for rn, row := range sheet.Rows {
			r := &Row{No: rn}
			for cn, cell := range row.Cells {
				c := &Cell{No: cn, Val: cell.String()}
				r.Cells = append(r.Cells, c)
			}
			ws.Rows = append(ws.Rows, r)
		}
		worksheets = append(worksheets, ws)
	}

	fmt.Println(worksheets)

	/*
	 * To Migration File
	 */

	/*
	 * To Go Struct
	 */
}

type WorkSheet struct {
	No   int
	Name string
	Rows []*Row
}

type Row struct {
	No    int
	Cells []*Cell
}

type Cell struct {
	No  int
	Val string
}
