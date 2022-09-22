package excel

import "github.com/xuri/excelize/v2"

// Read excel文件读取，将读取到数据封装为map，表格第一行视为表头，表头的每一列作为map的key（请尽量避免表头中包含中文）
func Read(filePath, sheet string) ([]map[string]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	var result = make([]map[string]string, len(rows))
	if len(rows) <= 1 {
		return result, nil
	}

	headerKeys := rows[0]

	for rowIndex := 1; rowIndex < len(rows); rowIndex++ {
		row := make(map[string]string)
		for colIndex, col := range headerKeys {

			if colIndex > (len(rows[rowIndex]) - 1) {
				continue
			}

			row[col] = rows[rowIndex][colIndex]
		}
		result[rowIndex-1] = row
	}

	return result, nil
}
