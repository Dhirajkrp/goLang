package lib

// a struct helps us group some related fields together.
// we can combine multiple fields as one struct eg:



type Job struct {
	designaion string
	slary float32
}	// we can also have nested struct.
type Person struct {
	Name string
	Age  int
	CurrntJob Job
}


// adding setters and getters to the struct.
type User struct{
	username string
	password string
}

//note that go uses, pass by value and not pass by reference, thus we have to pass the pointer to the struct type if we want to make any changes in the actual values. 

func SetUserName(user *User , name string){
	user.username = name
}
func GetUserName(user User) string{
	return  user.username
}