package main
import "fmt"
func sum (array []int) int {
	s:=0
	if (len(array)>1) {
		s+=array[0]
		return s+sum(array[1:])
	} else {
		return array[0]
	}
	
}
func main() {
	array:=[]int {2,3,4,5,6}
	fmt.Println("Исходный массив: ",array)
	fmt.Println("Сумма массива: ",sum(array))
}