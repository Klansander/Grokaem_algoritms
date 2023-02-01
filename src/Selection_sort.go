package main
 import (
	"fmt"
 )

 func find_Smallest(array []int) (index int) {
	min:= array[0]
	for i,value := range array {
		if min>value {
			min=value
			index=i
		}
	}
	return
 }
 func selectionSort ( array []int) []int {
	array_new :=make([]int,len(array))
	ind:=0
	for i :=range array {
		ind=find_Smallest(array)
		array_new[i]=array[ind]
		array=append(array[:ind],array[ind+1:]...)
	}
	return array_new
 }
 func main () {
	array:=[]int { 2, 3, 1, 4, 6, 43, 23, 6, 7, 4, 6, 86, 3}
	fmt.Println("Исходный массив: ",array)
	fmt.Println("Отсортированный массив: ",  selectionSort(array))
 }