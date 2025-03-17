package models

import "time"

type Phone struct {
	TypeID      int
	CountryCode int
	Operator    int
	Number      int
}

type Contact struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	GivenName  string    `json:"given_name"`
	FamilyName string    `json:"family_name"`
	Phones     []Phone   `json:"phones"`
	Email      []string  `json:"email"`
	Birthdate  time.Time `json:"birthdate"`
}

type Group struct {
	ID          int
	Title       string
	Description string
	Contacts    []int
}
