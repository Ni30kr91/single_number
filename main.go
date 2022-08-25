// single number
// Example 1:
// Input: nums = [2,2,1]
// Output: 1

// Example 2:
// Input: nums = [4,1,2,1,2]
// Output: 4

package main

import "fmt"

func main() {
	arr := []int{1, 0, 1}
	fmt.Println("elements of array")
	//temp := 0
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			fmt.Println("i", arr[i])
			fmt.Println("j", arr[j])
			if arr[i] == arr[j] && i != j {
				count++
				fmt.Println("count: ", "dggf")
			}
			fmt.Println("count: ", count)
		}

		fmt.Println("count: ", count)
		if count == 0 {
			fmt.Println(arr[i])
			break
		}
	}

}
