package main

import (
	"fmt"

	"github.com/ElecProg/Algorithms/sorting"
	"github.com/ElecProg/Algorithms/stringsorting"
)

func main() {
	fmt.Println("Selection sort")
	arr := []int{8, 3, 2, 4}
	sorting.SelectionSort(arr)
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Insertion sort")
	arr = []int{8, 3, 2, 4, 6}
	sorting.InsertionSort(arr)
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Merge sort")
	arr = []int{8, 3, 2, 4, 6}
	sorting.MergeSort(arr)
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Quick sort")
	arr = []int{8, 3, 2, 4, 6}
	sorting.QuickSort(arr)
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Counting sort")
	barr := [][]byte{[]byte{'c'}, []byte{'b'}, []byte{'b'}}
	stringsorting.CountingSort(barr, 0)
	fmt.Println(barr)
	fmt.Println()
}
