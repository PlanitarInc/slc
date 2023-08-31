package slc

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_IncludesFunctions(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Includes(nil, 1)).To(BeFalse())
		Expect(Includes(nil, 0)).To(BeFalse())
		Expect(Includes([]int{}, 9)).To(BeFalse())
		Expect(Includes([]int{1, 2, 3, 4, 5}, 9)).To(BeFalse())
		Expect(Includes([]int{1, 2, 3, 4, 5}, 1)).To(BeTrue())
		Expect(Includes([]int{1, 2, 3, 4, 5}, 5)).To(BeTrue())
		Expect(Includes([]int{1, 2, 3, 4, 5}, 4)).To(BeTrue())

		isFour := func(n int) bool { return n == 4 }
		isTen := func(n int) bool { return n == 10 }
		Expect(IncludesFunc(nil, isFour)).To(BeFalse())
		Expect(IncludesFunc(nil, isTen)).To(BeFalse())
		Expect(IncludesFunc([]int{}, isFour)).To(BeFalse())
		Expect(IncludesFunc([]int{}, isTen)).To(BeFalse())
		Expect(IncludesFunc([]int{1, 2, 3, 4, 5}, isFour)).To(BeTrue())
		Expect(IncludesFunc([]int{1, 2, 3, 4, 5}, isTen)).To(BeFalse())
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Includes(nil, "")).To(BeFalse())
		Expect(Includes(nil, "zz")).To(BeFalse())
		Expect(Includes([]string{}, "")).To(BeFalse())
		Expect(Includes([]string{"a", "b", "c"}, "zz")).To(BeFalse())
		Expect(Includes([]string{"a", "b", "c"}, "a")).To(BeTrue())
		Expect(Includes([]string{"a", "b", "c"}, "b")).To(BeTrue())
		Expect(Includes([]string{"a", "b", "c"}, "c")).To(BeTrue())

		isB := func(s string) bool { return s == "b" }
		Expect(IncludesFunc(nil, isB)).To(BeFalse())
		Expect(IncludesFunc([]string{}, isB)).To(BeFalse())
		Expect(IncludesFunc([]string{"a"}, isB)).To(BeFalse())
		Expect(IncludesFunc([]string{"a", "b"}, isB)).To(BeTrue())
		Expect(IncludesFunc([]string{"a", "b", "c"}, isB)).To(BeTrue())
	})
}

func ExampleIncludes() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Includes(n, 0))
	fmt.Println(Includes(n, 1))
	fmt.Println(Includes(n, 4))
	fmt.Println(Includes(n, 5))
	fmt.Println(Includes(n, 6))
	// Output:
	// false
	// true
	// true
	// true
	// false
}

func ExampleIncludesFunc() {
	type C struct {
		N int
	}

	n := []C{{1}, {2}, {3}, {4}, {5}}

	fmt.Println(IncludesFunc(n, func(c C) bool { return c.N == 0 }))
	fmt.Println(IncludesFunc(n, func(c C) bool { return c.N == 4 }))
	fmt.Println(IncludesFunc(n, func(c C) bool { return c.N > 3 }))
	fmt.Println(IncludesFunc(n, func(c C) bool { return c.N > 5 }))
	// Output:
	// false
	// true
	// true
	// false
}

