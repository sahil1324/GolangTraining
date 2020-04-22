package main
import (
	"fmt"
)
type userPrivilage interface{
	 deposit(amount float64) 
	 withdraw(amount float64) 
	 checkBalance() float64
}
type currentAccount struct{
	balance float64
}
type savingAccount struct{
	balance float64
	interestRate float64
}
func (ca currentAccount) deposit(amount float64){
	ca.balance += amount
}
func (ca currentAccount) withdraw(amount float64) {
  ca.balance-=amount
}
func (ca currentAccount) checkBalance() float64 {
  return ca.balance
}
func (sa savingAccount) deposit(amount float64) {
	sa.balance+=(amount+sa.interestRate*amount)
}
func (sa savingAccount) withdraw(amount float64) {
  sa.balance-=amount
}
func (sa savingAccount) checkBalance() float64 {
  return sa.balance
}
func accountDetail(u userPrivilage){
	fmt.Println("Deposit Money")
	u.deposit(100)
	u.withdraw(10) 
	fmt.Println(u.checkBalance())
}
func main(){
cA:=currentAccount{balance:0}
sA:=savingAccount{balance:0,interestRate:10}
fmt.Println("Balance in current account")
accountDetail(cA)
fmt.Println("Balance in saving account")
accountDetail(sA)

}
