package hashTable

import (
	"github.com/go-playground/assert"
	"testing"
)

func TestNewHSMap(t *testing.T) {

	hsMap, _ := NewHSTable(nil, 20)

	_ = hsMap.Set(1, 111)
	_ = hsMap.Set(2, 2222)

	v, _ := hsMap.Get(1)
	t.Logf("%v", v)

	_ = hsMap.Delete(2)
	_, ok := hsMap.Get(2)
	t.Logf("%v", ok)
}

func TestNewSelfHFHSMap(t *testing.T) {
	hsMap, _ := NewHSTable(testHasFunc, 20)

	_ = hsMap.Set(1, 111)
	_ = hsMap.Set(2, 2222)

	v, _ := hsMap.Get(1)
	t.Logf("%v", v)
}

func testHasFunc(a int) int {
	return a % 20
}

func TestKeyOutOfRange(t *testing.T) {
	hsMap, _ := NewHSTable(keyOutOfRangeFunc, 10)

	err := hsMap.Set(10, 111)
	assert.Equal(t, ErrHSKeyOutOfRange, err)

	err = hsMap.Delete(10)
	assert.Equal(t, ErrHSKeyOutOfRange, err)

	_, has := hsMap.Get(10)
	assert.Equal(t, false, has)

}

func keyOutOfRangeFunc(a int) int {
	return a * 2
}

func TestHSMapValueSizeInvalie(t *testing.T) {
	_, err := NewHSTable(keyOutOfRangeFunc, 3)

	assert.Equal(t, ErrHSMapValueSizeInvalid, err)
}
