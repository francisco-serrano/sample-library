package calculator

func Add(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func MultiMul(values ...int) int {
	accumulator := 0
	for _, v := range values {
		accumulator *= v
	}

	return accumulator
}

func DoSomething() string {
	return "something"
}
