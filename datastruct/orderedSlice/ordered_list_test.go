package orderedSlice

import "testing"

func TestOrderedSlice(t *testing.T) {

	orderList := NewOrderedSlice()

	orderList.Insert(1, "A")
	orderList.Insert(3, "C")

	t.Logf("%s", orderList)

	orderList.Insert(2, "B")

	t.Logf("%s", orderList)

	orderList.Insert(4, "B")

	t.Logf("%s", orderList)

	rangeNode := orderList.SearchKeyRange(2, 3)

	t.Logf("range [2,3]")
	for i := 0; i < len(rangeNode); i++ {
		t.Logf("%s", rangeNode[i])
	}

	rangeNode = orderList.SearchKeyRange(1, 1)

	t.Logf("range [1,1]")
	for i := 0; i < len(rangeNode); i++ {
		t.Logf("%s", rangeNode[i])
	}

	rangeNode = orderList.SearchKeyRange(4, 4)

	t.Logf("range [4,4]")
	for i := 0; i < len(rangeNode); i++ {
		t.Logf("%s", rangeNode[i])
	}

	orderList.Delete(2)

	t.Logf("%s", orderList)

	_, has := orderList.SearchKey(2)

	t.Logf("has key 2 : %v", has)

	orderList.Delete(4)

	t.Logf("%s", orderList)

	orderList.Delete(1)

	t.Logf("%s", orderList)

	orderList.Delete(3)

	t.Logf("%s", orderList)

	orderList.Insert(1, "A")
	t.Logf("%s", orderList)

	orderList.Insert(1, "C")

	t.Logf("%s", orderList)
}
