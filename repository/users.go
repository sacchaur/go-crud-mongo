package repository

import (
	"context"
	"crud_operation/configs"
	"crud_operation/dto"
	"crud_operation/utils"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	GetUser(ctx context.Context, userId int) (*dto.User, error)
	CreateUser(ctx context.Context, user *dto.User) (*dto.User, error)
	UpdateUser(ctx context.Context, userId int, user *dto.User) (*dto.User, error)
	DeleteUser(ctx context.Context, userId int) (bool, error)
	GetAllUsers(ctx context.Context) (*[]dto.User, error)
	AuthenticateUser(username, password string) (bool, error)
}

type userRepository struct {
	cfg             configs.ApiConfig
	storageInstance *mongo.Client
}

func NewUserRepository() UserRepository {
	return &userRepository{
		cfg:             configs.GetConfig(),
		storageInstance: StorageInstance,
	}
}

func (r *userRepository) GetUser(ctx context.Context, userId int) (*dto.User, error) {
	log.Println("Get user in repository")
	userCollection := GetCollection(r.cfg, r.storageInstance, r.cfg.MongoUserCollection)
	var user dto.User
	err := userCollection.FindOne(ctx, bson.M{"userId": userId}).Decode(&user)
	if err != nil {
		log.Printf("Error while fetching user: %+v \n", err)
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *dto.User) (*dto.User, error) {
	log.Println("Add user in repository")
	userCollection := GetCollection(r.cfg, r.storageInstance, r.cfg.MongoUserCollection)
	user.Id = primitive.NewObjectID()
	_, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error while inserting user: %+v \n", err)
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userId int, user *dto.User) (*dto.User, error) {
	log.Println("Update user in repository")

	userCollection := GetCollection(r.cfg, r.storageInstance, r.cfg.MongoUserCollection)
	_, err := userCollection.UpdateOne(ctx, bson.M{"userId": userId}, bson.M{"$set": user})

	if err != nil {
		log.Printf("Error while updating user: %+v \n", err)
		return nil, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userId int) (bool, error) {
	log.Println("Delete user in repository")
	userCollection := GetCollection(r.cfg, r.storageInstance, r.cfg.MongoUserCollection)
	result, err := userCollection.DeleteOne(ctx, bson.M{"userId": userId})
	if result.DeletedCount == 0 {
		log.Printf("User with id %d not found \n", userId)
		return false, nil
	}
	if err != nil {
		log.Printf("Error while deleting user: %+v \n", err)
		return false, err
	}
	return true, nil
}

func (r *userRepository) GetAllUsers(ctx context.Context) (*[]dto.User, error) {
	log.Println("Get all users in repository")
	// string to integer
	// r.cfg.MongoTimeout
	timeout, err := strconv.Atoi(r.cfg.MongoTimeout)
	if err != nil {
		log.Printf("Error while converting string to integer: %+v \n", err)
		return nil, err
	}

	context, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	userCollection := GetCollection(r.cfg, r.storageInstance, r.cfg.MongoUserCollection)
	var users []dto.User
	defer cancel()
	cursor, err := userCollection.Find(context, bson.M{})
	if err != nil {
		log.Printf("Error while fetching all users: %+v \n", err)
		return nil, err
	}
	defer cursor.Close(context)
	for cursor.Next(context) {
		var user dto.User
		if err = cursor.Decode(&user); err != nil {
			log.Printf("Error while decoding user: %+v \n", err)
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

func (r *userRepository) AuthenticateUser(username, password string) (bool, error) {
	collection := GetCollection(r.cfg, r.storageInstance, r.cfg.MongoUserCollection)
	// collection := db.GetCollection("users")
	var user dto.User
	err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return false, err
	}

	return utils.Compare(password, user.Salt, user.Password), nil
}
