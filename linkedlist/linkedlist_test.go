package linkedlist_test

import (
	"github.com/williammartin/collections/linkedlist"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Linkedlist", func() {

	var empty *linkedlist.LinkedList
	var one *linkedlist.LinkedList

	BeforeEach(func() {
		empty = linkedlist.New()
		one = linkedlist.New()
		one.Add("1")
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
})
