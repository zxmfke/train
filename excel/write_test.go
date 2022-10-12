package excel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperator_WriteHeader(t *testing.T) {
	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 3}})

	cases := []struct {
		name      string
		header    []string
		expectErr error
	}{
		{
			name:      "写入header超过上限",
			header:    []string{"header1", "header2", "header3", "header4"},
			expectErr: OutOfColumn,
		},
		{
			name:      "正常写入header",
			header:    []string{"header1", "header2", "header3"},
			expectErr: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			err := operator.WriteHeader(tt.header)
			assert.Equal(t, err, tt.expectErr)
		})
	}

}

func TestOperator_WriteRow(t *testing.T) {
	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 3}})

	cases := []struct {
		name      string
		fileName  string
		data      []*cell
		expectErr error
	}{
		{
			name:     "写入data col长度超过上限",
			fileName: "testdata/write_row_out_of_col.xlsx",
			data: []*cell{
				{
					Col:   "A",
					Row:   2,
					Value: "data1",
				},
				{
					Col:   "B",
					Row:   2,
					Value: "data2",
				},
				{
					Col:   "C",
					Row:   2,
					Value: "data3",
				},
				{
					Col:   "D",
					Row:   2,
					Value: "data4",
				},
			},
			expectErr: OutOfColumn,
		},
		{
			name:     "正常写入data",
			fileName: "testdata/write_row.xlsx",
			data: []*cell{
				{
					Col:   "A",
					Row:   2,
					Value: "data1",
				},
				{
					Col:   "B",
					Row:   2,
					Value: "data2",
				},
				{
					Col:   "C",
					Row:   2,
					Value: "data3",
				},
			},
			expectErr: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			err := operator.WriteRow(tt.data...)
			assert.Equal(t, tt.expectErr, err)

			_ = operator.SaveAs(tt.fileName)

			//defer os.Remove(tt.fileName)
		})
	}

	defer func() {
		_ = operator.Close()
	}()
}

func TestOperator_WriteAll(t *testing.T) {
	operator := New([]SheetConf{{Name: "第一个tab", ColLen: 3}})

	cases := []struct {
		name      string
		fileName  string
		data      [][]string
		expectErr error
	}{
		{
			name:     "写入col超过上限，只写入col最长",
			fileName: "testdata/write_all_ouf_of_col.xlsx",
			data: [][]string{
				{},
				{"data1", "data2", "data3", "data4", "data5", "data6"},
			},
			expectErr: nil,
		},
		{
			name:     "正常写入data",
			fileName: "testdata/write_all.xlsx",
			data: [][]string{
				{"data1", "data2", "data3"},
				{"data4", "data5", "data6"},
			},
			expectErr: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			err := operator.WriteAll(tt.data)
			assert.Equal(t, tt.expectErr, err)

			_ = operator.SaveAs(tt.fileName)

			//defer os.Remove(tt.fileName)
		})
	}

	defer func() {
		_ = operator.Close()

	}()
}
