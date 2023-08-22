package service

type CiInterface interface {
	Execute() (string, error)
}
