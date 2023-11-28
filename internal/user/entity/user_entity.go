package entity

import (
	"strings"
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
)

type UserStatus struct {
	status string
	valid  bool
}

func (u *UserStatus) String() string {
	return u.status
}

func (u *UserStatus) Set(status *string) {
	if status == nil {
		u.valid = false
		return
	}
	u.status = *status
	u.valid = true
}

func (u *UserStatus) Valid() bool {
	return u.valid
}

const name = "active"

const (
	ActiveUserStatus   = "active"
	InactiveUserStatus = "inactive"
)

type User struct {
	id        int
	userName  commonentity.CustomString
	fullName  commonentity.FullNameItems
	email     commonentity.Email
	age       commonentity.CustomInt
	isMarried commonentity.CustomBool
	password  commonentity.Password
	status    UserStatus
	createdAt time.Time
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *User) EmailDomain() string {
	return u.email.Domain()
}

func (u *User) EmailAddress() string {
	return u.email.Address()
}

func (u *User) Email() *commonentity.Email {
	return &u.email
}

func (u *User) SetEmail(email commonentity.Email) {
	u.email = email
}

func (u *User) UserName() *commonentity.CustomString {
	return &u.userName
}

func (u *User) SetUserName(userName commonentity.CustomString) {
	u.userName = userName
}

func (u *User) Status() *UserStatus {
	return &u.status
}

func (u *User) SetStatus(status *string) {
	if status == nil {
		u.status.valid = false
		return
	}
	u.status.status = *status
	u.status.valid = true
}

func (u *User) IsEmpty() bool {
	return u.id == 0
}

func (u *User) ID() int {
	return u.id
}

func (u *User) SetID(id int) {
	u.id = id
}

func (u *User) FullName() *commonentity.FullNameItems {
	return &u.fullName
}

func (u *User) SetFullName(fullName commonentity.FullNameItems) {
	u.fullName = fullName
}

func (u *User) Age() *commonentity.CustomInt {
	return &u.age
}

func (u *User) SetAge(age *int) error {
	if age == nil {
		return commonentity.ErrIsNotValid
	}
	if age != nil {
		if *age < 0 {
			return commonentity.ErrAgeCannotBeNegative
		}
		if *age < 18 {
			return commonentity.ErrAgeCannotBeLessThanEighteen
		}
		u.age = commonentity.CustomInt{
			Val:   *age,
			Valid: true,
		}
	}
	return nil
}

func (u *User) IsMarried() *commonentity.CustomBool {
	return &u.isMarried
}

func (u *User) SetIsMarried(isMarried *bool) error {
	if isMarried == nil {
		return commonentity.ErrIsNotValid
	}
	u.isMarried.Set(isMarried)
	return nil
}

func (u *User) Password() *commonentity.Password {
	return &u.password
}

func (u *User) SetPassword(password *string) error {
	if password == nil {
		return commonentity.ErrIsNotValid
	}
	if password != nil {
		if strings.Count(*password, "")-1 < 8 {
			return commonentity.ErrPasswordCannotBeLessThanEight
		}
		u.password.SetPassword(password)
	}
	return nil
}
