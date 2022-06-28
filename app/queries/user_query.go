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

func (q *UserQueries) GetUser(id string) (models.UserResponse, error) {
	user := models.UserResponse{}
	userFromDB := models.User{}
	filter, err := utils.FilterId(id)
	if err != nil {
		return user, err
	}

	collection := q.DB.Database("library").Collection("users")
	res := collection.FindOne(context.Background(), filter)
	if res.Err() != nil {
		return user, res.Err()
	}

	res.Decode(&userFromDB)
	user.ID = userFromDB.ID
	user.Username = userFromDB.Username
	user.Email = userFromDB.Email
	user.Role = userFromDB.Role

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

func (q *BookQueries) UpdateUser(id primitive.ObjectID, u *models.UserResponse) error {
	collection := q.DB.Database("library").Collection("users")

	user := models.User{}
	user.Email = u.Email
	user.Username = u.Username
	user.Role = u.Role

	_, err := collection.ReplaceOne(context.Background(), primitive.M{"_id": id}, user)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) DeleteUser(id primitive.ObjectID) error {
	collection := q.DB.Database("library").Collection("users")

	_, err := collection.DeleteOne(context.Background(), primitive.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
