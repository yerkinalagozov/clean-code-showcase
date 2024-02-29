package handler

import (
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/service"
)

type NewUser struct {
	ID        int       `json:"id,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Age       int       `json:"age,omitempty"`
	IsMarried string    `json:"is_married,omitempty"`
	Password  string    `json:"password,omitempty"`
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *NewUser) MapToService() service.NewUserReq {
	var user service.NewUserReq
	user.SetId(n.ID)
	user.SetUserName(n.UserName)
	user.SetFirstName(n.FirstName)
	user.SetLastName(n.LastName)
	user.SetEmail(n.Email)
	user.SetAge(n.Age)
	user.SetIsMarried(n.IsMarried)
	user.SetPassword(n.Password)
	user.SetStatus(n.Status)
	return user
}

type User struct {
	ID        int    `json:"id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Age       int    `json:"age,omitempty"`
	IsMarried string `json:"is_married,omitempty"`
	Password  string `json:"password,omitempty"`
	Status    string `json:"status,omitempty"`
}
