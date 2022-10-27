package hashmap

import (
	"github.com/zxmfke/train/datastruct/list"
	"testing"
)

func TestNewHSMap(t *testing.T) {
	hsMap := NewHSMap(nil, list.NewListRoot())

	hsMap.Set("a", 111)
	hsMap.Set("b", 2222)

	v, _ := hsMap.Get("a")
	t.Logf("%v", v)

	hsMap.Delete("b")
	_, ok := hsMap.Get("b")
	t.Logf("%v", ok)
}
