package _interface

type CRUD interface {
	Create() (msg string, err error)
	Update() (msg string, err error)
	Retrieve() (msg string, err error)
	Delete() (msg string, err error)
}
