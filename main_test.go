package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("main.go", func() {
	It("should convert strings to loads", func ()  {
		loads := map[int]load{
			1: {1, point{-50.1, 80.0}, point{90.1, 12.2}, 155.7333618721435, false},
			2: {2, point{-24.5, -19.2}, point{98.5, 0}, 124.48951763100378, false},
			3: {3, point{0.3, 8.9}, point{40.9, 55}, 61.42939035998974, false},
			4: {4, point{5.3, -61.1}, point{77.8, -5.4}, 91.42614505708966, false},
		}
		sampleLines := []string{
			"1 (-50.1,80.0) (90.1,12.2)",
			"2 (-24.5,-19.2) (98.5,1,8)",
			"3 (0.3,8.9) (40.9,55.0)",
			"4 (5.3,-61.1) (77.8,-5.4)",
		}
		result := convertStringsToLoads(sampleLines)
		Expect(result).To(Equal(loads))
	})
	It("should get distance to point of origin", func ()  {
		prevDis := 0
		cur := point{x: 0, y: 0}
		nextLoad:= load{ 1,  point {x: -50.1, y: 80}, point {x: 90.1, y: 12.2},  155.7333618721435, false}
		expected := 341.0484306841512
		result := getDistanceToHome(float64(prevDis), cur, nextLoad)
		Expect(result).To(Equal(expected))
	})
	It("should find a load that was previously visited invalid", func() {
		testLoad := load{1, point{-50.1, 80.0}, point{90.1, 12.2}, 155.7333618721435, true}
		Expect(isValid(0, point{0,0}, testLoad)).To(Equal(false))
	})
	It("should find a load with a distance violation invalid", func() {
		testLoad := load{1, point{-50.1, 80.0}, point{90.1, 12.2}, 155.7333618721435, false}
		prevDistance := 500.000
		Expect(isValid(prevDistance, point{-1.99,2.99}, testLoad)).To(Equal(false))
	})
	It("should convert a list of drivers to a list of formatted strings", func() {
		loads := []driver{
			{[]int{1,2,3,4}, 455},
			{[]int{5,6,7,8}, 455},
			{[]int{9,10,11, 12}, 455},
		}
		expected := []string{
			"[1, 2, 3, 4]",
			"[5, 6, 7, 8]",
			"[9, 10, 11, 12]",
		}
		result := loadsToString(loads)
		Expect(result).To(Equal(expected))
	})
})

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "main.go")
}