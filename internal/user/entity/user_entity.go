package entity

import (
	"strings"
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/internal/commonentity"
)

const (
	Married            MarriageStatus = "married"
	NotMarried         MarriageStatus = "not_married"
	ActiveUserStatus   UserStatus     = "active"
	InactiveUserStatus UserStatus     = "inactive"
)

type User struct {
	id        int
	userName  string
	fullName  FullNameItems
	email     Email
	age       int
	isMarried MarriageStatus
	password  Password
	status    UserStatus
	createdAt time.Time
}

func (u *User) Status() string {
	return u.status.String()
}

func (u *User) SetStatus(status string) error {
	err := u.status.Set(status)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) IsMarried() bool {
	return u.isMarried.IsMarried()
}

func (u *User) SetMarried(married string) error {
	err := u.isMarried.Set(married)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FullName() string {
	return u.fullName.FullName()
}

func (u *User) SetFirstLastName(firstName, lastName string) {
	u.fullName.Set(firstName, lastName)
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) SetEmail(email string) error {
	err := u.email.AddEmail(email)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Email() string {
	return u.email.String()
}

func (u *User) UserName() string {
	return u.userName
}

func (u *User) SetUserName(userName string) {
	u.userName = userName
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

func (u *User) Age() int {
	return u.age
}

func (u *User) SetAge(age int) error {
	if age < 18 {
		return commonentity.ErrAgeCannotBeLessThanEighteen
	}
	u.age = age
	return nil
}

func (u *User) Password() string {
	return u.password.Password()
}

func (u *User) SetNewPassword(password string) error {
	err := u.password.SetPassword(password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) SetPassword(password string) {
	u.password.password = password
}

///

type Password struct {
	password string
}

func (p *Password) Password() string {
	return p.password
}

func (p *Password) SetPassword(password string) error {
	if password == "" {
		return commonentity.ErrIsNotValid
	}
	if strings.Count(password, "")-1 < 8 {
		return commonentity.ErrPasswordCannotBeLessThanEight
	}
	p.password = password
	return nil
}

////////////////////

type MarriageStatus string

func (m *MarriageStatus) String() string {
	return string(*m)
}

func (m *MarriageStatus) IsMarried() bool {
	var status bool
	switch *m {
	case Married:
		status = true
	case NotMarried:
		status = false
	}
	return status
}

func (m *MarriageStatus) Set(status string) error {
	switch status {
	case string(Married):
		*m = Married
	case string(NotMarried):
		*m = NotMarried
	default:
		return commonentity.ErrIsNotValid
	}
	return nil
}

//////////////

type Email struct {
	address string
	domain  string
}

func (e *Email) String() string {
	return e.address + "@" + e.domain
}

func (e *Email) Domain() string {
	return e.domain
}

func (e *Email) Address() string {
	return e.address
}

func (e *Email) AddEmail(email string) error {
	if email == "" {
		return commonentity.ErrIsNotValid
	}
	if !strings.Contains(email, "@") {
		return commonentity.ErrEmailIsNotValid
	}
	fd := strings.SplitAfter(email, "@")
	fh := strings.ReplaceAll(fd[0], "@", "")
	e.address = fh
	e.domain = fd[1]

	return nil
}

////////

type FullNameItems struct {
	firstName string
	lastName  string
}

func (n *FullNameItems) FirstName() string {
	return n.firstName
}

func (n *FullNameItems) LastName() string {
	return n.lastName
}

func (n *FullNameItems) Set(firstName, lastName string) {
	n.firstName = firstName
	n.lastName = lastName
}

func (n *FullNameItems) FullName() string {
	return n.firstName + " " + n.lastName
}

//////////////

type UserStatus string

func (u *UserStatus) String() string {
	return string(*u)
}

func (u *UserStatus) Set(status string) error {
	switch status {
	case string(ActiveUserStatus):
		*u = ActiveUserStatus
	case string(InactiveUserStatus):
		*u = InactiveUserStatus
	default:
		return commonentity.ErrUserStatusIsNotValid
	}
	return nil
}

/////////////////////
