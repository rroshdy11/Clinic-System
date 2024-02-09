package service

import (
	entities "Clinic_System/entity/User"
	repository "Clinic_System/repository"
	"context"
)

// UserService implements the user.Service interface
type UserService interface {
	GetAll(ctx context.Context) ([]entities.User, error)
	GetAllDoctors(ctx context.Context) ([]entities.User, error)
	SignUp(ctx context.Context, user entities.User) (entities.User, error)
	LogIn(ctx context.Context, email string, password string) (entities.User, error)
}

// make the struct that implements the UserService interface
type userService struct {
	repo repository.UserRepository
}

// New creates a new user service object
func New(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// implement the methods of the UserService interface
func (s userService) GetAll(ctx context.Context) ([]entities.User, error) {
	return s.repo.GetAll()
}

func (s userService) GetAllDoctors(ctx context.Context) ([]entities.User, error) {
	return s.repo.GetAllDoctors()
}

func (s userService) SignUp(ctx context.Context, user entities.User) (entities.User, error) {

	// call the SignUp function to write the new user to the database
	return s.repo.SignUp(user)

}

func (s userService) LogIn(ctx context.Context, email string, password string) (entities.User, error) {
	return s.repo.LogIn(email, password)
}
