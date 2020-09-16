package math_test

import (
	. "github.com/muxxer/iota.go/math"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {

	Context("AbsInt64()", func() {
		It("should only return positive values", func() {
			v := AbsInt64(-9223372036854775807)
			Expect(v).To(Equal(int64(9223372036854775807)))
		})
	})
})