func Test_IndexFunctions(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Index(nil, 1)).To(Equal(-1))
		Expect(Index(nil, 0)).To(Equal(-1))
		Expect(Index([]int{}, 9)).To(Equal(-1))
		Expect(Index([]int{1, 2, 3, 4, 5}, 9)).To(Equal(-1))
		Expect(Index([]int{1, 2, 3, 4, 5}, 1)).To(Equal(0))
		Expect(Index([]int{1, 2, 3, 4, 5}, 5)).To(Equal(4))
		Expect(Index([]int{1, 2, 3, 4, 5}, 4)).To(Equal(3))

		isFour := func(n int) bool { return n == 4 }
		isTen := func(n int) bool { return n == 10 }
		Expect(IndexFunc(nil, isFour)).To(Equal(-1))
		Expect(IndexFunc(nil, isTen)).To(Equal(-1))
		Expect(IndexFunc([]int{}, isFour)).To(Equal(-1))
		Expect(IndexFunc([]int{}, isTen)).To(Equal(-1))
		Expect(IndexFunc([]int{1, 2, 3, 4, 5}, isFour)).To(Equal(3))
		Expect(IndexFunc([]int{1, 2, 3, 4, 5}, isTen)).To(Equal(-1))
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Index(nil, "")).To(Equal(-1))
		Expect(Index(nil, "zz")).To(Equal(-1))
		Expect(Index([]string{}, "")).To(Equal(-1))
		Expect(Index([]string{"a", "b", "c"}, "zz")).To(Equal(-1))
		Expect(Index([]string{"a", "b", "c"}, "a")).To(Equal(0))
		Expect(Index([]string{"a", "b", "c"}, "b")).To(Equal(1))
		Expect(Index([]string{"a", "b", "c"}, "c")).To(Equal(2))

		isB := func(s string) bool { return s == "b" }
		Expect(IndexFunc(nil, isB)).To(Equal(-1))
		Expect(IndexFunc([]string{}, isB)).To(Equal(-1))
		Expect(IndexFunc([]string{"a"}, isB)).To(Equal(-1))
		Expect(IndexFunc([]string{"a", "b"}, isB)).To(Equal(1))
		Expect(IndexFunc([]string{"a", "b", "c"}, isB)).To(Equal(1))
	})
}

func ExampleIndex() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Index(n, 0))
	fmt.Println(Index(n, 1))
	fmt.Println(Index(n, 4))
	fmt.Println(Index(n, 5))
	fmt.Println(Index(n, 6))
	// Output:
	// -1
	// 0
	// 3
	// 4
	// -1
}

func ExampleIndexFunc() {
	type C struct {
		N int
	}

	n := []C{{1}, {2}, {3}, {4}, {5}}

	fmt.Println(IndexFunc(n, func(c C) bool { return c.N == 0 }))
	fmt.Println(IndexFunc(n, func(c C) bool { return c.N == 4 }))
	fmt.Println(IndexFunc(n, func(c C) bool { return c.N > 2 }))
	fmt.Println(IndexFunc(n, func(c C) bool { return c.N > 5 }))
	// Output:
	// -1
	// 3
	// 2
	// -1
}

func TestEvery(t *testing.T) {
	RegisterTestingT(t)

	t.Run("none", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return false }

		Expect(Every(nil, matchFn)).To(BeTrue())
		Expect(Every([]int{}, matchFn)).To(BeTrue())
		Expect(Every([]int{1, 2, 3, 4, 5}, matchFn)).To(BeFalse())
	})

	t.Run("all", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return true }

		Expect(Every(nil, matchFn)).To(BeTrue())
		Expect(Every([]int{}, matchFn)).To(BeTrue())
		Expect(Every([]int{1, 2, 3, 4, 5}, matchFn)).To(BeTrue())
	})

	t.Run("filter", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return n%2 == 0 }

		Expect(Every(nil, matchFn)).To(BeTrue())
		Expect(Every([]int{}, matchFn)).To(BeTrue())
		Expect(Every([]int{1, 2, 3, 4, 5}, matchFn)).To(BeFalse())
	})

	t.Run("nonComprableType", func(t *testing.T) {
		RegisterTestingT(t)

		type tt struct {
			S string
			V bool
		}

		matchFn := func(v tt) bool { return v.V }

		Expect(Every(nil, matchFn)).To(BeTrue())
		Expect(Every([]tt{}, matchFn)).To(BeTrue())
		Expect(Every([]tt{
			{"one", true},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, matchFn)).To(BeFalse())
		Expect(Every([]tt{
			{"one", true},
			{"two", true},
			{"three", true},
			{"four", true},
			{"five", true},
		}, matchFn)).To(BeTrue())
	})
}

func ExampleEvery() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Every(n, func(n int) bool { return n > 0 }))
	fmt.Println(Every(n, func(n int) bool { return n < 3 }))
	// Output:
	// true
	// false
}

