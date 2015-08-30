package set_test

import (
	"github.com/williammartin/collections/set"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"
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

	fmt.Println(empty)

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
})
