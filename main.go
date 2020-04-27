package main

import (
	"fmt"
	sorting "github.com/gregoriusdev/golang_dsnalg/sorting_alg"
	log "github.com/sirupsen/logrus"
)

func main() {

	arr := []int{10, 30, 27, 5, 60, 4, 1, 90, 61, 3}
    fmt.Println("before : ", arr)
    var sortedArr []int
    /*
    sortedArr = sorting.MergeSort(arr)
    fmt.Println("after merge sort : ", sortedArr)
    sortedArr = sorting.SelectionSort(arr)
    fmt.Println("after selection sort : ", sortedArr)
    sortedArr = sorting.BubbleSort(arr)
    fmt.Println("after bubble sort : ", sortedArr)
    sortedArr = sorting.InsertionSort(arr)
    fmt.Println("after insertion sort : ", sortedArr)
    */
    sortedArr = sorting.QuickSort(arr)
	fmt.Println("after quick sort: ", sortedArr)
	log.Info("sorting finished")
}