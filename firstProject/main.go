package main

import (
	"fmt"
	"goLang/firstProject/lib"
)

func main(){
	 count := 0;
	lib.Increment(count);
	fmt.Println("The value of count after executing Increment function is :" , count)
}
