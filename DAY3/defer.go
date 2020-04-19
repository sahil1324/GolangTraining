package main
import "fmt"
func func1() {
	fmt.Println("func1")
}
func func2() string{
    defer func1()
	return "func2"
}
func main() {
	fmt.Println(func2())
}
