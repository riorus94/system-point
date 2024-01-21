package domain

type UserCredentials struct {
	Id       int
	UserId   int
	Password string
	Salt     string
}
