package rev

import (
	"fmt"
	"testing"
)

// page65
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 4.3 指定长度，使用指针避免值拷贝
func reverse2(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

// func TestReverse(t *testing.T) {
// 	a := [...]int{0, 1, 2, 3, 4, 5}
// 	reverse2(&a)
// 	fmt.Println(a)
// }

// 练习4.4 倒序append
func rotate(s []int) []int {
	r := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}

// func TestRotate(t *testing.T) {
// 	a := [...]int{0, 1, 2, 3, 4, 5}
// 	fmt.Println(a[:])
// 	q := rotate(a[:])
// 	fmt.Println(q)
// }

// 练习4.5
func DeleteRepeat(s []string) []string {
	if len(s) < 1 {
		return s
	}
	dist := s[0:1]
	for i := 1; i < len(s); i++ {
		for s[i] != dist[len(dist)-1] {
			dist = append(dist, s[i])
		}
	}
	return dist
}

func TestDeleteRepeat(t *testing.T) {
	s := []string{"a", "b", "c", "d", "d", "f", "g"}
	ss := DeleteRepeat(s)
	fmt.Println(ss)
}
