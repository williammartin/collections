package linkedlist_test

import (
	"github.com/williammartin/collections/linkedlist"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linkedlist", func() {

	var empty *linkedlist.LinkedList
	var one *linkedlist.LinkedList
	var many *linkedlist.LinkedList

	BeforeEach(func() {
		empty = linkedlist.New()
		one = linkedlist.New()
		many = linkedlist.New()
		one.Add("1")
		many.Add("1")
		many.Add("2")
		many.Add("3")
	})

	Describe("LinkedList emptiness", func() {
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

	Describe("LinkedList size", func() {
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

	Describe("LinkedList containment", func() {
		Context("With zero elements", func() {
			It("Should not contain an element", func() {
				Expect(empty.Contains("1")).To(Equal(false))
			})
		})

		Context("With one element", func() {
			It("Should contain only the previously added element", func() {
				Expect(one.Contains("1")).To(Equal(true))
				Expect(one.Contains("2")).To(Equal(false))
			})
		})

		Context("With many elements", func() {
			It("Should contain all of but only the previously added elements", func() {
				Expect(many.Contains("1")).To(Equal(true))
				Expect(many.Contains("2")).To(Equal(true))
				Expect(many.Contains("3")).To(Equal(true))
				Expect(many.Contains("4")).To(Equal(false))
			})
		})
	})

	Describe("LinkedList ordering", func() {
		Context("With many elements", func() {
			It("Should be ordered as added", func() {
				first := many.Head()
				second := first.Next
				third := second.Next

				Expect(first.Value).To(Equal("1"))
				Expect(second.Value).To(Equal("2"))
				Expect(third.Value).To(Equal("3"))
			})
		})
	})
})
