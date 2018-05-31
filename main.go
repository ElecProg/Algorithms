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

	fmt.Println("Heap sort")
	arr = []int{8, 3, 2, 4, 6}
	sorting.HeapSort(arr)
	fmt.Println(arr)
	fmt.Println()

	fmt.Println("Counting sort")
	strngs := []string{"c", "b", "b"}
	stringsorting.CountingSort(strngs, 0)
	fmt.Println(strngs)
	fmt.Println()

	fmt.Println("LSD sort")
	strngs = []string{"Bob", "Bart", "Bar", "Charlie"}
	stringsorting.LSDSort(strngs, 4)
	fmt.Println(strngs)
	fmt.Println()

	fmt.Println("MSD sort")
	strngs = []string{"Bob", "Bart", "Bar", "Charlie", "Charlia"}
	stringsorting.MSDSort(strngs)
	fmt.Println(strngs)
	fmt.Println()

	fmt.Println("3-way string quick sort")
	strngs = []string{"Bob", "Bart", "Bar", "Charlie", "Charlia"}
	stringsorting.ThreeWayQuickSort(strngs)
	fmt.Println(strngs)
	fmt.Println()
}
