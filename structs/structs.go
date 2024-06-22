package main
import (
	"fmt"
	"reflect"
)

type Person struct {
	name      string
	age       int
	isMarried bool
	friends   []string
}

// we can also have embedded structs, mimicing the composition pattern
type Job struct {
	Role   string
	Salary float64
}

type Teacher struct {
	Job
	Name string
}

// we can have tags to provide additional information about the attributes of the struct

type Credentials struct {
	Username string `required:"true" max:"20"`
	Password string
}

// the tags does not have significance on their own, but we can parse these to have additional logics in the application
// on their own they are just plain strings.
// backtick is used instead of quotes because they are literal representation of the string and does not consider any escape characters.
// they treat everything as string , special symbols , escape characters etc.

// this is just a syntactic sugar , under the hood it stores a reference variable to the acutal struct.

func changeJob(teacher *Teacher, role string, salary float64) {
	teacher.Role = role
	teacher.Salary = salary
}
func main() {

	dhiraj := Person{
		name:      "Dhiraj Prasad",
		age:       100,
		isMarried: false,
		friends: []string{
			"Praful",
			"Dipesh",
		},
	}
	fmt.Println(dhiraj)

	// using the teacher type
	ram := Teacher{}
	ram.Name = "Ram Sharma"
	// we can directly access the embedded struct vars.
	// but this is not inheritance.
	ram.Role = "Senior Professor"
	ram.Salary = 45000
	fmt.Println(ram)
	// if we want to have a literal declaration, then we cannot directly reference the embedded Structs.

	shyam := Teacher{
		Name: "Shyam Sharma",
		Job:  Job{Role: "Junior Professor", Salary: 35000},
	}

	fmt.Println(shyam)

	// structs are passed by values, i.e a new copy of the struct is passed in the fucntion calls or when we assing it to a new varible,
	// thus if we want to manipulate the actual struct, we have to pass the pointer to the struct and not just the struct variable.

	changeJob(&shyam, "Assistant Professor", 38000)
	fmt.Println(shyam)

	// we can also have anonymous structs.
	suresh := struct{ accessRole string }{"Admin"}
	fmt.Println(suresh)

	// parsing the tags from the struct

	// to do this we used the reflect package.
	t := reflect.TypeOf(Credentials{})
	field, _ := t.FieldByName("Username")
	fmt.Println(field.Tag)
}
