package array

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("collection 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		actual := Add(numbers)
		expect := 15

		if actual != expect {
			t.Errorf("actual %d, expect %d, data %v", actual, expect, numbers)
		}
	})

	t.Run("collection any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		actual := Add(numbers)
		expect := 6

		if actual != expect {
			t.Errorf("actual %d, expect %d, data %v", actual, expect, numbers)
		}
	})
}

func TestAddAll(t *testing.T) {
	t.Run("add all", func(t *testing.T) {
		actual := AddAll([]int{1, 2}, []int{0, 9})
		expect := []int{3, 9}

		// Go dont compare slices
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual %d, expect %d", actual, expect)
		}
	})

	t.Run("add all 2", func(t *testing.T) {
		actual := AddAll2([]int{1, 2}, []int{0, 9})
		expect := []int{3, 9}

		// Go dont compare slices
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf("actual %d, expect %d", actual, expect)
		}
	})
}

func TestAddRest(t *testing.T) {
	verifyAdds := func(t *testing.T, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	}

	t.Run("add slices", func(t *testing.T) {
		actual := AddRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		verifyAdds(t, actual, expected)
	})

	t.Run("add empty slices", func(t *testing.T) {
		actual := AddRest([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		verifyAdds(t, actual, expected)
	})
}
