package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	a := []int{10, 9, 2, 8, 3, 4, 6}
	BubbleSort(a)
	t.Logf("%v\n", a)
}

func TestInsertionSort(t *testing.T) {
	a := []int{10, 9, 2, 8, 3, 4, 6}
	InsertionSort(a)
	t.Logf("%v\n", a)
}
