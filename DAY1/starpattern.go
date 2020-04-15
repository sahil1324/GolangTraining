package main
import "fmt"

func pattern(row int){
     for i:=1;i<=row; i++ {
         for j:=0; j<i; j++{
             fmt.Print("*")
         }
		 fmt.Println()
     }
}
func main() {
    row := 5
    pattern(row)
}
