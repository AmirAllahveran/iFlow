package store

import (
	"github.com/jinzhu/gorm"
	"github.com/smf8/http-monitor/model"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

// GetUserByUserName retrieves user from database based on it's ID
// this method loads user's URLs and Requests lists
// returns error if user was not found
func (s *Store) GetUserByUserName(username string) (*model.User, error) {
	user := new(model.User)
	// remove preloading in the future if necessary
	if err := s.db.First(user, model.User{Username: username}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user from database with given id
// returns error if user was not found
func (s *Store) GetUserByID(id uint) (*model.User, error) {
	usr := &model.User{}
	usr.ID = id
	if err := s.db.Model(usr).Find(usr).Error; err != nil {
		return nil, err
	}
	return usr, nil
}

// GetAllUsers retrieves all users from database
func (s *Store) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// AddUser adds a user to the database
func (s *Store) AddUser(user *model.User) error {
	return s.db.Create(user).Error
}

func (s *Store) AddExperiment(experiment *model.Experiment) error {
	return s.db.Create(experiment).Error
}

func (s *Store) AddModule(module *model.Module) error {
	return s.db.Create(module).Error
}

func (s *Store) AddPackage(experimentPackage *model.Package) error {
	return s.db.Create(experimentPackage).Error
}
