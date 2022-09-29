package ejercicios_tt

type User struct {
	name     string
	lastName string
	age      int
	email    string
	password string
}

func (u *User) SetFullName(name string, lastName string) {
	u.name = name
	u.lastName = lastName
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}
