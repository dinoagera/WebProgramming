package service

type Login interface {
	Login(email, password string) (string, error)
}
type Register interface {
	Register(email, password string) error
}
