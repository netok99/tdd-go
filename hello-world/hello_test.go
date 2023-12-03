package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	messageVerification := func(t *testing.T, actual, expected string) {
		t.Helper()
		if actual != expected {
			t.Errorf("actual '%s', expected '%s'", actual, expected)
		}
	}

	t.Run("say hello with name", func(t *testing.T) {
		actual := Ola("Neto", "")
		expect := "Hello Neto"
		messageVerification(t, actual, expect)
	})

	t.Run("say hello with name", func(t *testing.T) {
		actual := Ola("Neto", "english")
		expect := "Hello Neto"
		messageVerification(t, actual, expect)
	})

	t.Run("say hello spanish with name", func(t *testing.T) {
		actual := Ola("Neto", "spanish")
		expect := "Hola Neto"
		messageVerification(t, actual, expect)
	})

	t.Run("say hello portuguese with name", func(t *testing.T) {
		actual := Ola("Neto", "portuguese")
		expect := "Olá Neto"
		messageVerification(t, actual, expect)
	})

	t.Run("say hello without name", func(t *testing.T) {
		actual := Ola("", "")
		expect := "Hello"
		messageVerification(t, actual, expect)
	})

	t.Run("say ola without name", func(t *testing.T) {
		actual := Ola("", "portuguese")
		expect := "Olá"
		messageVerification(t, actual, expect)
	})

	t.Run("say hola without name", func(t *testing.T) {
		actual := Ola("", "spanish")
		expect := "Hola"
		messageVerification(t, actual, expect)
	})
}
