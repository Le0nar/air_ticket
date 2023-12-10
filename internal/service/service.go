package service

type Service struct {
}

type Repository interface{}

func NewService(repository Repository) *Service {
	return &Service{
	}
}
