package cap

import "fmt"

func ExampleCap() {
	nums := []int{1,2,3,4,5}

	// nums, len, cap
	fmt.Println(nums)
	fmt.Println(len(nums), cap(nums))

	sliced1 := nums[:3]
	// sliced1, len, cap of sliced1
	fmt.Println(sliced1)
	fmt.Println(len(sliced1), cap(sliced1))

	sliced2 := nums[2:]
	// sliced2, len, cap of sliced2
	fmt.Println(sliced2)
	fmt.Println(len(sliced2), cap(sliced2))

	sliced3 := sliced1[:4]
	// sliced3, len, cap of sliced3
	fmt.Println(sliced3)
	fmt.Println(len(sliced3), cap(sliced3))

	nums[2] = 100
	// nums, sliced1, sliced2, sliced3
	fmt.Println(nums)
	fmt.Println(sliced1)
	fmt.Println(sliced2)
	fmt.Println(sliced3)

	// Output:
	// [1 2 3 4 5]
	// 5 5
	// [1 2 3]
	// 3 5
	// [3 4 5]
	// 3 3
	// [1 2 3 4]
	// 4 5
	// [1 2 100 4 5]
	// [1 2 100]
	// [100 4 5]
	// [1 2 100 4]
}
