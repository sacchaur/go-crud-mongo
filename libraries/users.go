package libraries

import (
	"context"
	"crud_operation/configs"
	"crud_operation/dto"
	"crud_operation/repository"
	"crud_operation/utils"
	"log"
)

type UserService interface {
	Get(ctx context.Context, userId int) (*dto.User, error)
	Add(ctx context.Context, user *dto.User) (*dto.User, error)
	Update(ctx context.Context, userId int, user *dto.User) (*dto.User, error)
	Delete(ctx context.Context, userId int) (bool, error)
	GetAll(ctx context.Context) (*[]dto.User, error)
	AuthenticateUser(username, password string) (bool, error)
}

type userService struct {
	config   configs.ApiConfig
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		config:   configs.GetConfig(),
		userRepo: userRepo,
	}
}

func (us *userService) Get(ctx context.Context, userId int) (*dto.User, error) {
	return us.userRepo.GetUser(ctx, userId)
}

func (us *userService) Add(ctx context.Context, user *dto.User) (*dto.User, error) {
	log.Println("Add user in service")
	salt := utils.GenerateSalt()
	user.Salt = salt
	user.Password = utils.Encrypt(user.Password, salt)
	return us.userRepo.CreateUser(ctx, user)
}

func (us *userService) Update(ctx context.Context, userId int, user *dto.User) (*dto.User, error) {
	return us.userRepo.UpdateUser(ctx, userId, user)
}

func (us *userService) Delete(ctx context.Context, userId int) (bool, error) {
	return us.userRepo.DeleteUser(ctx, userId)
}

func (us *userService) GetAll(ctx context.Context) (*[]dto.User, error) {
	return us.userRepo.GetAllUsers(ctx)
}

func (us *userService) AuthenticateUser(username, password string) (bool, error) {
	return us.userRepo.AuthenticateUser(username, password)
}
