package excel

// WriteHeader 写入header
func (o *Operator) WriteHeader(headers []string) error {

	curSheet := o.getCurSheet()

	if len(headers) > curSheet.colMaxLen {
		return OutOfColumn
	}

	for i := 0; i < len(headers); i++ {
		curSheet.rowCellTemplate[i].Row = 1
		curSheet.rowCellTemplate[i].Value = headers[i]
	}

	return o.WriteRow(curSheet.rowCellTemplate...)
}

// WriteAll 写入所有的数据，不包括header
func (o *Operator) WriteAll(cells [][]interface{}) error {
	var (
		rows     = len(cells)
		rowBias  = 2
		curSheet = o.getCurSheet()
	)

	for i := 0; i < rows; i++ {
		cols := len(cells[i])

		if cols == 0 {
			continue
		}

		if cols > curSheet.colMaxLen {
			cols = curSheet.colMaxLen
		}

		for j := 0; j < cols; j++ {
			curSheet.rowCellTemplate[j].Row = i + rowBias
			curSheet.rowCellTemplate[j].Value = cells[i][j]
		}

		if err := o.WriteRow(curSheet.rowCellTemplate...); err != nil {
			return err
		}
	}

	return nil
}

// WriteRow 写入单行数据
func (o *Operator) WriteRow(cells ...*cell) error {
	curSheet := o.getCurSheet()

	if len(cells) > curSheet.colMaxLen {
		return OutOfColumn
	}

	for _, c := range cells {
		if err := o.excelFile.SetCellValue(o.curSheet, c.location(), c.Value); err != nil {
			return err
		}
	}
	return nil
}
