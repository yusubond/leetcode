package template

import "testing"

func TestMyQueue(t *testing.T) {
	q := NewMyQueue()
	for i := 1; i <= 5; i++ {
		q.EnQueue(i)
	}
	if q.IsEmpty() {
		t.Fatalf("fail, expect queue isn't empty")
	}

	for j := 1; j <= 5; j++ {
		x := q.DeQueue()
		if x != j {
			t.Fatalf("fail dequeue, expect %d but got %d\n", j, x)
		}
	}

	if !q.IsEmpty() {
		t.Fatalf("fail, expect queue is empty")
	}
}
