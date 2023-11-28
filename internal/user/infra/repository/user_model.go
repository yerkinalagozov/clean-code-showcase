package repository

import (
	"database/sql"
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/entity"
)

type UserData struct {
	ID        int            `json:"id,omitempty"`
	UserName  sql.NullString `json:"user_name,omitempty"`
	FirstName sql.NullString `json:"first_name,omitempty"`
	LastName  sql.NullString `json:"last_name,omitempty"`
	Email     sql.NullString `json:"email,omitempty"`
	Age       sql.NullInt32  `json:"age,omitempty"`
	IsMarried sql.NullBool   `json:"is_married,omitempty"`
	Password  sql.NullString `json:"password,omitempty"`
	Status    sql.NullString `json:"status,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
}

func (u *UserData) MapToRepoNew(user entity.User) {
	if user.UserName().Valid == true {
		u.UserName.Valid = true
		u.UserName.String = user.UserName().Val
	}
	if user.FullName().FirstName().Valid {
		u.FirstName.Valid = true
		u.FirstName.String = user.FullName().FirstName().Val
	}
	if user.FullName().LastName().Valid {
		u.LastName.Valid = true
		u.LastName.String = user.FullName().LastName().Val
	}
	if user.Email().Valid() {
		u.Email.Valid = true
		u.Email.String = user.Email().Address()
	}
	if user.Age().Valid {
		u.Age.Valid = true
		u.Age.Int32 = int32(user.Age().Val)
	}
	if user.IsMarried().Valid {
		u.IsMarried.Valid = true
		u.IsMarried.Bool = user.IsMarried().Val
	}
	if user.Password().Valid() {
		u.Password.Valid = true
		u.Password.String = user.Password().String()
	}
	if user.Status().Valid() {
		u.Status.Valid = true
		u.Status.String = user.Status().String()
	}
	u.CreatedAt = time.Now().UTC()
}

func (u *UserData) MapToRepoNewList(users []entity.User) []UserData {
	var result []UserData
	for _, user := range users {
		userData := UserData{}
		userData.MapToRepoNew(user)
		result = append(result, userData)
	}
	return result
}

func (u *UserData) MapToRepo(user entity.User) {
	u.ID = user.ID()
	if user.UserName().Valid == true {
		u.UserName.Valid = true
		u.UserName.String = user.UserName().Val
	}
	if user.FullName().FirstName().Valid {
		u.FirstName.Valid = true
		u.FirstName.String = user.FullName().FirstName().Val
	}
	if user.FullName().LastName().Valid {
		u.LastName.Valid = true
		u.LastName.String = user.FullName().LastName().Val
	}
	if user.Email().Valid() {
		u.Email.Valid = true
		u.Email.String = user.Email().Address()
	}
	if user.Age().Valid {
		u.Age.Valid = true
		u.Age.Int32 = int32(user.Age().Val)
	}
	if user.IsMarried().Valid {
		u.IsMarried.Valid = true
		u.IsMarried.Bool = user.IsMarried().Val
	}
	if user.Password().Valid() {
		u.Password.Valid = true
		u.Password.String = user.Password().String()
	}
	if user.Status().Valid() {
		u.Status.Valid = true
		u.Status.String = user.Status().String()
	}
	u.CreatedAt = time.Now().UTC()
}

func (u *UserData) MapToRepoList(users []entity.User) []UserData {
	var result []UserData
	for _, user := range users {
		userData := UserData{}
		userData.MapToRepo(user)
		result = append(result, userData)
	}
	return result
}

func (u *UserData) MapToEntity() (entity.User, error) {
	var user entity.User
	user.SetID(u.ID)
	if u.UserName.Valid {
		user.SetUserName(commonentity.CustomString{
			Val:   u.UserName.String,
			Valid: u.UserName.Valid,
		})
	}
	if u.FirstName.Valid && u.LastName.Valid {
		fullNameItem := commonentity.FullNameItems{}
		fullName := fullNameItem.Set(u.FirstName.String, u.LastName.String)
		user.SetFullName(*fullName)
	}
	if u.Email.Valid {
		emailItems := commonentity.Email{}
		emailItems.AddEmail(u.Email.String)
		user.SetEmail(emailItems)
	}
	if u.Age.Valid {
		ageInt := int(u.Age.Int32)
		err := user.SetAge(&ageInt)
		if err != nil {
			return entity.User{}, err
		}
	}
	if u.IsMarried.Valid {
		err := user.SetIsMarried(&u.IsMarried.Bool)
		if err != nil {
			return entity.User{}, err
		}
	}
	if u.Password.Valid {
		err := user.SetPassword(&u.Password.String)
		if err != nil {
			return entity.User{}, err
		}
	}
	if u.Status.Valid {
		user.SetStatus(&u.Status.String)
	}
	user.SetCreatedAt(u.CreatedAt)
	return user, nil
}

func (u *UserData) MapToEntityList(users []UserData) ([]entity.User, error) {
	var result []entity.User
	for _, user := range users {
		userEntity, err := user.MapToEntity()
		if err != nil {
			return nil, err
		}
		result = append(result, userEntity)
	}
	return result, nil
}
