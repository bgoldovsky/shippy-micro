package main

import (
	pb "github.com/bgoldovsky/shippy-micro/userserver/proto/user"
	"github.com/jinzhu/gorm"
)

// Repository ...
type Repository interface {
	Create(user *pb.User) error
	Get(id string) (*pb.User, error)
	GetAll() ([]*pb.User, error)
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

// UserRepository ...
type UserRepository struct {
	db *gorm.DB
}

// GetAll ...
func (repo *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Get ...
func (repo *UserRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetByEmailAndPassword ...
func (repo *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Create ...
func (repo *UserRepository) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
