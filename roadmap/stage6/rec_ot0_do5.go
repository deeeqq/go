package stage6

import (
	"fmt"
)

func schet(n int) int {
	if n == 6 {
		return n
	}
	fmt.Println(n)
	return n * schet(n+1)

}

func Ot0do5rec() {
	defer schet1(5)
	schet(0)
}
func schet1(b int) int {
	if b == 0 {
		return b
	}
	fmt.Println(b)
	return b * schet1(b-1)
}
