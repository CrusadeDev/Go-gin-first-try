package Model

type User struct {
	tableName struct{} `pg:"user"`
	ID        int
	FirstName string
	LastName  string
}

func (u *User) GetId() int {
	return u.ID
}

func (u *User) GetFullName() string {
	return u.FirstName + "" + u.LastName
}
