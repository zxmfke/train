package excel

import (
	"testing"
)

func TestRead(t *testing.T) {
	cases := []struct {
		name     string
		fileName string
		expect   error
	}{
		{
			name:     "文件不存在",
			fileName: "testdata/new1.xlsx",
		},
		{
			name:     "文件存在",
			fileName: "testdata/new.xlsx",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = Read(tt.fileName, "测试中文Tab")
		})
	}
}
