package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums [3]int

	fmt.Scan(&nums[0])
	fmt.Scan(&nums[1])
	fmt.Scan(&nums[2])

	sort.Ints(nums[:])

	fmt.Println(nums[0], nums[1], nums[2])
}
