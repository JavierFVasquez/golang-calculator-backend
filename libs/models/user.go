package models

type User struct {
	ID      string  `faker:"uuid_hyphenated" json:"id" bson:"id"`
	Name    string  `faker:"name" json:"name" bson:"name"`
	Email   string  `faker:"email" json:"email" bson:"email"`
	Balance float64 `faker:"balance" json:"balance" bson:"balance"`
}

func NewUser(id, name, email string, balance float64) *User {
	return &User{
		ID:      id,
		Name:    name,
		Email:   email,
		Balance: balance,
	}
}