func TestSome(t *testing.T) {
	RegisterTestingT(t)

	t.Run("none", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return false }

		Expect(Some(nil, matchFn)).To(BeFalse())
		Expect(Some([]int{}, matchFn)).To(BeFalse())
		Expect(Some([]int{1, 2, 3, 4, 5}, matchFn)).To(BeFalse())
	})

	t.Run("all", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return true }

		Expect(Some(nil, matchFn)).To(BeFalse())
		Expect(Some([]int{}, matchFn)).To(BeFalse())
		Expect(Some([]int{1, 2, 3, 4, 5}, matchFn)).To(BeTrue())
	})

	t.Run("filter", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return n%2 == 0 }

		Expect(Some(nil, matchFn)).To(BeFalse())
		Expect(Some([]int{}, matchFn)).To(BeFalse())
		Expect(Some([]int{1, 2, 3, 4, 5}, matchFn)).To(BeTrue())
	})

	t.Run("nonComprableType", func(t *testing.T) {
		RegisterTestingT(t)

		type tt struct {
			S string
			V bool
		}

		matchFn := func(v tt) bool { return v.V }

		Expect(Some(nil, matchFn)).To(BeFalse())
		Expect(Some([]tt{}, matchFn)).To(BeFalse())
		Expect(Some([]tt{
			{"one", true},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, matchFn)).To(BeTrue())
	})
}

func ExampleSome() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Some(n, func(n int) bool { return n > 0 }))
	fmt.Println(Some(n, func(n int) bool { return n < 3 }))
	fmt.Println(Some(n, func(n int) bool { return n < 0 }))
	// Output:
	// true
	// true
	// false
}

func Test_FindFunctions(t *testing.T) {
	RegisterTestingT(t)

	t.Run("none", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return false }

		Expect(Find(nil, matchFn)).To(Equal(0))
		Expect(Find([]int{}, matchFn)).To(Equal(0))
		Expect(Find([]int{7}, matchFn)).To(Equal(0))
		Expect(Find([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal(0))

		Expect(FindPtr(nil, matchFn)).To(BeNil())
		Expect(FindPtr([]int{}, matchFn)).To(BeNil())
		Expect(FindPtr([]int{7}, matchFn)).To(BeNil())
		Expect(FindPtr([]int{1, 2, 3, 4, 5}, matchFn)).To(BeNil())
	})

	t.Run("all", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return true }

		Expect(Find(nil, matchFn)).To(Equal(0))
		Expect(Find([]int{}, matchFn)).To(Equal(0))
		Expect(Find([]int{7}, matchFn)).To(Equal(7))
		Expect(Find([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal(1))

		var s []int
		Expect(FindPtr(nil, matchFn)).To(BeNil())
		Expect(FindPtr([]int{}, matchFn)).To(BeNil())
		s = []int{7}
		Expect(FindPtr(s, matchFn)).To(BeIdenticalTo(&s[0]))
		s = []int{1, 2, 3, 4, 5}
		Expect(FindPtr([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal(&s[0]))
	})

	t.Run("filter", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return n%2 == 0 }

		Expect(Find(nil, matchFn)).To(Equal(0))
		Expect(Find([]int{}, matchFn)).To(Equal(0))
		Expect(Find([]int{7}, matchFn)).To(Equal(0))
		Expect(Find([]int{8}, matchFn)).To(Equal(8))
		Expect(Find([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal(2))

		var s []int
		Expect(FindPtr(nil, matchFn)).To(BeNil())
		Expect(FindPtr([]int{}, matchFn)).To(BeNil())
		Expect(FindPtr([]int{7}, matchFn)).To(BeNil())
		s = []int{8}
		Expect(FindPtr(s, matchFn)).To(BeIdenticalTo(&s[0]))
		s = []int{1, 2, 3, 4, 5}
		Expect(FindPtr(s, matchFn)).To(BeIdenticalTo(&s[1]))
	})

	t.Run("nonComprableType", func(t *testing.T) {
		RegisterTestingT(t)

		type tt struct {
			S string
			V bool
		}

		matchFn := func(v tt) bool { return v.V }

		Expect(Find(nil, matchFn)).To(Equal(tt{}))
		Expect(Find([]tt{}, matchFn)).To(Equal(tt{}))
		Expect(Find([]tt{
			{"one", false},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, matchFn)).To(Equal(tt{"two", true}))

		Expect(FindPtr(nil, matchFn)).To(BeNil())
		Expect(FindPtr([]tt{}, matchFn)).To(BeNil())
		s := []tt{
			{"one", false},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}
		Expect(FindPtr(s, matchFn)).To(BeIdenticalTo(&s[1]))
	})
}

func ExampleFind() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Find(n, func(n int) bool { return n > 0 }))
	fmt.Println(Find(n, func(n int) bool { return n > 3 }))
	fmt.Println(Find(n, func(n int) bool { return n < 0 }))
	// Output:
	// 1
	// 4
	// 0
}

func ExampleFindPtr() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(FindPtr(n, func(n int) bool { return n > 0 }) == &n[0])
	fmt.Println(FindPtr(n, func(n int) bool { return n > 3 }) == &n[3])
	fmt.Println(FindPtr(n, func(n int) bool { return n < 0 }) == nil)
	// Output:
	// true
	// true
	// true
}

