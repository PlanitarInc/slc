package slc

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestIncludes(t *testing.T) {
	RegisterTestingT(t)

	Expect(Includes(nil, 1)).To(BeFalse())
	Expect(Includes(nil, 0)).To(BeFalse())
	Expect(Includes([]int{}, 9)).To(BeFalse())
	Expect(Includes([]int{1, 2, 3, 4, 5}, 9)).To(BeFalse())
	Expect(Includes([]int{1, 2, 3, 4, 5}, 1)).To(BeTrue())
	Expect(Includes([]int{1, 2, 3, 4, 5}, 5)).To(BeTrue())
	Expect(Includes([]int{1, 2, 3, 4, 5}, 4)).To(BeTrue())

	Expect(Includes(nil, "")).To(BeFalse())
	Expect(Includes(nil, "zz")).To(BeFalse())
	Expect(Includes([]string{}, "")).To(BeFalse())
	Expect(Includes([]string{"a", "b", "c"}, "zz")).To(BeFalse())
	Expect(Includes([]string{"a", "b", "c"}, "a")).To(BeTrue())
	Expect(Includes([]string{"a", "b", "c"}, "b")).To(BeTrue())
	Expect(Includes([]string{"a", "b", "c"}, "c")).To(BeTrue())
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

func TestDiff(t *testing.T) {
	RegisterTestingT(t)

	t.Run("empty", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Diff[int](nil, nil)).To(BeNil())
		Expect(Diff([]int{}, nil)).To(BeNil())
		Expect(Diff(nil, []int{})).To(BeNil())
		Expect(Diff([]int{}, []int{})).To(BeNil())
	})

	t.Run("leftEmpty", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Diff(nil, []int{1})).To(BeNil())
		Expect(Diff(nil, []int{1, 2, 3})).To(BeNil())
		Expect(Diff([]int{}, []int{1})).To(BeNil())
		Expect(Diff([]int{}, []int{1, 2, 3})).To(BeNil())
	})

	t.Run("rightEmpty", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Diff([]int{1}, nil)).To(Equal([]int{1}))
		Expect(Diff([]int{1, 2, 3}, nil)).To(Equal([]int{1, 2, 3}))
		Expect(Diff([]int{1}, []int{})).To(Equal([]int{1}))
		Expect(Diff([]int{1, 2, 3}, []int{})).To(Equal([]int{1, 2, 3}))
	})

	t.Run("combos", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Diff([]int{1}, []int{1})).To(BeNil())
		Expect(Diff([]int{1}, []int{2})).To(Equal([]int{1}))
		Expect(Diff([]int{1}, []int{1, 2, 3})).To(BeNil())
		Expect(Diff([]int{1}, []int{2, 3})).To(Equal([]int{1}))

		Expect(Diff([]int{1, 2, 3}, []int{1})).To(Equal([]int{2, 3}))
		Expect(Diff([]int{1, 2, 3}, []int{2})).To(Equal([]int{1, 3}))
		Expect(Diff([]int{1, 2, 3}, []int{3})).To(Equal([]int{1, 2}))
		Expect(Diff([]int{1, 2, 3}, []int{4})).To(Equal([]int{1, 2, 3}))

		Expect(Diff([]int{1, 2, 3}, []int{1, 2})).To(Equal([]int{3}))
		Expect(Diff([]int{1, 2, 3}, []int{2, 1})).To(Equal([]int{3}))
		Expect(Diff([]int{1, 2, 3}, []int{1, 3})).To(Equal([]int{2}))
		Expect(Diff([]int{1, 2, 3}, []int{3, 1})).To(Equal([]int{2}))
		Expect(Diff([]int{1, 2, 3}, []int{2, 3})).To(Equal([]int{1}))
		Expect(Diff([]int{1, 2, 3}, []int{3, 2})).To(Equal([]int{1}))
		Expect(Diff([]int{1, 2, 3}, []int{1, 4})).To(Equal([]int{2, 3}))
		Expect(Diff([]int{1, 2, 3}, []int{4, 1})).To(Equal([]int{2, 3}))

		Expect(Diff([]int{1, 2, 3}, []int{1, 2, 3})).To(BeNil())
		Expect(Diff([]int{1, 2, 3}, []int{4, 1, 3})).To(Equal([]int{2}))
		Expect(Diff([]int{1, 2, 3}, []int{4, 5, 6})).To(Equal([]int{1, 2, 3}))
	})

	t.Run("dups", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Diff([]int{1, 2, 3}, []int{1, 1, 1})).To(Equal([]int{2, 3}))
		Expect(Diff([]int{1, 2, 3}, []int{4, 4, 4})).To(Equal([]int{1, 2, 3}))
		Expect(Diff([]int{1, 2, 3}, []int{4, 3, 4, 3, 4})).To(Equal([]int{1, 2}))

		Expect(Diff([]int{1, 2, 1, 3, 1}, []int{1})).To(Equal([]int{2, 3}))
		Expect(Diff([]int{1, 2, 2, 3, 1}, []int{2})).To(Equal([]int{1, 3, 1}))
		Expect(Diff([]int{1, 3, 2, 3, 3}, []int{3})).To(Equal([]int{1, 2}))
	})

	t.Run("strings", func(t *testing.T) {
		RegisterTestingT(t)

		Expect(Diff([]string{"a", "b", "c"}, []string{"a"})).To(Equal([]string{"b", "c"}))
		Expect(Diff([]string{"a", "b", "c"}, []string{"b", "d"})).To(Equal([]string{"a", "c"}))
		Expect(Diff([]string{"a", "b", "c"}, []string{"d", "b", "d"})).To(Equal([]string{"a", "c"}))
	})
}
