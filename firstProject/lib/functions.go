package lib

func Concat(s1 string, s2 string) string {
	return s1 + s2
}

// we can use the comma separated params if they are of the same type an just declare the type once.
func add(a, b int) int {
	return a + b
}

// we can also have a naked return statment , i.e specifying the return value beforehand and not at the time of the return statement.

func GetFullName(firstname string, lastnname string) (fullname string) {
	fullname = firstname + " " + lastnname
	return
}

// in the above function we have predifned the retun variable along with the return type.
