package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/dejarc/minimum-routing-cost-path"
)

var _ = Describe("main.go", func() {
	It("should convert strings to Loads", func() {
		Loads := map[int]Load{
			1: {1, Point{-50.1, 80.0}, Point{90.1, 12.2}, 155.7333618721435, false},
			2: {2, Point{-24.5, -19.2}, Point{98.5, 0}, 124.48951763100378, false},
			3: {3, Point{0.3, 8.9}, Point{40.9, 55}, 61.42939035998974, false},
			4: {4, Point{5.3, -61.1}, Point{77.8, -5.4}, 91.42614505708966, false},
		}
		sampleLines := []string{
			"1 (-50.1,80.0) (90.1,12.2)",
			"2 (-24.5,-19.2) (98.5,1,8)",
			"3 (0.3,8.9) (40.9,55.0)",
			"4 (5.3,-61.1) (77.8,-5.4)",
		}
		result := ConvertStringsToLoads(sampleLines)
		Expect(result).To(Equal(Loads))
	})
	It("should get distance to Point of origin", func() {
		prevDis := 0
		cur := Point{X: 0, Y: 0}
		nextLoad := Load{1, Point{X: -50.1, Y: 80}, Point{X: 90.1, Y: 12.2}, 155.7333618721435, false}
		expected := 341.0484306841512
		result := GetDistanceToHome(float64(prevDis), cur, nextLoad)
		Expect(result).To(Equal(expected))
	})
	It("should find a Load that was previously visited invalid", func() {
		testLoad := Load{1, Point{-50.1, 80.0}, Point{90.1, 12.2}, 155.7333618721435, true}
		Expect(IsValid(0, Point{0, 0}, testLoad)).To(Equal(false))
	})
	It("should find a Load with a distance violation invalid", func() {
		testLoad := Load{1, Point{-50.1, 80.0}, Point{90.1, 12.2}, 155.7333618721435, false}
		prevDistance := 500.000
		Expect(IsValid(prevDistance, Point{-1.99, 2.99}, testLoad)).To(Equal(false))
	})
	It("should convert a list of drivers to a list of formatted strings", func() {
		Loads := []Driver{
			{[]int{1, 2, 3, 4}, 455},
			{[]int{5, 6, 7, 8}, 455},
			{[]int{9, 10, 11, 12}, 455},
		}
		expected := []string{
			"[1, 2, 3, 4]",
			"[5, 6, 7, 8]",
			"[9, 10, 11, 12]",
		}
		result := LoadsToString(Loads)
		Expect(result).To(Equal(expected))
	})
})
