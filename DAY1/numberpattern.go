package main
import "fmt"

func pattern(row int){
     sum:= 1
     for i:=1;i<=row; i++ {
         for j:=0; j<i; j++{
             fmt.Print(sum," ")
			 sum++
         }
		 fmt.Println()
     }
}
func main() {
    row := 5
    pattern(row)
}
