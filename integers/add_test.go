package integers

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(2, 2)
	expected := 4

	if actual != expected {
		t.Errorf("actual '%d', expected '%d'", actual, expected)
	}
}
