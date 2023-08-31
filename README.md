# slc [![GoDoc](https://godoc.org/github.com/PlanitarInc/slc?status.svg)](https://godoc.org/github.com/PlanitarInc/slc) [![CI](https://github.com/PlanitarInc/slc/actions/workflows/go.yml/badge.svg)](https://github.com/PlanitarInc/slc/actions/workflows/go.yml) [![Coverage Status](https://coveralls.io/repos/github/PlanitarInc/slc/badge.svg)](https://coveralls.io/github/PlanitarInc/slc)

A generics-based toolset for working with slices.

## Usage

### Download the package

```sh
go get github.com/PlanitarInc/slc
```

### Example

Playground link: https://go.dev/play/p/qu0t390uxt_p

```go
package main

import (
	"fmt"

	"github.com/PlanitarInc/slc"
)

func main() {
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
}
```