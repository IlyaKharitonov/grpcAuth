package authStorage

type User struct {
	Name     string
	Lastname string
	Pass     string
	Email    string
	Token    string
}

var AuthStorage = make(map[string]*User)
