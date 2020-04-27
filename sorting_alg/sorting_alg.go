package sorting_alg

import (
        "math/rand"
)
func SelectionSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }
        min := 0
        for i := 0; i < len(arr)-1; i++ {
                min = i
                for j := i + 1; j < len(arr); j++ {
                        if arr[min] > arr[j] {
                                arr[min], arr[j] = arr[j], arr[min]
                        }
                }
        }
        return arr
}

func BubbleSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }
        for i := len(arr); i > 0; i-- {
                for j := 0; j < len(arr[:i])-1; j++ {
                        if arr[j] > arr[j+1] {
                                arr[j], arr[j+1] = arr[j+1], arr[j]
                        }
                }
        }
        return arr
}

func InsertionSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }
        for i := 1; i < len(arr); i++ {
                for j := 0; j < i; j++ {
                        if arr[j] > arr[i] {
                                arr[j], arr[i] = arr[i], arr[j]
                        }
                }
        }
        return arr
}

func MergeSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }
        left := MergeSort(arr[:len(arr)/2])
        right := MergeSort(arr[len(arr)/2:])
        i, j := 0, 0
        rst := make([]int, 0)
        for i < len(left) || j < len(right) {
                if i >= len(left) {
                        rst = append(rst, right[j])
                        j++
                } else if j >= len(right) {
                        rst = append(rst, left[i])
                        i++
                } else {
                        if left[i] < right[j] {
                                rst = append(rst, left[i])
                                i++
                        } else {
                                rst = append(rst, right[j])
                                j++
                        }
                }
        }
        return rst
}

func QuickSort(arr []int) []int {
        if len(arr) <= 1 {
                return arr
        }

        p := len(arr) - 1
        i := 0
        j := p - 1

        for i <= j {
                if arr[i] < arr[p] {
                        //fmt.Println("arr[i] < arr[p]", arr, i, j, p)
                        i++
                } else {
                        if arr[j] > arr[p] {
                                //fmt.Println("arr[i] >= arr[p] & arr[j] > arr[p]", arr, i, j, p)
                                j--
                        } else {
                                arr[i], arr[j] = arr[j], arr[i]
                                //fmt.Println("arr[i] >= arr[p] & arr[j] <= arr[p]", arr, i, j, p)
                                i++
                                j--
                        }
                }
        }
        arr[p], arr[i] = arr[i], arr[p]

        QuickSort(arr[:i])
        QuickSort(arr[i+1:])

        return arr
}

func QuickSort2(a []int) []int {

        if len(a) < 2 {
                return a
        }

        left, right := 0, len(a)-1

        pivot := rand.Int() % len(a)

        a[pivot], a[right] = a[right], a[pivot]

        for i, _ := range a {
                if a[i] < a[right] {
                        a[left], a[i] = a[i], a[left]
                        left++
                }
        }

        a[left], a[right] = a[right], a[left]

        QuickSort2(a[:left])
        QuickSort2(a[left+1:])

        return a
}