package orderedSlice

import (
	"github.com/go-playground/assert"
	"testing"
)

func TestOrderedSlice_String(t *testing.T) {

	orderList := NewOrderedSlice()

	// 打印空
	t.Logf("%s", orderList)

	orderList.Insert(1, "A")
	t.Logf("%s", orderList)
}

func TestOrderedSlice_Insert(t *testing.T) {

	cases := []struct {
		name      string
		insertKey int
		wantOrder []int
	}{
		{
			name:      "插入第一个",
			insertKey: 1,
			wantOrder: []int{1},
		},
		{
			name:      "插入第二个，比第一个大",
			insertKey: 3,
			wantOrder: []int{1, 3},
		},
		{
			name:      "插入第三个，比第二个小",
			insertKey: 2,
			wantOrder: []int{1, 2, 3},
		},
		{
			name:      "插入第四个，比第一个小",
			insertKey: 0,
			wantOrder: []int{0, 1, 2, 3},
		},
		{
			name:      "插入第五个，已存在，更新",
			insertKey: 1,
			wantOrder: []int{0, 1, 2, 3},
		},
	}

	orderList := NewOrderedSlice()

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			orderList.Insert(tt.insertKey, "A")

			if len(tt.wantOrder) != orderList.count {
				assert.Equal(t, len(tt.wantOrder), orderList.count)
				return
			}

			for i := 0; i < len(tt.wantOrder); i++ {
				assert.Equal(t, tt.wantOrder[i], orderList.data[i].K)
			}
		})
	}
}

func TestFindInsertIndex(t *testing.T) {
	cases := []struct {
		name       string
		insertKey  int
		wantOrder  []int
		wantIndex  int
		wantAction insertAction
	}{
		{
			name:       "插入第二个，比第一个大",
			insertKey:  3,
			wantOrder:  []int{1, 3},
			wantIndex:  1,
			wantAction: Nothing,
		},
		{
			name:       "插入第三个，比第二个小",
			insertKey:  2,
			wantOrder:  []int{1, 2, 3},
			wantIndex:  1,
			wantAction: MoveInsertToNext,
		},
		{
			name:       "插入第四个，比第一个小",
			insertKey:  0,
			wantOrder:  []int{0, 1, 2, 3},
			wantIndex:  0,
			wantAction: MoveInsertToNext,
		},
		{
			name:       "插入第五个，已存在，更新",
			insertKey:  1,
			wantOrder:  []int{0, 1, 2, 3},
			wantIndex:  1,
			wantAction: Update,
		},
	}

	orderList := NewOrderedSlice()
	orderList.Insert(1, "A")

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			index, action := orderList.findKeyInsertIndex(0, orderList.count-1, tt.insertKey)
			assert.Equal(t, tt.wantIndex, index)
			assert.Equal(t, tt.wantAction, action)

			orderList.Insert(tt.insertKey, "A")
		})
	}
}

func TestOrderedSlice_SearchKey(t *testing.T) {
	cases := []struct {
		name       string
		insertList []int
		searchKey  int
		wantKey    int
		wantHas    bool
	}{
		{
			name:       "查询空",
			insertList: nil,
			searchKey:  0,
			wantKey:    -1,
			wantHas:    false,
		},
		{
			name:       "查询存在的值",
			insertList: []int{1, 2, 3},
			searchKey:  2,
			wantKey:    2,
			wantHas:    true,
		},
		{
			name:       "查询首个key",
			insertList: []int{1, 2},
			searchKey:  1,
			wantKey:    1,
			wantHas:    true,
		},
		{
			name:       "查询末尾key",
			insertList: []int{1, 2},
			searchKey:  2,
			wantKey:    2,
			wantHas:    true,
		},
		{
			name:       "查询不存在的key",
			insertList: []int{1},
			searchKey:  0,
			wantKey:    -1,
			wantHas:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			orderList := NewOrderedSlice()

			for i := 0; i < len(tt.insertList); i++ {
				orderList.Insert(tt.insertList[i], "A")
			}

			node, has := orderList.SearchKey(tt.searchKey)
			if has {
				assert.Equal(t, tt.wantKey, node.K)
			}

			assert.Equal(t, tt.wantHas, has)

		})
	}
}

func TestOrderedSlice_SearchKeyRange(t *testing.T) {
	cases := []struct {
		name       string
		insertList []int
		min        int
		max        int
		wantKey    []int
	}{
		{
			name:       "查询空",
			insertList: nil,
			min:        0,
			max:        1,
			wantKey:    nil,
		},
		{
			name:       "查询min max相反",
			insertList: nil,
			min:        1,
			max:        0,
			wantKey:    nil,
		},
		{
			name:       "查询中间数据",
			insertList: []int{1, 2, 3, 4, 5},
			min:        2,
			max:        4,
			wantKey:    []int{2, 3, 4},
		}, {
			name:       "查询包含头",
			insertList: []int{1, 2, 3, 4, 5},
			min:        1,
			max:        3,
			wantKey:    []int{1, 2, 3},
		}, {
			name:       "查询包含尾",
			insertList: []int{1, 2, 3, 4, 5},
			min:        3,
			max:        5,
			wantKey:    []int{3, 4, 5},
		}, {
			name:       "查询不全在orderedSlice里的元素",
			insertList: []int{1, 2, 4},
			min:        3,
			max:        5,
			wantKey:    []int{4},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			orderList := NewOrderedSlice()

			for i := 0; i < len(tt.insertList); i++ {
				orderList.Insert(tt.insertList[i], "A")
			}

			nodes := orderList.SearchKeyRange(tt.min, tt.max)
			if len(nodes) != len(tt.wantKey) {
				assert.Equal(t, len(tt.wantKey), len(nodes))
				return
			}

			for i := 0; i < len(tt.wantKey); i++ {
				assert.Equal(t, tt.wantKey[i], nodes[i].K)
			}

		})
	}
}

func TestOrderedSlice_Delete(t *testing.T) {

	cases := []struct {
		name       string
		insertList []int
		delKey     int
		wantKey    int
		wantIndex  int
	}{
		{
			name:       "删除空",
			insertList: nil,
			delKey:     0,
			wantKey:    -1,
			wantIndex:  -1,
		},
		{
			name:       "删除存在的值",
			insertList: []int{1, 2, 3},
			delKey:     2,
			wantKey:    2,
			wantIndex:  1,
		},
		{
			name:       "删除首个key",
			insertList: []int{1, 2},
			delKey:     1,
			wantKey:    1,
			wantIndex:  0,
		},
		{
			name:       "删除末尾key",
			insertList: []int{1, 2},
			delKey:     2,
			wantKey:    2,
			wantIndex:  1,
		},
		{
			name:       "删除不存在的key",
			insertList: []int{1},
			delKey:     2,
			wantKey:    -1,
			wantIndex:  -1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			orderList := NewOrderedSlice()

			for i := 0; i < len(tt.insertList); i++ {
				orderList.Insert(tt.insertList[i], "A")
			}

			node, delIndex := orderList.Delete(tt.delKey)

			if node != nil {
				assert.Equal(t, tt.wantKey, node.K)
			}
			assert.Equal(t, tt.wantIndex, delIndex)
		})
	}

}
