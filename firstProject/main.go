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
	ram.Name = "Ram Sharma"
	ram.Age = 21;

	fmt.Printf("Name %v , Age %v", ram.Name, ram.Age)
	fmt.Println(lib.GetPersonName(ram))
}
