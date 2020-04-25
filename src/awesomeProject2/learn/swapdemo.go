package main

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}
