package user

import (
	"gopkg.in/mgo.v2"

	usermodel "github.com/sf9v/go-db-models/model/user"
)

//Store implements usermodel.Storer
type Store struct {
	session *mgo.Session
}

func (s *Store) List(filters ...usermodel.Filter) ([]usermodel.User, error) {
	return nil, nil
}

func (s *Store) Get(filters ...usermodel.Filter) (usermodel.User, error) {
	return usermodel.User{}, nil
}

func (s *Store) Create(user usermodel.User) error {
	return nil
}

func (s *Store) Update(user usermodel.User) error {
	return nil
}

func (s *Store) Delete(filters ...usermodel.Filter) error {
	return nil
}

func New(session *mgo.Session) usermodel.Storer {
	return &Store{session}
}
