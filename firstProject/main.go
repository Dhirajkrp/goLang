package main

import (
	"fmt"
	"goLang/firstProject/lib"
)

func main(){
	fmt.Println(lib.Concat("dhiraj" , "Prasad"))
	a , b := 10,3
	coef ,rem := lib.Divide(a,b)


	fmt.Printf("The coeffecient and reamainder of %v / %v is : %v , %v ", a, b, coef, rem);
}