package array

func Add(numbers []int) int {
	value := 0
	for _, number := range numbers {
		value += number
	}
	return value
}

func AddAll(add_numbers ...[]int) (adds []int) {
	adds = make([]int, len(add_numbers))
	for i, numbers := range add_numbers {
		adds[i] = Add(numbers)
	}
	return
}

func AddAll2(add_numbers ...[]int) (adds []int) {
	for _, numbers := range add_numbers {
		adds = append(adds, Add(numbers))
	}
	return
}

func AddRest(add_numbers ...[]int) []int {
	var adds []int
	for _, numbers := range add_numbers {
		if len(numbers) == 0 {
			adds = append(adds, 0)
		} else {
			final := numbers[1:]
			adds = append(adds, Add(final))
		}
	}
	return adds
}
