package queries

import (
	"context"
	"fmt"

	"github.com/caiostarke/restApi-and-grpc/app/models"
	"github.com/caiostarke/restApi-and-grpc/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserQueries struct {
	DB *mongo.Client
}

func (q *UserQueries) GetUser(id string) (models.User, error) {
	user := models.User{}
	filter, err := utils.FilterId(id)
	if err != nil {
		return user, err
	}

	collection := q.DB.Database("library").Collection("users")
	res := collection.FindOne(context.Background(), filter)
	if res.Err() != nil {
		return user, res.Err()
	}

	res.Decode(&user)

	return user, nil
}

func (q *BookQueries) CreateUser(u *models.SignUpRequest) error {
	fmt.Println("Creating a user")
	user := models.User{}

	user.Email = u.Email
	user.Username = u.Username
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = passwordEncrypted

	collection := q.DB.Database("library").Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) UpdateUser(id primitive.ObjectID, u *models.User) error {
	collection := q.DB.Database("library").Collection("users")

	_, err := collection.ReplaceOne(context.Background(), primitive.M{"_id": id}, u)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) DeleteUser(id string) error {
	collection := q.DB.Database("library").Collection("users")
	oid, err := utils.FilterId(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.Background(), oid)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) Login(user *models.LoginRequest) (models.UserResponse, error) {
	userRes := models.UserResponse{}
	userModel := models.User{}

	collection := q.DB.Database("library").Collection("users")

	fmt.Println(user.ID)

	res := collection.FindOne(context.Background(), primitive.M{"_id": user.ID})
	if res.Err() != nil {
		return userRes, res.Err()
	}

	if err := res.Decode(&userModel); err != nil {
		return userRes, err
	}

	err := bcrypt.CompareHashAndPassword(userModel.Password, []byte(user.Password))
	if err != nil {
		return userRes, err
	}

	userRes.ID = userModel.ID.Hex()
	userRes.Email = userModel.Email
	userRes.Username = userModel.Username
	userRes.Role = userModel.Role

	return userRes, nil
}
