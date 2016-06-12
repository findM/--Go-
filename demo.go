package main

import "fmt"

//插入排序
func insertion_sort(unsorted []int) {
	fmt.Println("排序前")
	fmt.Println(unsorted)

	for i := 1; i < len(unsorted); i++ {
		key := unsorted[i]

		var j int

		for j = i - 1; j >= 0; j-- {
			if key < unsorted[j] { //查找合适的位置
				unsorted[j+1] = unsorted[j]
			} else {
				break
			}
		}
		unsorted[j+1] = key
	}

	fmt.Println("排序后")
	fmt.Println(unsorted)
}

//桶排序或者箱排序
func bucket_sort(unsorted []int, max int) {
	fmt.Println("排序前")
	fmt.Println(unsorted)

	a := make([]int, max+1)

	for i := 0; i < len(unsorted); i++ {
		key := unsorted[i]
		a[key]++
	}

	index := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < a[i]; j++ {
			unsorted[index] = i
			index++
		}
	}

	fmt.Println("排序后")
	fmt.Println(unsorted)
}
