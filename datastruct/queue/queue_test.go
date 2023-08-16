package queue

import "testing"

func TestListQueue(t *testing.T) {
	listQueue := NewListQueue()

	listQueue.Enqueue(1, 1)
	listQueue.Enqueue(2, 2)

	t.Logf("%s", listQueue.Print())

	for {
		v, has := listQueue.Dequeue()
		if !has {
			break
		}

		t.Logf("pop value %v", v)
	}
}
