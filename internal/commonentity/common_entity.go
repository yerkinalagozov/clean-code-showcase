package commonentity

import "strings"

type FullNameItems struct {
	firstName CustomString
	lastName  CustomString
}

func (n *FullNameItems) FirstName() *CustomString {
	return &n.firstName
}

func (n *FullNameItems) LastName() *CustomString {
	return &n.lastName
}

func (n *FullNameItems) Set(firstName, lastName string) *FullNameItems {
	n.firstName = CustomString{
		Val:   firstName,
		Valid: true,
	}
	n.lastName = CustomString{
		Val:   lastName,
		Valid: true,
	}
	return n
}

func (n *FullNameItems) FullName() CustomString {
	return CustomString{
		Val:   n.firstName.Val + " " + n.lastName.Val,
		Valid: true,
	}
}

type Password struct {
	password string
	valid    bool
}

func (p *Password) Valid() bool {
	return p.valid
}

func (p *Password) SetPassword(password *string) {
	// TODO: add hashing logic
	if password != nil {
		p.password = *password
		p.valid = true
	}
	if password == nil {
		p.valid = false
	}
}

func (p *Password) String() string {
	return p.password
}

type CustomInt struct {
	Val   int
	Valid bool
}

func (c *CustomInt) Set(val *int) {
	if val != nil {
		c.Val = *val
		c.Valid = true
	}
	if val == nil {
		c.Valid = false
	}
}

type CustomBool struct {
	Val   bool
	Valid bool
}

func (c *CustomBool) Set(val *bool) {
	if val != nil {
		c.Val = *val
		c.Valid = true
	}
	if val == nil {
		c.Valid = false
	}
}

type CustomString struct {
	Val   string
	Valid bool
}

func (c *CustomString) Set(val *string) {
	if val != nil {
		c.Val = *val
		c.Valid = true
	}
	if val == nil {
		c.Valid = false
	}
}

func (c *CustomString) IsNil() bool {
	return c.Valid
}

type CustomFloat struct {
	Val   float64
	Valid bool
}

func (c *CustomFloat) Set(val *float64) {
	if val != nil {
		c.Val = *val
		c.Valid = true
	}
	if val == nil {
		c.Valid = false
	}
}

type Email struct {
	address string
	domain  string
	valid   bool
}

func (e *Email) Valid() bool {
	return e.valid
}

//type Email email

func (e *Email) String() string {
	return e.address + "@" + e.domain
}

func (e *Email) Domain() string {
	return e.domain
}

func (e *Email) Address() string {
	return e.address
}

func (e *Email) AddEmail(email string) {
	fd := strings.SplitAfter(email, "@")
	fh := strings.ReplaceAll(fd[0], "@", "")
	e.address = fh
	e.domain = fd[1]
	e.valid = true
}
