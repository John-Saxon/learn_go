package main

import (
	"fmt"
	"log"

	"github.com/extrame/xls"
)

func main() {
	// pwd, _ := os.Getwd()
	xlsPath := `test.xls`
	xlsFile, closer, err := xls.OpenWithCloser(xlsPath, "utf-8")
	if err != nil {
		log.Fatal(err)
	}
	// 获取xls文件的第一个sheet
	sheet := xlsFile.GetSheet(0)
	fmt.Println("max row : ", sheet.MaxRow)
	// 从第二行开始，遍历xls文件，然后按行调用insertRowFromXls函数
	for j := 0; j < int(sheet.MaxRow)+1; j++ {
		xlsRow := sheet.Row(j)
		fmt.Println("xlsRow: ", xlsRow)
		lastI := xlsRow.LastCol()
		fmt.Println("lastI: ", lastI)
		fI := xlsRow.FirstCol()
		fmt.Println("fI: ", fI)
		for i := 0; i < lastI; i++ {
			// fmt.Println("k: ", i)
			// fmt.Println("content: ", xlsRow.Col(i))
		}
		rowColCount := xlsRow.LastCol()
		fmt.Println("rowColCount: ", rowColCount)
	}
	closer.Close()
}

// get sheet count
// select sheet
// multi data type
// to csv
