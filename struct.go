package main
import "fmt"

type User struct {
	Name string
	Age int
}

func namdAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Jane Doe", 23}
	case 1:
		return User{"John Doe", 27}
	default:
		return User{"None", -1}
	}
}

func incrementAge1(user User) {
	user.Age++
	fmt.Println(user.Age)
}

func incrementAge2(user *User) {	// with 'reference'
	user.Age++
	fmt.Println(user.Age)
}

func main() {
	john := User{"John", 19}
	incrementAge1(john)
	fmt.Println(john.Age)	// not changed

	incrementAge2(&john)
	fmt.Println(john.Age)	// changed
}