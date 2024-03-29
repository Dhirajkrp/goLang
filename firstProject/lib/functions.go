package lib

import "fmt"

func Concat(s1 string, s2 string) string {
	return s1 + s2
}

// we can use the comma separated params if they are of the same type an just declare the type once.
func Add(a, b int) int {
	return a + b
}

// we can also have a naked return statment , i.e specifying the return value beforehand and not at the time of the return statement.

func GetFullName(firstname string, lastnname string) (fullname string) {
	fullname = firstname + " " + lastnname
	return
}

// in the above function we have predifned the retun variable along with the return type.

/*-----------------funcitons as varibles----------------*/

// functions are treated as variables in go, i.e they are first class citizens, thus we can pass and return a fuction to and from and another fucntion.

//lets take a fucntion example which takes a  function as its argument.

func AddThree(addTwo func(int, int) int, c int) int {
	return addTwo(2, 3) + c
}

// we can also return more than one value from a function , this patter is common where the value and the error is returned from a function.

func Divide(a, b int) (coef, rem int) {
	coef = a / b
	rem = a % b
	return
}
// passing data into fucntions : data is passed by value and not by reference to the fucntions.
// this means that the values are send as a copy to the function and any updates to that copy will only be temporary and will not reflect on the actual variable itself.

//eg : lets take an example of funciton which takes a variable count increments it but 1 and prints the value.

func Increment(count int) {
	count++;
	fmt.Println("The value of the variable count is:  " , count)
}
// the above function will increment the current value and print it, but the acutal value of the variable passed will not change.

// thus in order to increment the actual variable we can return the incremented value and update the variable with it.


func IncrementAndUpdate(count int) int {
	count = count +1
	return count
}

// the above function returns the updated value, so we can do something like,
// count = IncrementAndUpdate(count);
// thus successfully updating the cout.

/* ---------------- Ignoring some return values ------------------- */

// if we want to ignore a return value, we can use the underscore, '_' , this is different from other languages such as javascript where the undescore is just a naming convention, in go the value is acutally ignored and we can no longer use the varialble.


// the above Divide function can be used just to get the coeffecient and not the remainder.

// coef, _ := Divide(10,3);

func GetPersonName(person Person) string {
	return person.Name
}        