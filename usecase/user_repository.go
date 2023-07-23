package usecase

import "github.com/daisuke0925m/sample-saas/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
