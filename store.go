package sesame

import (
	"errors"
	"sync"
)

type UserStore interface {
	GetByEmail(string) (*User, error)
	Save(*User) error
}

const (
	NoSuchUserError = "No such User"
)

type InMemoryStore struct {
	sync.RWMutex
	Users map[string]*User
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{Users: make(map[string]*User)}
}

func (store *InMemoryStore) GetByEmail(email string) (u *User, err error) {
	store.RLock()
	u, ok := store.Users[email]
	store.RUnlock()

	if !ok {
		return u, errors.New(NoSuchUserError)
	}

	return
}

func (store *InMemoryStore) Save(u *User) (err error) {
	store.Lock()
	store.Users[u.Email] = u
	store.Unlock()

	return
}

func (store *InMemoryStore) Reset() {
	store.Lock()
	store.Users = make(map[string]*User)
	store.Unlock()
}
