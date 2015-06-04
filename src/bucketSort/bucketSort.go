package main

import(
	"fmt"
	"math/rand"
	"time"
)

func main(){
	inSlice := createRandSlice(100)
	sortedSlice := bucketSort(inSlice)
	fmt.Println("-------beforeSort-------")
	fmt.Println(inSlice)
	fmt.Println("-------afterSort-------")
	fmt.Print(sortedSlice)

}

func bucketSort(inSlice []int) []int{
	var target []int

	bucket := make([][]int, 100)
	for i:=0; i<len(inSlice); i++{
		bucket[inSlice[i]] = append(bucket[inSlice[i]], inSlice[i])
	}

	for i:=0; i<100; i++{
		if len(bucket[i]) == 0{
			continue
		}

		for k:=0; k<len(bucket[i]); k++{
			target = append(target, bucket[i][k])
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