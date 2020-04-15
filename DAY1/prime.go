package main
import "fmt"

func prime(num int){
     prime := 0
	 for i:=1;i<=num; i++ {
       	 if (num%i)==0{
		 prime++
		 }
     }
     if prime == 2 {
	 fmt.Println("Prime Number")
	 }else{
	 fmt.Println("Not Prime Number")
}}
func main() {
    num := 11
    prime(num)
}
