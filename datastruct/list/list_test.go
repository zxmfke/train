package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_String(t *testing.T) {
	root := NewListRoot()

	t.Logf("%s", root)

	root.Set(1, 1)
	root.Set(2, 2)

	t.Logf("%s", root)
}

func TestNewListRootWithInit(t *testing.T) {
	root := NewListRootWithInit(1, 1)

	_, has := root.Get(1)
	assert.Equal(t, true, has)

	root.Set(2, 2)
	t.Logf("%s", root)
}

func TestList_Set(t *testing.T) {
	root := NewListRoot()

	cases := []struct {
		name      string
		setKey    int
		setValue  interface{}
		wantHas   bool
		wantValue interface{}
	}{
		{
			name:      "插入第一个 key",
			setKey:    1,
			setValue:  1,
			wantHas:   true,
			wantValue: 1,
		}, {
			name:      "插入第二个 key",
			setKey:    2,
			setValue:  2,
			wantHas:   true,
			wantValue: 2,
		}, {
			name:      "插入相同的 key",
			setKey:    2,
			setValue:  3,
			wantHas:   true,
			wantValue: 3,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root.Set(tt.setKey, tt.setValue)

			v, has := root.Get(tt.setKey)

			assert.Equal(t, has, tt.wantHas)
			if has {
				assert.Equal(t, v, tt.wantValue)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	root := NewListRoot()
	root.Set(1, 1)
	root.Set(2, 2)

	cases := []struct {
		name      string
		searchKey int
		wantHas   bool
		wantValue interface{}
	}{
		{
			name:      "搜索一个不存在的 key",
			searchKey: 3,
			wantHas:   false,
			wantValue: 0,
		}, {
			name:      "搜索第一个 key",
			searchKey: 1,
			wantHas:   true,
			wantValue: 1,
		}, {
			name:      "搜索 nextNode 的 key",
			searchKey: 2,
			wantHas:   true,
			wantValue: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v, has := root.Get(tt.searchKey)

			assert.Equal(t, has, tt.wantHas)
			if has {
				assert.Equal(t, v, tt.wantValue)
			}
		})
	}
}

func TestList_Delete(t *testing.T) {
	root := NewListRoot()
	root.Set(1, 1)
	root.Set(2, 2)
	root.Set(3, 3)
	root.Set(4, 3)

	cases := []struct {
		name      string
		delKey    int
		wantHas   bool
		print     bool
		wantPrint string
	}{
		{
			name:    "删除一个不存在的 key",
			delKey:  5,
			wantHas: false,
		}, {
			name:    "删除中间的 key",
			delKey:  2,
			wantHas: false,
		}, {
			name:    "删除头的 key",
			delKey:  1,
			wantHas: false,
		}, {
			name:    "删除尾的 key",
			delKey:  4,
			wantHas: false,
		}, {
			name:      "删空",
			delKey:    3,
			wantHas:   false,
			print:     true,
			wantPrint: "empty List",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root.Delete(tt.delKey)

			_, has := root.search(tt.delKey)

			assert.Equal(t, tt.wantHas, has)

			if tt.print {
				assert.Equal(t, fmt.Sprintf("%s", root), tt.wantPrint)
			}
		})
	}
}

func TestList_Search(t *testing.T) {
	root := NewListRoot()
	root.Set(1, 1)
	root.Set(2, 2)

	cases := []struct {
		name      string
		searchKey int
		wantHas   bool
		wantValue interface{}
	}{
		{
			name:      "搜索一个不存在的 key",
			searchKey: 3,
			wantHas:   false,
			wantValue: 0,
		}, {
			name:      "搜索第一个 key",
			searchKey: 1,
			wantHas:   true,
			wantValue: 1,
		}, {
			name:      "搜索 nextNode 的 key",
			searchKey: 2,
			wantHas:   true,
			wantValue: 2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			node, has := root.search(tt.searchKey)

			assert.Equal(t, has, tt.wantHas)
			if has {
				assert.Equal(t, node.v, tt.wantValue)
			}
		})
	}
}
