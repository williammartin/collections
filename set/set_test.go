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

	BeforeEach(func() {
		empty = set.New()
		one = set.New()
		one.Add("1")
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
})
