package ch_1

func add(x, y int) int {
	return x + y
}

func cover(x, y int) int {
	switch x {
	case 1:
		return x + y
	case 2:
		return x + y
	case 3:
		return x + y
	case 10:
		return x + y
	}
	return x + y
}
