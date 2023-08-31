package slc_test

import (
	"fmt"

	"github.com/PlanitarInc/slc"
)

func Example_Combo() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(slc.Includes(nums, 4)) // true
	fmt.Println(slc.Includes(nums, 9)) // false

	fmt.Println(slc.Every(nums, func(n int) bool {
		return n > 0
	})) // true
	fmt.Println(slc.Every(nums, func(n int) bool {
		return n%2 == 1
	})) // false

	type C struct {
		S string
		N int
	}

	objs := []C{
		{"a", 9},
		{"b", 8},
		{"c", 7},
	}
	fmt.Println(slc.Filter(objs, func(o C) bool {
		return o.N > 7
	})) // [{a 9}, {b 8}]
	fmt.Println(slc.FilterOut(objs, func(o C) bool {
		return o.S == "b"
	})) // [{a 9}, {c 7}]

	// Output:
	// true
	// false
	// true
	// false
	// [{a 9} {b 8}]
	// [{a 9} {c 7}]
}