func TestMap(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Map(nil, strconv.Itoa)).To(BeNil())
		Expect(Map([]int{}, strconv.Itoa)).To(BeNil())
		Expect(Map([]int{1}, strconv.Itoa)).To(Equal([]string{"1"}))
		Expect(Map([]int{1, 2}, strconv.Itoa)).To(Equal([]string{"1", "2"}))
		Expect(Map([]int{1, 2, 3, 4, 5}, strconv.Itoa)).To(Equal([]string{"1", "2", "3", "4", "5"}))
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		isAB := func(s string) bool { return s == "a" || s == "b" }

		Expect(Map(nil, isAB)).To(BeNil())
		Expect(Map([]string{}, isAB)).To(BeNil())
		Expect(Map([]string{"a", "b", "c"}, isAB)).To(Equal([]bool{true, true, false}))
	})

	t.Run("nonComprableType", func(t *testing.T) {
		RegisterTestingT(t)

		type tt struct {
			S string
			V bool
		}

		mapFn := func(v tt) string { return "_" + v.S + "_" }

		Expect(Map(nil, mapFn)).To(BeNil())
		Expect(Map([]tt{}, mapFn)).To(BeNil())

		Expect(Map([]tt{
			{"one", true},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, mapFn)).To(Equal([]string{
			"_one_",
			"_two_",
			"_three_",
			"_four_",
			"_five_",
		}))
	})
}

func ExampleMap() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(strings.Join(Map(n, strconv.Itoa), ","))
	// Output:
	// 1,2,3,4,5
}

func TestReduce(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		sum := func(acc, n int) int { return acc + n }

		Expect(Reduce(nil, sum)).To(Equal(0))
		Expect(Reduce([]int{}, sum)).To(Equal(0))
		Expect(Reduce([]int{1}, sum)).To(Equal(1))
		Expect(Reduce([]int{1, 2}, sum)).To(Equal(3))
		Expect(Reduce([]int{1, 2, 3, 4, 5}, sum)).To(Equal(15))
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		count := func(acc int, s string) int { return acc + 1 }

		Expect(Reduce(nil, count)).To(Equal(0))
		Expect(Reduce([]string{}, count)).To(Equal(0))
		Expect(Reduce([]string{"a", "b", "c"}, count)).To(Equal(3))
	})

	t.Run("nonComprableType", func(t *testing.T) {
		RegisterTestingT(t)

		type tt struct {
			S string
			V bool
		}

		concat := func(acc string, v tt) string { return acc + "," + v.S }

		Expect(Reduce(nil, concat)).To(Equal(""))
		Expect(Reduce([]tt{}, concat)).To(Equal(""))

		Expect(Reduce([]tt{
			{"one", true},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, concat)).To(Equal(",one,two,three,four,five"))
	})
}

func ExampleReduce() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Reduce(n, func(acc, n int) int { return acc + n }))
	// Output:
	// 15
}

