package main

import (
	"fmt"
	"train/excel"
)

func main() {
	operator := excel.New([]excel.SheetConf{
		{Name: "测试中文Tab", ColLen: 20}, {Name: "测试第二个Tab", ColLen: 10},
	})

	headers := []string{}

	for i := 0; i < 20; i++ {
		headers = append(headers, fmt.Sprintf("%d", i))
	}

	if err := operator.WriteHeader(headers); err != nil {
		fmt.Println(fmt.Errorf("write header err : %s", err.Error()))
		return
	}

	data := make([][]interface{}, 5)

	for i := 0; i < 5; i++ {
		data[i] = make([]interface{}, 20)
		for j := 0; j < 20; j++ {
			data[i][j] = fmt.Sprintf("%d%d", i, j)
		}
	}

	if err := operator.WriteAll(data); err != nil {
		fmt.Println(fmt.Errorf("write data err : %s", err.Error()))
		return
	}

	operator.SwitchSheet("测试第二个Tab")

	data2 := make([][]interface{}, 4)

	for i := 0; i < 4; i++ {
		data2[i] = make([]interface{}, 10)
		for j := 0; j < 10; j++ {
			data2[i][j] = fmt.Sprintf("%d%d", i, j)
		}
	}

	headers2 := []string{}

	for i := 0; i < 10; i++ {
		headers2 = append(headers2, fmt.Sprintf("%d", i))
	}

	if err := operator.WriteHeader(headers2); err != nil {
		fmt.Println(fmt.Errorf("write header err : %s", err.Error()))
		return
	}

	if err := operator.WriteHeader(headers2); err != nil {
		fmt.Println(fmt.Errorf("write header err : %s", err.Error()))
		return
	}

	if err := operator.WriteAll(data2); err != nil {
		fmt.Println(fmt.Errorf("write data err : %s", err.Error()))
		return
	}

	_ = operator.SaveAs("testdata/new.xlsx")
}
