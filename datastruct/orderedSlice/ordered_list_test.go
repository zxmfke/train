package orderedSlice

import "testing"

func TestOrderedSlice(t *testing.T) {

	orderList := NewOrderedSlice(OrderDesc)

	orderList.Insert(1, "A")
	orderList.Insert(3, "C")

	t.Logf("%s", orderList)

	orderList.Insert(2, "B")

	t.Logf("%s", orderList)

	orderList.Insert(4, "B")

	t.Logf("%s", orderList)

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
