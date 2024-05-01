package math

func Sum[T int | float64](nums ...T) T {
	var s T
	for _, x := range nums {
		s = x + s
	}
	return s
}

func Mul[T int | float64](a T, b T) T {
	return a * b
}
