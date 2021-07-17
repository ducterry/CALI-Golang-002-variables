package main

import "fmt"

/*
	- Session 002: Variables
	- Date: 18.07.2021
	- By: Nguyen Dang Duc
*/
func main() {
	/*
		TH1: Khai bao thong thuong
	*/
	var myStr string = "Nguyen Dang Duc"
	var myInt int = 100
	var myfloat float32 = 12.13
	fmt.Println("TH1: Khai bao thong thuong")
	fmt.Println(myStr, myInt, myfloat)
	fmt.Println("=====================")

	/*
		TH2: Khai bao tap chung
	*/
	var (
		employeeId          int    = 1991
		firstName, lastName string = "Duc", "Nguyen Dang"
	)
	fmt.Println(employeeId, firstName, lastName)
	fmt.Println("=====================")

	/*
		TH3: Khai bao ngan gon
	 */
	name:="Duc"
	age, salary,isProgramer :=20,18000000,true
	fmt.Println(name,age,salary,isProgramer)
}
