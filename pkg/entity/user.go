package entity

type User struct {
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (User, error) {
	return User{
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}