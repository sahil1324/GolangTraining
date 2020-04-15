package main
import "fmt"

func factorial(num int){
     fact := 1
     for i:=1;i<=num; i++ {
       	 fact = fact * i
     }
	 fmt.Println(fact)
}
func main() {
    num := 6
    factorial(num)
}
