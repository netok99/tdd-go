package structinterfaces

import (
	"testing"
)

func TestPrimeter(t *testing.T) {
	actual := Perimeter(10.0, 10.0)
	expected := 40.0

	if actual != expected {
		t.Errorf("actual %.2f esperado %.2f", actual, expected)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		actual := shape.Area()

		if actual != expected {
			t.Errorf("actual %.2f, expected %.2f", actual, expected)
		}
	}

	t.Run("retangles", func(t *testing.T) {
		retangle := Retangle{12.0, 6.0}
		verifyArea(t, retangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		verifyArea(t, circle, 314.1592653589793)
	})
}

func TestTableDrivenArea(t *testing.T) {
	testArea := []struct {
		shape    Shape
		expected float64
	}{
		// naming
		{shape: Retangle{Width: 12, Height: 6}, expected: 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, test := range testArea {
		actual := test.shape.Area()
		if actual != test.expected {
			t.Errorf("actual %.2f, expected %.2f", actual, test.expected)
		}
	}
}
