package excel

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		name         string
		sheetConf    []SheetConf
		expectedName string
	}{
		{
			name:         "不传sheet，初始化一个excel",
			expectedName: "Sheet1",
		},
		{
			name:         "初始化一个sheet",
			sheetConf:    []SheetConf{{Name: "第一个tab", ColLen: 10}},
			expectedName: "第一个tab",
		},
		{
			name:         "初始化两个sheet",
			sheetConf:    []SheetConf{{Name: "第一个tab", ColLen: 10}, {Name: "第二个tab", ColLen: 20}},
			expectedName: "第一个tab",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			operator := New(tt.sheetConf)

			curSheet := operator.getCurSheet()

			assert.Equal(t, tt.expectedName, curSheet.name)

			_ = operator.Close()
		})
	}
}

func TestOperator_SwitchSheet(t *testing.T) {
	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 10}, {Name: "第二个tab", ColLen: 20}})

	curSheet := operator.getCurSheet()

	assert.Equal(t, "第一个tab", curSheet.name)

	operator.SwitchSheet("第二个tab")

	curSheet = operator.getCurSheet()

	assert.Equal(t, "第二个tab", curSheet.name)

	_ = operator.Close()
}

func TestOperator_SaveAs(t *testing.T) {

	newFile := "testdata/test.xlsx"
	newDirFile := "testdata/test2.xlsx"

	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 10}, {Name: "第二个tab", ColLen: 20}})

	if err := operator.SaveAs(newFile); err != nil {
		t.Errorf("save err : %s", err.Error())
		return
	}

	if err := operator.SaveAs(newDirFile); err != nil {
		t.Errorf("save with newDir err : %s", err.Error())
		return
	}

	defer func() {
		_ = os.Remove(newFile)
		_ = os.Remove(newDirFile)
	}()
}
