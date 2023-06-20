package models

type User struct {
	ID    string   `faker:"uuid_hyphenated" json:"id" bson:"id"`
	Name  string   `faker:"name" json:"name" bson:"name"`
	Email string   `faker:"email" json:"email" bson:"email"`
	Roles []string `json:"roles" bson:"roles"`
}

func NewUser(id, name, email string, roles []string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
		Roles: roles,
	}
}
