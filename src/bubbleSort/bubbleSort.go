package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	unSortedSlice := createRandSlice(100)
	sortedSlice := bubbleSort(unSortedSlice)
	fmt.Println("----Before sort----")
	fmt.Println(unSortedSlice)
	fmt.Println("----After sort----")
	fmt.Println(sortedSlice)
}


func bubbleSort(inSlice []int) []int{
	//ここはまった sliceはポインタ
	target := make([]int, len(inSlice))
	copy(target, inSlice)
	for i:=0; i<len(target)-1; i++{
		var tmp int
		for j:=len(target)-1; j>i; j--{
			if(target[j] < target[j-1]) {
				tmp = target[j]
				target[j] = target[j-1]
				target[j-1] = tmp
			}
		}
	}
	return target
}

func createRandSlice(num int) []int{
	var slice []int
	//Seedを時間で変えることで毎回違う乱数列を発生させる。
	//Seedを行わない場合は、暗黙にSeed(1)として同じ乱数列を発生させる。
	rand.Seed(time.Now().UnixNano())
	for i:=0; i < 100; i++{
		slice = append(slice, rand.Intn(100))
	}
	return slice
}

