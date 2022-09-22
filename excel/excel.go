package excel

import (
	"github.com/xuri/excelize/v2"
	"os"
	"path/filepath"
)

// Operator excel操作器
type Operator struct {
	curSheet  string
	excelFile *excelize.File
	sheetMap  map[string]*sheet
}

// SheetConf excel页签配置
type SheetConf struct {
	Name   string `json:"name"`
	ColLen int    `json:"colLen"`
}

var defaultSheets = []SheetConf{
	{
		Name:   "Sheet1",
		ColLen: 20,
	},
}

// New 初始化返回一个Excel操作类，默认使用Sheet1作为首页和当前操作的sheet
func New(sheets []SheetConf) *Operator {

	if len(sheets) == 0 {
		sheets = defaultSheets
	}

	var (
		defaultSheet = sheets[0]
		operator     = &Operator{
			curSheet:  defaultSheet.Name,
			excelFile: excelize.NewFile(),
			sheetMap:  make(map[string]*sheet),
		}
	)

	defer func() {
		// excelize.NewFile会默认创建一个Sheet1的sheet
		if len(sheets) != 0 {
			operator.excelFile.DeleteSheet("Sheet1")
		}
	}()

	for i := 0; i < len(sheets); i++ {
		sheetIndex := operator.newSheet(sheets[i].Name)
		operator.sheetMap[sheets[i].Name] = &sheet{
			index:           sheetIndex,
			name:            sheets[i].Name,
			colMaxLen:       sheets[i].ColLen,
			rowCellTemplate: nil,
		}
	}

	operator.SwitchSheet(defaultSheet.Name)

	return operator
}

// newSheet 创建新的sheet
func (o *Operator) newSheet(sheetName string) int {
	return o.excelFile.NewSheet(sheetName)
}

// SwitchSheet 切换excel的tab
func (o *Operator) SwitchSheet(sheetName string) {

	var (
		sheetTab = o.sheetMap[sheetName]
	)

	sheetTab.initHeader()
	o.excelFile.SetActiveSheet(sheetTab.index)
	o.curSheet = sheetName
}

// getCurSheet 获取当前sheet
func (o *Operator) getCurSheet() *sheet {
	return o.sheetMap[o.curSheet]
}

// SaveAs 存储，如果文件夹不存在则创建文件夹后存储
func (o *Operator) SaveAs(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0766); err != nil {
		return err
	}

	if err := o.excelFile.SaveAs(path); err != nil {
		return err
	}

	return nil
}

func (o *Operator) Close() error {
	return o.excelFile.Close()
}
