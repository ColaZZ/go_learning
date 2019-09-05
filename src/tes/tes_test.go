package tes

import (
	"fmt"
	"testing"
)

var container = []string{"zero", "one", "two"}

func TestContainer(t *testing.T) {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	value, ok := interface{}(container).(map[int]string)
	fmt.Println(container)
	fmt.Println(value, ok)
	fmt.Printf("The element is %q.\n", container[1])
}
