package main

// Greatest Common Divisor
func gcd(x, y int) int {
	var tmp int
	for {
		tmp = (x % y)
		if tmp > 0 {
			x = y
			y = tmp
		} else {
			return y
		}
	}
}

// Least Common Multiple
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}
