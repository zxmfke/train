package excel

import "fmt"

var alphabetMaxNo = 26
var headerAlphabet = map[int]string{
	0:  "A",
	1:  "B",
	2:  "C",
	3:  "D",
	4:  "E",
	5:  "F",
	6:  "G",
	7:  "H",
	8:  "I",
	9:  "J",
	10: "K",
	11: "L",
	12: "M",
	13: "N",
	14: "O",
	15: "P",
	16: "Q",
	17: "R",
	18: "S",
	19: "T",
	20: "U",
	21: "V",
	22: "W",
	23: "X",
	24: "Y",
	25: "Z",
}

type sheet struct {
	index           int
	name            string
	colMaxLen       int
	rowCellTemplate []*cell
}

type cell struct {
	Col   string // 横轴字母表
	Row   int    // 第几行
	Value interface{}
}

// location 获取此cell位置，比如A1 AA2
func (c *cell) location() string {
	return fmt.Sprintf("%s%d", c.Col, c.Row)
}

// initHeader 初始化col的Y轴
func (s *sheet) initHeader() {

	if len(s.rowCellTemplate) != 0 {
		return
	}

	for i := 0; i < s.colMaxLen; i++ {
		quotient := i / alphabetMaxNo
		mod := i % alphabetMaxNo
		key := headerAlphabet[mod]
		if quotient != 0 {
			key += headerAlphabet[quotient]
		}

		s.rowCellTemplate = append(s.rowCellTemplate, &cell{
			Col: key,
		})
	}
}
