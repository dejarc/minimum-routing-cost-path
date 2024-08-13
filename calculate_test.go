package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dejarc/minimum-routing-cost-path"
)

var _ = Describe("calculate.go", func() {
	It("should calculate distance correctly", func ()  {
		var expected float64 = 141.4371954377631
		start := Point{0, 0}
		end := Point{-100.0101, -100.0123}
		Expect(GetDist(start, end)).To(Equal(expected))
	})
	It("should calculate total cost correctly", func ()  {
		drivers := []Driver{
			{[]int{1, 2, 3, 4}, 455},
			{[]int{5, 6, 7, 8}, 455},
			{[]int{9, 10, 11, 12}, 455},
		}
		var expected float64 = 2865
		Expect(CalculateTotalCost(drivers)).To(Equal(expected))
	})
})
