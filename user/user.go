package user

type User struct {
	UserID    UserID
	Username  Username
	UserEmail UserEmail
}

func New(i, u, e string) *User {
	return &User{
		UserID:    UserID(i),
		Username:  Username(u),
		UserEmail: UserEmail(e),
	}
}

type UserID string

func (id UserID) String() string {
	return string(id)
}

type Username string

func (u Username) String() string {
	return string(u)
}

type UserEmail string

func (e UserEmail) String() string {
	return string(e)
}
