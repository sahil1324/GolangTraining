package main
import "fmt"
func slice() {
list:=[]int{1,2,3}
list = append(list,4)
list = append(list,5)
fmt.Println(list)
list1:=list
list1=append(list1,6)
fmt.Println(list1)
}
func main() {
	slice()
}
