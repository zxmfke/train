1. copy 用法写反，应该是 dest 在前
- copy(o.data[insertIndex:], o.data[insertIndex+1:]) wrong
- copy(o.data[insertIndex+1:], o.data[insertIndex:]) right

2. Sorter 的比较需要加上等于，比大小，4 < 4 不应该认为小，而是一样，所以加上等号
3. findKey，递归循环的边界，如果是位置是 0，然后又比 index 0 的 key 大，middleIndex 要加 1，而不是直接返回 middleIndex
4. Delete，删除最后一个元素，不能只是最后一个元素置为 nil，同时要把这个 o.data减少一个元素