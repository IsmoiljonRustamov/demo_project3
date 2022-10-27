package main

func Triangle(a,b,c int) bool {
	if (a * a) + (b * b) == (c * c) {
		return true
	}
	return false
}