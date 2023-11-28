package commonentity

import "fmt"

var (
	ErrAgeCannotBeNegative           = fmt.Errorf("age cannot be negative")
	ErrAgeCannotBeLessThanEighteen   = fmt.Errorf("age cannot be less than 18")
	ErrIsNotValid                    = fmt.Errorf("is not valid")
	ErrPasswordCannotBeLessThanEight = fmt.Errorf("password cannot be less than 8")
	ErrUnknownStatus                 = fmt.Errorf("unknown status")
)

type DatabaseError3 struct {
	Err       error
	CustomErr error
	Message   string
}

func (e *DatabaseError3) Error() string {
	return fmt.Sprintf("database error: %v, custom error: %v, message: %v", e.Err, e.CustomErr, e.Message)
}

func NewDatabaseError(err error, customErr error, message string) *DatabaseError3 {
	return &DatabaseError3{
		Err:       err,
		CustomErr: customErr,
		Message:   message,
	}
}
