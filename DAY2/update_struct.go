package main

import "fmt"

type student struct {
	first_name string
	last_name  string
	contact_no contact
	course     string
	marks      []int
}
type contact struct {
	phone_number int
}

func main() {
	info := student{first_name: "sahil",
		last_name:  "dhiman",
		contact_no: contact{phone_number: 123},
		course:     "cse",
		marks:      []int{10, 10}}
	fmt.Println("Before Update")
	print(info.contact_no.phone_number)
	fmt.Println("\nAfter Update")
	ptr_info := &info
	ptr_info.contact_no.phone_number = 456
	print(info.contact_no.phone_number)
}
func (stu student) print() {

	fmt.Println(stu)
}
