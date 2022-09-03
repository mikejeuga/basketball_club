package models

import (
	"github.com/google/uuid"
)

type Birthday struct {
	Day, Month, Year int
}

func NewBirthday(day int, month int, year int) *Birthday {
	return &Birthday{Day: day, Month: month, Year: year}
}

type Player struct {
	ID                  uuid.UUID
	FirstName, LastName string
	Dob                 Birthday
}

func NewPlayer(firstName string, lastName string, dob Birthday) *Player {
	return &Player{FirstName: firstName, LastName: lastName, Dob: dob}
}
