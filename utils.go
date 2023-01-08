package utils

import "fmt"

func init() {
	fmt.Println("INIT UTILS")
}

func InSlice(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
func Containsint(a []int, x int) bool {
	for _, value := range a {
		if value == x {
			return true
		}
	}
	return false
}
