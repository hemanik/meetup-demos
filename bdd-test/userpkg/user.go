package userpkg

// User struct
type User struct {
	FirstName string
	LastName  string
}

// New constructor
func New() *User {
	var u User
	return &u
}

// FullName returns concatenation
func (u User) FullName() string {
	return u.FirstName + u.LastName
}
