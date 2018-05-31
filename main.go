package main

import (
	"fmt"

	"github.com/ElecProg/Algorithms/sorting"
)

func main() {
	fmt.Println("Selection sort")
	arr := []int{8, 3, 2, 4}
	sorting.SelectionSort(arr)
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Insertion sort")
	arr = []int{8, 3, 2, 4}
	sorting.InsertionSort(arr)
	fmt.Println(arr)
	fmt.Println()
}
