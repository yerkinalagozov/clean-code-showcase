package entity

import "time"

type UserQueryOption func(*UserQuery) error

type UserQuery struct {
	id        []int
	userName  []string
	firstName []string
	lastName  []string
	email     []string
	age       []int
	isMarried []bool
	status    []string
	createdAt []time.Time
}

func (u *UserQuery) CreatedAt() []time.Time {
	return u.createdAt
}

func (u *UserQuery) SetCreateAt(createAt []time.Time) {
	u.createdAt = createAt
}

func (u *UserQuery) Email() []string {
	return u.email
}

func (u *UserQuery) SetEmail(email []string) {
	u.email = email
}

func (u *UserQuery) UserName() []string {
	return u.userName
}

func (u *UserQuery) SetUserName(userName []string) {
	u.userName = userName
}

func (u *UserQuery) Status() []string {
	return u.status
}

func (u *UserQuery) SetStatus(status []string) {
	u.status = status
}

func NewUserQuery(opts ...UserQueryOption) (UserQuery, error) {
	var q UserQuery
	var err error
	for _, opt := range opts {
		err = opt(&q)
		if err != nil {
			return UserQuery{}, err
		}
	}
	return q, nil
}

func UserWithWhereById(id int) UserQueryOption {
	return func(q *UserQuery) error {
		q.id = append(q.id, id)
		return nil
	}
}

func UserWithWhereByFirstName(firstName string) UserQueryOption {
	return func(q *UserQuery) error {
		q.firstName = append(q.firstName, firstName)
		return nil
	}
}

func UserWithWhereByLastName(lastName string) UserQueryOption {
	return func(q *UserQuery) error {
		q.lastName = append(q.lastName, lastName)
		return nil
	}
}

func UserWithWhereByAge(age int) UserQueryOption {
	return func(q *UserQuery) error {
		q.age = append(q.age, age)
		return nil
	}
}

func UserWithWhereByIsMarried(isMarried bool) UserQueryOption {
	return func(q *UserQuery) error {
		q.isMarried = append(q.isMarried, isMarried)
		return nil
	}
}

func UserWithWhereByUserName(userName string) UserQueryOption {
	return func(q *UserQuery) error {
		q.userName = append(q.userName, userName)
		return nil
	}
}

func UserWithWhereByStatus(status string) UserQueryOption {
	return func(q *UserQuery) error {
		q.status = append(q.status, status)
		return nil
	}
}

func UserWithWhereByEmail(email string) UserQueryOption {
	return func(q *UserQuery) error {
		q.email = append(q.email, email)
		return nil
	}
}

func UserWithWhereByCreatedAt(createdAt time.Time) UserQueryOption {
	return func(q *UserQuery) error {
		q.createdAt = append(q.createdAt, createdAt)
		return nil
	}
}

func (u *UserQuery) IDs() []int {
	return u.id
}

func (u *UserQuery) SetID(id []int) {
	u.id = id
}

func (u *UserQuery) FirstName() []string {
	return u.firstName
}

func (u *UserQuery) SetFirstName(firstName []string) {
	u.firstName = firstName
}

func (u *UserQuery) LastName() []string {
	return u.lastName
}

func (u *UserQuery) SetLastName(lastName []string) {
	u.lastName = lastName
}

func (u *UserQuery) Age() []int {
	return u.age
}

func (u *UserQuery) SetAge(age []int) {
	u.age = age
}

func (u *UserQuery) IsMarried() []bool {
	return u.isMarried
}

func (u *UserQuery) SetIsMarried(isMarried []bool) {
	u.isMarried = isMarried
}