func Test_FilterFunctions(t *testing.T) {
	RegisterTestingT(t)

	t.Run("none", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return false }

		Expect(Filter(nil, matchFn)).To(BeEmpty())
		Expect(FilterOut(nil, matchFn)).To(BeEmpty())
		Expect(Filter([]int{}, matchFn)).To(BeEmpty())
		Expect(FilterOut([]int{}, matchFn)).To(BeEmpty())
		Expect(Filter([]int{1, 2, 3, 4, 5}, matchFn)).To(BeEmpty())
		Expect(FilterOut([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal([]int{1, 2, 3, 4, 5}))
	})

	t.Run("all", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return true }

		Expect(Filter(nil, matchFn)).To(BeEmpty())
		Expect(FilterOut(nil, matchFn)).To(BeEmpty())
		Expect(Filter([]int{}, matchFn)).To(BeEmpty())
		Expect(FilterOut([]int{}, matchFn)).To(BeEmpty())
		Expect(Filter([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal([]int{1, 2, 3, 4, 5}))
		Expect(FilterOut([]int{1, 2, 3, 4, 5}, matchFn)).To(BeEmpty())
	})

	t.Run("filter", func(t *testing.T) {
		RegisterTestingT(t)

		matchFn := func(n int) bool { return n%2 == 0 }

		Expect(Filter(nil, matchFn)).To(BeEmpty())
		Expect(FilterOut(nil, matchFn)).To(BeEmpty())
		Expect(Filter([]int{}, matchFn)).To(BeEmpty())
		Expect(FilterOut([]int{}, matchFn)).To(BeEmpty())
		Expect(Filter([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal([]int{2, 4}))
		Expect(FilterOut([]int{1, 2, 3, 4, 5}, matchFn)).To(Equal([]int{1, 3, 5}))
	})

	t.Run("nonComprableType", func(t *testing.T) {
		RegisterTestingT(t)

		type tt struct {
			S string
			V bool
		}

		matchFn := func(v tt) bool { return v.V }

		Expect(Filter(nil, matchFn)).To(BeEmpty())
		Expect(FilterOut(nil, matchFn)).To(BeEmpty())
		Expect(Filter([]tt{}, matchFn)).To(BeEmpty())
		Expect(FilterOut([]tt{}, matchFn)).To(BeEmpty())

		Expect(Filter([]tt{
			{"one", true},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, matchFn)).To(Equal([]tt{
			{"one", true},
			{"two", true},
			{"five", true},
		}))
		Expect(FilterOut([]tt{
			{"one", true},
			{"two", true},
			{"three", false},
			{"four", false},
			{"five", true},
		}, matchFn)).To(Equal([]tt{
			{"three", false},
			{"four", false},
		}))
	})
}

func ExampleFilter() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Filter(n, func(n int) bool { return n > 0 }))
	fmt.Println(Filter(n, func(n int) bool { return n > 3 }))
	fmt.Println(Filter(n, func(n int) bool { return n < 0 }))
	// Output:
	// [1 2 3 4 5]
	// [4 5]
	// []
}

func ExampleFilterOut() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(FilterOut(n, func(n int) bool { return n > 0 }))
	fmt.Println(FilterOut(n, func(n int) bool { return n > 3 }))
	fmt.Println(FilterOut(n, func(n int) bool { return n < 0 }))
	// Output:
	// []
	// [1 2 3]
	// [1 2 3 4 5]
}

func TestOverlap(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		groupedTestCases := getBinaryOperationCases(id[int], id[int])

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(Overlap(tc.Left, tc.Right)).To(BeFalse())
						} else {
							Expect(Overlap(tc.Left, tc.Right)).To(BeTrue())
						}
					})
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		genFn := func(n int) string {
			return string([]byte{byte('a') + byte(n)})
		}

		groupedTestCases := getBinaryOperationCases(genFn, genFn)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(Overlap(tc.Left, tc.Right)).To(BeFalse())
						} else {
							Expect(Overlap(tc.Left, tc.Right)).To(BeTrue())
						}
					})
				}
			})
		}
	})
}

func ExampleOverlap() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Overlap(n, []int{9, 8, 7}))
	fmt.Println(Overlap(n, []int{1}))
	fmt.Println(Overlap(n, []int{1, 2, 3, 4, 5}))
	// Output:
	// false
	// true
	// true
}

