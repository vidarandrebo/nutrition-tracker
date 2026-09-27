package user

type User struct {
	ID    int64
	Name  string
	Email string
}

func NewUser() *User {
	return &User{}
}

func (u *User) FromTable(user TableUser) {
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
}
