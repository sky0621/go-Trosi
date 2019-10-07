package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

func main() {
	excelFileName := "テーブル定義書「商品」.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		os.Exit(-1)
	}
	for sn, sheet := range xlFile.Sheets {
		fmt.Printf("\n<SheetNo: %d>\n", sn)
		for rn, row := range sheet.Rows {
			fmt.Printf("\n<RowNo: %d>\n", rn)
			for cn, cell := range row.Cells {
				text := cell.String()
				//nf := cell.GetNumberFormat()
				//fm := cell.Formula()
				//st := cell.GetStyle()
				//fmt.Printf("[Value: %s][NumberFormat: %s][Formula: %s][Style: %#v]\n\n", text, nf, fm, st)
				fmt.Printf("[%d] %s\n", cn, text)
			}
		}
	}
}