func TestOverlapFunc(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		equalsInt := func(a, b int) bool { return a == b }
		groupedTestCases := getBinaryOperationCases(id[int], id[int])

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(OverlapFunc(tc.Left, tc.Right, equalsInt)).To(BeFalse())
						} else {
							Expect(OverlapFunc(tc.Left, tc.Right, equalsInt)).To(BeTrue())
						}
					})
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		equalsStr := func(a, b string) bool { return a == b }
		genFn := func(n int) string {
			return string([]byte{byte('a') + byte(n)})
		}
		groupedTestCases := getBinaryOperationCases(genFn, genFn)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(OverlapFunc(tc.Left, tc.Right, equalsStr)).To(BeFalse())
						} else {
							Expect(OverlapFunc(tc.Left, tc.Right, equalsStr)).To(BeTrue())
						}
					})
				}
			})
		}
	})

	t.Run("mix", func(t *testing.T) {
		RegisterTestingT(t)

		equalsFn := func(a int, b string) bool { return strconv.Itoa(a) == b }
		groupedTestCases := getBinaryOperationCases(id[int], strconv.Itoa)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(OverlapFunc(tc.Left, tc.Right, equalsFn)).To(BeFalse())
						} else {
							Expect(OverlapFunc(tc.Left, tc.Right, equalsFn)).To(BeTrue())
						}
					})
				}
			})
		}
	})
}

func ExampleOverlapFunc() {
	n := []int{1, 2, 3, 4, 5}

	equalsFn := func(a int, b string) bool { return strconv.Itoa(a) == b }

	fmt.Println(OverlapFunc(n, []string{"9", "8", "7"}, equalsFn))
	fmt.Println(OverlapFunc(n, []string{"1"}, equalsFn))
	fmt.Println(OverlapFunc(n, []string{"1", "2", "3", "4", "5"}, equalsFn))
	// Output:
	// false
	// true
	// true
}

func TestIntersect(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		groupedTestCases := getBinaryOperationCases(id[int], id[int])

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(Intersect(tc.Left, tc.Right)).To(BeNil())
						} else {
							Expect(Intersect(tc.Left, tc.Right)).To(Equal(tc.Intersect))
						}
					})
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		genFn := func(n int) string {
			return string([]byte{byte('a') + byte(n)})
		}

		groupedTestCases := getBinaryOperationCases(genFn, genFn)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(Intersect(tc.Left, tc.Right)).To(BeNil())
						} else {
							Expect(Intersect(tc.Left, tc.Right)).To(Equal(tc.Intersect))
						}
					})
				}
			})
		}
	})
}

func ExampleIntersect() {
	n := []int{1, 2, 3, 4, 5}

	fmt.Println(Intersect(n, []int{9, 8, 7}))
	fmt.Println(Intersect(n, []int{1}))
	fmt.Println(Intersect(n, []int{1, 2, 3, 4, 5, 6, 7}))
	// Output:
	// []
	// [1]
	// [1 2 3 4 5]
}

func TestIntersectFunc(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		equalsInt := func(a, b int) bool { return a == b }
		groupedTestCases := getBinaryOperationCases(id[int], id[int])

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(IntersectFunc(tc.Left, tc.Right, equalsInt)).To(BeNil())
						} else {
							Expect(IntersectFunc(tc.Left, tc.Right, equalsInt)).To(Equal(tc.Intersect))
						}
					})
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		equalsStr := func(a, b string) bool { return a == b }
		genFn := func(n int) string {
			return string([]byte{byte('a') + byte(n)})
		}
		groupedTestCases := getBinaryOperationCases(genFn, genFn)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(IntersectFunc(tc.Left, tc.Right, equalsStr)).To(BeNil())
						} else {
							Expect(IntersectFunc(tc.Left, tc.Right, equalsStr)).To(Equal(tc.Intersect))
						}
					})
				}
			})
		}
	})

	t.Run("mix", func(t *testing.T) {
		RegisterTestingT(t)

		equalsFn := func(a int, b string) bool { return strconv.Itoa(a) == b }
		groupedTestCases := getBinaryOperationCases(id[int], strconv.Itoa)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Intersect) == 0 {
							Expect(IntersectFunc(tc.Left, tc.Right, equalsFn)).To(BeNil())
						} else {
							Expect(IntersectFunc(tc.Left, tc.Right, equalsFn)).To(Equal(tc.Intersect))
						}
					})
				}
			})
		}
	})
}

