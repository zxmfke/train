package excel

import "errors"

var (
	OutOfColumn = errors.New("超过设置的最大长度")
	EmptySheet  = errors.New("SHEET为空")
)
