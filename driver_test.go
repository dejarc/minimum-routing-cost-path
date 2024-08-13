package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dejarc/minimum-routing-cost-path"
)

var _ = Describe("Driver", func() {
	It("should create a driver", func() {
		Expect(CreateDriver()).To(Equal(Driver{make([]int, 0), 0}))
	})
})
