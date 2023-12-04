package iteration

func Repeat(char string, repetitions int) string {
	var value string
	for i := 0; i < repetitions; i++ {
		value = value + char
	}
	return value
}
