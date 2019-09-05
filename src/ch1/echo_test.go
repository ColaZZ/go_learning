package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// func TestEcho1(t *testing.T) {
// 	// s, sep := "", ""
// 	fmt.Println(os.Args[0])
// }

// func TestEchho2(t *testing.T) {
// 	for k, v := range os.Args[1:] {
// 		fmt.Printf("%d\t%s\n", k, v)
// 	}
// }

func TestEcho3(t *testing.T) {
	s, sep := "", ""
	time1 := time.Now()
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	// fmt.Println(s)
	time2 := time.Now()
	fmt.Printf("%d %d", time.Since(time1), time.Since(time2))
	// fmt.Println(strings.Join(os.Args[1:], " "))
	time3 := time.Now()
	fmt.Printf("%d %d", time.Since(time2), time.Since(time3))
}
