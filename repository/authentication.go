package repository

import (
	"context"
	"crud_operation/configs"
	"crud_operation/dto"
	"crud_operation/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthenticationRepository interface {
	AuthenticateUser(username, password string) (bool, error)
	Login(username, password string) (bool, error)
	ResetPassword(username, password string) (bool, error)
}

type authenticationRepository struct {
	cfg             configs.ApiConfig
	storageInstance *mongo.Client
}

func NewAuthenticationRepository() AuthenticationRepository {
	return &authenticationRepository{
		cfg:             configs.GetConfig(),
		storageInstance: StorageInstance,
	}
}

func (auth *authenticationRepository) AuthenticateUser(username, password string) (bool, error) {
	collection := GetCollection(auth.cfg, auth.storageInstance, auth.cfg.MongoUserCollection)
	var user dto.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false, err
	}

	return utils.Compare(password, user.Salt, user.Password), nil
}

func (auth *authenticationRepository) Login(username, password string) (bool, error) {
	return auth.AuthenticateUser(username, password)
}

func (auth *authenticationRepository) ResetPassword(username, password string) (bool, error) {
	collection := GetCollection(auth.cfg, auth.storageInstance, auth.cfg.MongoUserCollection)
	var user dto.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false, err
	}

	salt := utils.GenerateSalt()
	password = utils.Encrypt(password, salt)
	_, err = collection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"password": password, "salt": salt}})
	if err != nil {
		return false, err
	}

	return true, nil
}
