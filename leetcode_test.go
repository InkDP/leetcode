package main

import (
	"fmt"
	"testing"
)

func Test204(t *testing.T) {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes2(10))
	fmt.Println(countPrimes3(10))
}

func Test303(t *testing.T){
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	param_1 := obj.SumRange(0, 5)
	fmt.Println(param_1)
	obj = Constructor2([]int{-2, 0, 3, -5, 2, -1})
	param_1 = obj.SumRange2(0, 5)
	fmt.Println(param_1)
}

func Test349(t *testing.T) {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9,4,9,8,4}))
	fmt.Println(intersection2([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection2([]int{4, 9, 5}, []int{9,4,9,8,4}))
}
