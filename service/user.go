package service

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
	"github.com/Smart-Purveyance-Tracker/backend/repository"
)

type User interface {
	Create(user entity.User) (entity.User, error) // todo
	Login(email, password string) (entity.User, error)
}

type UserImpl struct {
	userRepo repository.User
}

func NewUserImpl(userRepo repository.User) *UserImpl {
	return &UserImpl{
		userRepo: userRepo,
	}
}

func (i *UserImpl) Create(user entity.User) (entity.User, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}

	user.Password = string(pwdHash)
	u, err := i.userRepo.Insert(user)
	if err != nil {
		return entity.User{}, err
	}
	return u, nil
}

func (i *UserImpl) Login(email, password string) (entity.User, error) {
	user, err := i.userRepo.FindByEmail(email)
	if err != nil {
		return entity.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
