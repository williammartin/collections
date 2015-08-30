package set_test

import (
	"github.com/williammartin/collections/set"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {

	var empty *set.Set
	var one *set.Set
	var many *set.Set

	BeforeEach(func() {
		empty = set.New()
		one = set.New()
		many = set.New()
		one.Add("1")
		many.Add("1")
		many.Add("2")
		many.Add("3")
	})

	Describe("Set emptiness", func() {
		Context("With zero elements", func() {
			It("should be empty", func() {
				Expect(empty.IsEmpty()).To(Equal(true))
			})
		})

		Context("With one element", func() {
			It("should not be empty", func() {
				Expect(one.IsEmpty()).To(Equal(false))
			})
		})
	})

	Describe("Set size", func() {
		Context("With zero elements", func() {
			It("should have a size of zero", func() {
				Expect(empty.Size()).To(Equal(0))
			})
		})

		Context("With one element", func() {
			It("should have a size of one", func() {
				Expect(one.Size()).To(Equal(1))
			})
		})

		Context("With many elements", func() {
			It("should have a size of three", func() {
				Expect(many.Size()).To(Equal(3))
			})
		})
	})

	Describe("Set containment", func() {
		Context("With zero elements", func() {
			It("should not contain any elements", func() {
				Expect(empty.Contains("1")).To(Equal(false))
			})
		})

		Context("With elements added", func() {
			It("should contain those elements added", func() {
				Expect(one.Contains("1")).To(Equal(true))
				Expect(many.Contains("1")).To(Equal(true))
				Expect(many.Contains("2")).To(Equal(true))
				Expect(many.Contains("3")).To(Equal(true))
			})

			It("should not contain elements not added", func() {
				Expect(one.Contains("2")).To(Equal(false))
				Expect(many.Contains("5")).To(Equal(false))
			})
		})
	})

	Describe("Element removal", func() {
		Context("With one element", func() {
			It("should remove that element from the set", func() {
				setUnderTest := set.New()
				setUnderTest.Add("1")
				setUnderTest.Remove("1")
				Expect(setUnderTest.Contains("1")).To(Equal(false))
			})
		})

		Context("With many elements", func() {
			It("should remove many elements from the set", func() {
				setUnderTest := set.New()
				setUnderTest.Add("1")
				setUnderTest.Add("2")
				setUnderTest.Add("3")
				setUnderTest.Remove("1")
				setUnderTest.Remove("3")

				Expect(setUnderTest.Contains("1")).To(Equal(false))
				Expect(setUnderTest.Contains("2")).To(Equal(true))
				Expect(setUnderTest.Contains("3")).To(Equal(false))
			})
		})
	})

	Describe("Element addition", func() {
		Context("With an element added", func() {
			It("should not add an element with the same value", func() {
				one.Add("1")
				Expect(one.Size()).To(Equal(1))
			})
		})
	})
})
