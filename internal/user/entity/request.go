package entity

type UserRequest struct {
	id        int
	userName  string
	firstName string
	lastName  string
	isMarried bool
	status    string
}

func (g *UserRequest) Id() int {
	return g.id
}

func (g *UserRequest) SetId(id int) {
	g.id = id
}

func (g *UserRequest) UserName() string {
	return g.userName
}

func (g *UserRequest) SetUserName(userName string) {
	g.userName = userName
}

func (g *UserRequest) FirstName() string {
	return g.firstName
}

func (g *UserRequest) SetFirstName(firstName string) {
	g.firstName = firstName
}

func (g *UserRequest) LastName() string {
	return g.lastName
}

func (g *UserRequest) SetLastName(lastName string) {
	g.lastName = lastName
}

func (g *UserRequest) IsMarried() bool {
	return g.isMarried
}

func (g *UserRequest) SetIsMarried(isMarried bool) {
	g.isMarried = isMarried
}

func (g *UserRequest) Status() string {
	return g.status
}

func (g *UserRequest) SetStatus(status string) {
	g.status = status
}

/////////////////////////////

type GetUsersRequest struct {
	UsersRequest []UserRequest
}

type GetProductRequest struct {
	id          int
	description string
}
