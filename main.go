package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(narcissistic(153))
	fmt.Println(narcissistic(111))

	fmt.Println(findOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(findOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	// will return 0 means false
	fmt.Println(findOutlier([]int{11, 13, 15, 19, 9, 13, -21}))

	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	needle := "blue"

	i := findNeedle(haystack, needle)
	if i != -1 {
		fmt.Printf("The needle '%s' is found at %d\n", needle, i)
	} else {
		fmt.Printf("The needle '%s' was not found in the haystack")
	}

	fmt.Println(blueOceanReverse([]int{1, 2, 3, 4, 6, 10}, []int{1}))
	fmt.Println(blueOceanReverse([]int{1, 5, 5, 5, 5, 3}, []int{5}))

}

func narcissistic(number int) bool {
	numStr := strconv.Itoa(number)
	numDigits := len(numStr)

	sum := 0
	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		sum += int(math.Pow(float64(digit), float64(numDigits)))
	}

	return sum == number
}

func findOutlier(integers []int) int {
	ec, oc := 0, 0
	var even, odd int

	for _, num := range integers {
		if num%2 == 0 {
			ec++
			even = num
		} else {
			oc++
			odd = num
		}

		if ec > 1 && oc == 1 {
			return odd
		} else if oc > 1 && ec == 1 {
			return even
		}
	}

	return 0
}

func findNeedle(haystack []string, needle string) int {

	for i, str := range haystack {
		if str == needle {
			return i
		}
	}
	// i will return -1 if i cant find the needle in the haystack
	return -1
}

func blueOceanReverse(list1, list2 []int) []int {
	result := []int{}

	list2Map := make(map[int]bool)
	for _, num := range list2 {
		list2Map[num] = true
	}

	for _, num := range list1 {
		if !list2Map[num] {
			result = append(result, num)
		}
	}

	return result
}
