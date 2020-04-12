package model

import "github.com/bbdLe/iGame/comm"

type User struct {
	ServerId string
	Session comm.Session
}

func NewUser(comm comm.Session, ServerId string) *User {
	return &User{
		ServerId: ServerId,
		Session: comm,
	}
}