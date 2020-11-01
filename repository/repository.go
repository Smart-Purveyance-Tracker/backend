package repository

import (
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/Smart-Purveyance-Tracker/backend/entity"
)

type User interface {
	Insert(user entity.User) (entity.User, error)
	Find(id uint64) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}

type UserInMem struct {
	mx           sync.RWMutex
	userIDToInfo map[uint64]entity.User
}

func NewUserInMem() *UserInMem {
	return &UserInMem{
		userIDToInfo: make(map[uint64]entity.User),
	}
}

func (m *UserInMem) Insert(user entity.User) (entity.User, error) {
	id := uuid.New().ID()
	user.ID = uint64(id)
	_, err := m.FindByEmail(user.Email)
	if err == nil {
		return entity.User{}, errors.New("already exists")
	}
	m.mx.Lock()
	m.userIDToInfo[uint64(id)] = user
	m.mx.Unlock()
	return user, nil
}

func (m *UserInMem) Find(id uint64) (entity.User, error) {
	m.mx.RLock()
	user := m.userIDToInfo[id]
	m.mx.RUnlock()
	return user, nil
}

func (m *UserInMem) FindByEmail(email string) (entity.User, error) {
	m.mx.RLock()
	for _, user := range m.userIDToInfo {
		if user.Email == email {
			return user, nil
		}
	}
	m.mx.RUnlock()
	return entity.User{}, errors.New("not found")
}
