package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dejarc/minimum-routing-cost-path"
)

var _ = Describe("Load", func() {
	It("should create a load", func() {
		id := 1
		var startX float64 = 0
		var startY float64 = 0
		var endX float64 = 100
		var endY float64 = -100
		expected := Load{
			1,
			Point{X: 0, Y: 0},
			Point{X: 100, Y: -100},
			141.4213562373095,
			false,
		}
		Expect(CreateLoad(id, startX, startY, endX, endY)).To(Equal(expected))
	})
})
