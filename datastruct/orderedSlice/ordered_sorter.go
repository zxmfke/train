package orderedSlice

type Sorter interface {
	Less(a, b int) bool
}

type OrderedListSortedWay int

const (
	_         OrderedListSortedWay = iota
	OrderIncr                      // 小到大
	OrderDesc                      // 大到小
)

func newSorter(sortWay OrderedListSortedWay) Sorter {
	switch sortWay {
	case OrderIncr:
		return incrSorter{}
	case OrderDesc:
		return descSorter{}
	default:
		return incrSorter{}
	}
}

type incrSorter struct{}

func (i incrSorter) Less(a, b int) bool {
	return a >= b
}

type descSorter struct{}

func (d descSorter) Less(a, b int) bool {
	return a <= b
}
