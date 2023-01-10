package data_structure

import "fmt"

func InitArray() {
	var arr1 [5]int
	fmt.Printf("%v\n", arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", arr2)

	arr3 := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5},
	}
	fmt.Printf("%v\n", arr3)
}

func InitSlice() {
	var slice []int
	slice = append(slice, 0, 1, 2, 3, 4, 5)
	fmt.Println(slice)

	slice = append(slice[0:2], slice[3:]...)
	fmt.Println(slice)
}
