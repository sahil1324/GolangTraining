package main
import "fmt"
func map1() {
dict := map[int]string{1:"virat kohli"}
dict[2]="rohit sharma"
fmt.Println(dict)
dict[1]="KL rahul"
fmt.Println(dict)
delete(dict,1)
fmt.Println(dict)
}
func main() {
	map1()
}