func ExampleIntersectFunc() {
	n := []int{1, 2, 3, 4, 5}

	equalsFn := func(a int, b string) bool { return strconv.Itoa(a) == b }

	fmt.Println(IntersectFunc(n, []string{"9", "8", "7"}, equalsFn))
	fmt.Println(IntersectFunc(n, []string{"1"}, equalsFn))
	fmt.Println(IntersectFunc(n, []string{"1", "2", "3", "4", "5", "6", "7"}, equalsFn))
	// Output:
	// []
	// [1]
	// [1 2 3 4 5]
}

func TestDiff(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		groupedTestCases := getBinaryOperationCases(id[int], id[int])

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Diff) == 0 {
							Expect(Diff(tc.Left, tc.Right)).To(BeNil())
						} else {
							Expect(Diff(tc.Left, tc.Right)).To(Equal(tc.Diff))
						}
					})
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		genFn := func(n int) string {
			return string([]byte{byte('a') + byte(n)})
		}
		groupedTestCases := getBinaryOperationCases(genFn, genFn)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Diff) == 0 {
							Expect(Diff(tc.Left, tc.Right)).To(BeNil())
						} else {
							Expect(Diff(tc.Left, tc.Right)).To(Equal(tc.Diff))
						}
					})
				}
			})
		}
	})
}

func ExampleDiff() {
	n1 := []int{1, 2, 3, 4, 5}
	n2 := []int{2, 4, 6, 8}

	fmt.Println(Diff(n1, nil))
	fmt.Println(Diff(n1, n2))
	fmt.Println(Diff(nil, n2))
	// Output:
	// [1 2 3 4 5]
	// [1 3 5]
	// []
}

func TestDiffFunc(t *testing.T) {
	RegisterTestingT(t)

	t.Run("int", func(t *testing.T) {
		RegisterTestingT(t)

		equalsInt := func(a, b int) bool { return a == b }
		groupedTestCases := getBinaryOperationCases(id[int], id[int])

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Diff) == 0 {
							Expect(DiffFunc(tc.Left, tc.Right, equalsInt)).To(BeNil())
						} else {
							Expect(DiffFunc(tc.Left, tc.Right, equalsInt)).To(Equal(tc.Diff))
						}
					})
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		RegisterTestingT(t)

		equalsStr := func(a, b string) bool { return a == b }
		genFn := func(n int) string {
			return string([]byte{byte('a') + byte(n)})
		}
		groupedTestCases := getBinaryOperationCases(genFn, genFn)

		for groupName, testCases := range groupedTestCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Diff) == 0 {
							Expect(DiffFunc(tc.Left, tc.Right, equalsStr)).To(BeNil())
						} else {
							Expect(DiffFunc(tc.Left, tc.Right, equalsStr)).To(Equal(tc.Diff))
						}
					})
				}
			})
		}
	})

	t.Run("mix", func(t *testing.T) {
		RegisterTestingT(t)

		equalsFn := func(a int, b string) bool { return strconv.Itoa(a) == b }
		testCases := getBinaryOperationCases(id[int], strconv.Itoa)

		for groupName, testCases := range testCases {
			t.Run(groupName, func(t *testing.T) {
				RegisterTestingT(t)

				for _, tc := range testCases {
					t.Run(tc.Name(), func(t *testing.T) {
						RegisterTestingT(t)

						if len(tc.Diff) == 0 {
							Expect(DiffFunc(tc.Left, tc.Right, equalsFn)).To(BeNil())
						} else {
							Expect(DiffFunc(tc.Left, tc.Right, equalsFn)).To(Equal(tc.Diff))
						}
					})
				}
			})
		}
	})
}

func ExampleDiffFunc() {
	n1 := []int{1, 2, 3, 4, 5}
	n2 := []string{"2", "4", "6", "8"}

	equalsFn := func(a int, b string) bool { return strconv.Itoa(a) == b }

	fmt.Println(DiffFunc(n1, nil, equalsFn))
	fmt.Println(DiffFunc(n1, n2, equalsFn))
	fmt.Println(DiffFunc(nil, n2, equalsFn))
	// Output:
	// [1 2 3 4 5]
	// [1 3 5]
	// []
}

type binaryTestCase[T, S any] struct {
	Left      []T
	Right     []S
	Diff      []T
	Intersect []T
}

