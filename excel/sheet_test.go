package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitHeader(t *testing.T) {
	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 3}})

	// 第一次初始化header
	operator.getCurSheet().initHeader()

	// 已经初始化过的，不再初始化header
	operator.getCurSheet().initHeader()

	curHeader := operator.getCurSheet().rowCellTemplate

	expectHeader := []*cell{
		{
			Col:   "A",
			Row:   1,
			Value: nil,
		},
		{
			Col:   "B",
			Row:   1,
			Value: nil,
		},
		{
			Col:   "C",
			Row:   1,
			Value: nil,
		},
	}

	if len(expectHeader) != len(curHeader) {
		t.Errorf("header len wrong, expect : %d , result : %d", len(expectHeader), len(curHeader))
	}

	for i := 0; i < 3; i++ {
		assert.Equal(t, curHeader[i].Col, expectHeader[i].Col)
	}
}

func TestInitLongHeader(t *testing.T) {
	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 30}})

	// 初始化超过长度26的header
	operator.getCurSheet().initHeader()
}
