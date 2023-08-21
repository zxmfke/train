package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayList_String(t *testing.T) {
	root := NewArrayListRoot()

	t.Logf("%s", root)

	root.Set(1, 1)
	root.Set(2, 2)

	t.Logf("%s", root)
}

func TestArrayList_NewArrayListRootWithSize(t *testing.T) {
	root := NewArrayListRootWithSize(10)

	_, has := root.Get(1)
	assert.Equal(t, false, has)

	root.Set(2, 2)
	t.Logf("%s", root)
}

func TestArrayList_Set(t *testing.T) {
	root := NewArrayListRoot()

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

func TestArrayList_Get(t *testing.T) {
	root := NewArrayListRoot()
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

func TestArrayList_Delete(t *testing.T) {
	root := NewArrayListRoot()
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
			wantPrint: "empty ArrayList",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			root.Delete(tt.delKey)

			_, _, has := root.search(tt.delKey)

			assert.Equal(t, tt.wantHas, has)

			t.Logf("%s", root)

			if tt.print {
				assert.Equal(t, fmt.Sprintf("%s", root), tt.wantPrint)
			}
		})
	}
}

func TestArrayList_Search(t *testing.T) {
	root := NewArrayListRoot()
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
			v, has := root.Search(tt.searchKey)

			assert.Equal(t, has, tt.wantHas)
			if has {
				assert.Equal(t, v.(int), tt.wantValue)
			}
		})
	}
}

func TestArrayList_DeleteTail(t *testing.T) {
	root := NewArrayListRoot()
	root.Set(1, 1)
	root.Set(2, 2)
	root.Set(3, 3)
	root.Set(4, 3)

	t.Logf("%s", root)

	cases := []struct {
		name   string
		expect int
		has    bool
	}{
		{
			name:   "删除尾部第1个",
			expect: 3,
			has:    true,
		},
		{
			name:   "删除尾部第2个",
			expect: 3,
			has:    true,
		},
		{
			name:   "删除尾部第3个",
			expect: 2,
			has:    true,
		},
		{
			name:   "删除尾部第4个",
			expect: 1,
			has:    true,
		},
		{
			name:   "删除一个空的list",
			expect: 0,
			has:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v, has := root.DeleteTail()

			t.Logf("%s", root)

			assert.Equal(t, tt.has, has)

			if has {
				assert.Equal(t, tt.expect, v.(int))
			}

		})
	}

}

func TestArrayList_DeleteHead(t *testing.T) {
	root := NewArrayListRoot()
	root.Set(1, 1)
	root.Set(2, 2)
	root.Set(3, 3)
	root.Set(4, 3)

	t.Logf("%s", root)

	cases := []struct {
		name   string
		expect int
		has    bool
	}{
		{
			name:   "删除头部第1个",
			expect: 1,
			has:    true,
		},
		{
			name:   "删除头部第2个",
			expect: 2,
			has:    true,
		},
		{
			name:   "删除头部第3个",
			expect: 3,
			has:    true,
		},
		{
			name:   "删除头部第4个",
			expect: 3,
			has:    true,
		},
		{
			name:   "删除一个空的list",
			expect: 0,
			has:    false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			v, has := root.DeleteHead()

			t.Logf("%s", root)

			assert.Equal(t, tt.has, has)

			if has {
				assert.Equal(t, tt.expect, v.(int))
			}

		})
	}

}
