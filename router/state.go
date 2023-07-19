package router

import (
	"sync"
	"udyweber/go_rest_api/models"
)

type State struct {
	data  map[string]models.User
	mutex sync.Mutex
}

func NewGlobalState() *State {
	return &State{
		data: make(map[string]models.User),
	}
}

func (s *State) Get(key string) (models.User, bool) {
	defer s.mutex.Unlock()

	s.mutex.Lock()

	value, ok := s.data[key]
	return value, ok
}

func (s *State) Set(key string, value models.User) {
	defer s.mutex.Unlock() // Release the mutex to another thread use it safely.

	s.mutex.Lock()
	s.data[key] = value
}

func (s *State) GetAll() []models.User {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var users []models.User

	for _, user := range s.data {
		users = append(users, user)
	}

	return users
}

var GlobalState = NewGlobalState()
