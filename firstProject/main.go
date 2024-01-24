package main

import (
	"fmt"
	"goLang/firstProject/lib"
)

func main(){
	 count := 0;
	lib.Increment(count);
	fmt.Println("The value of count after executing Increment function is :" , count)
	ram := lib.Person{}
	ram.Name= "Ram Sharma"
	ram.Age = 21;

	ram.CurrntJob = lib.Job{
		Designaion  : "Intern Engineer",
		Salary : 30500.00,
	}

	dhiraj := lib.User{}

	lib.SetUserName(&dhiraj , "Dhiraj Prasad")

	fmt.Printf("The name of the use dhiraj is : %v ", lib.GetUserName(dhiraj))


	fmt.Printf("Name %v , Age %v", ram.Name, ram.Age)
	fmt.Println(lib.GetPersonName(ram))
	lib.PrintPersonDetails(ram)
}