func (t binaryTestCase[T, S]) Name() string {
	aVal := "nil"
	if t.Left != nil {
		aVal = fmt.Sprintf("%v", t.Left)
	}

	bVal := "nil"
	if t.Right != nil {
		bVal = fmt.Sprintf("%v", t.Right)
	}

	return aVal + "-" + bVal
}

func getBinaryOperationCases[T, S any](
	leftGenerate func(int) T,
	rightGenerate func(int) S,
) map[string][]binaryTestCase[T, S] {
	testCases := getIntBinaryOperationCases()

	// Remap ints to T values
	res := make(map[string][]binaryTestCase[T, S])
	for name, cases := range testCases {
		res[name] = make([]binaryTestCase[T, S], len(cases))
		for i, c := range cases {
			res[name][i] = binaryTestCase[T, S]{
				Left:      Map(c.Left, leftGenerate),
				Right:     Map(c.Right, rightGenerate),
				Diff:      Map(c.Diff, leftGenerate),
				Intersect: Map(c.Intersect, leftGenerate),
			}
		}
	}
	return res
}

func getIntBinaryOperationCases() map[string][]binaryTestCase[int, int] {
	res := make(map[string][]binaryTestCase[int, int])

	res["empty"] = []binaryTestCase[int, int]{
		{nil, nil, nil, nil},
		{nil, []int{}, nil, nil},
		{[]int{}, nil, nil, nil},
		{[]int{}, []int{}, nil, nil},
	}

	res["leftEmpty"] = []binaryTestCase[int, int]{
		{nil, []int{1}, nil, nil},
		{nil, []int{1, 2, 3}, nil, nil},
		{[]int{}, []int{1}, nil, nil},
		{[]int{}, []int{1, 2, 3}, nil, nil},
	}

	res["rightEmpty"] = []binaryTestCase[int, int]{
		{[]int{1}, nil, []int{1}, nil},
		{[]int{1, 2, 3}, nil, []int{1, 2, 3}, nil},
		{[]int{1}, []int{}, []int{1}, nil},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}, nil},
	}

	res["combos"] = []binaryTestCase[int, int]{
		{[]int{1}, []int{1}, nil, []int{1}},
		{[]int{1}, []int{2}, []int{1}, nil},
		{[]int{1}, []int{1, 2, 3}, nil, []int{1}},
		{[]int{1}, []int{2, 3}, []int{1}, nil},

		{[]int{1, 2, 3}, []int{1}, []int{2, 3}, []int{1}},
		{[]int{1, 2, 3}, []int{2}, []int{1, 3}, []int{2}},
		{[]int{1, 2, 3}, []int{3}, []int{1, 2}, []int{3}},
		{[]int{1, 2, 3}, []int{4}, []int{1, 2, 3}, nil},

		{[]int{1, 2, 3}, []int{1, 2}, []int{3}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{2, 1}, []int{3}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 3}, []int{2}, []int{1, 3}},
		{[]int{1, 2, 3}, []int{3, 1}, []int{2}, []int{1, 3}},
		{[]int{1, 2, 3}, []int{2, 3}, []int{1}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{3, 2}, []int{1}, []int{2, 3}},
		{[]int{1, 2, 3}, []int{1, 4}, []int{2, 3}, []int{1}},
		{[]int{1, 2, 3}, []int{4, 1}, []int{2, 3}, []int{1}},

		{[]int{1, 2, 3}, []int{1, 2, 3}, nil, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{4, 1, 3}, []int{2}, []int{1, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}, nil},
	}

	res["dups"] = []binaryTestCase[int, int]{
		{[]int{1, 2, 3}, []int{1, 1, 1}, []int{2, 3}, []int{1}},
		{[]int{1, 2, 3}, []int{4, 4, 4}, []int{1, 2, 3}, nil},
		{[]int{1, 2, 3}, []int{4, 3, 4, 3, 4}, []int{1, 2}, []int{3}},

		{[]int{1, 2, 1, 3, 1}, []int{1}, []int{2, 3}, []int{1}},
		{[]int{1, 2, 2, 3, 1}, []int{2}, []int{1, 3, 1}, []int{2}},
		{[]int{1, 3, 2, 3, 3}, []int{3}, []int{1, 2}, []int{3}},
	}

	return res
}

func id[T any](v T) T { return v }
