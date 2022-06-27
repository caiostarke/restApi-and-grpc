package queries

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/caiostarke/restApi-and-grpc/app/models"
	"github.com/caiostarke/restApi-and-grpc/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookQueries struct {
	DB *mongo.Client
}

func (q *BookQueries) GetBooks() ([]models.Book, error) {
	books := []models.Book{}
	collection := q.DB.Database("library").Collection("books")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		book := models.Book{}
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return books, nil
}

func (q *BookQueries) GetBook(id string) (models.Book, error) {
	book := models.Book{}
	oid, err := utils.FilterId(id)
	if err != nil {
		return book, err
	}

	fmt.Println(oid)

	collection := q.DB.Database("library").Collection("books")
	res := collection.FindOne(context.Background(), oid)
	res.Decode(&book)

	return book, nil
}

func (q *BookQueries) CreateBook(b *models.Book) error {
	fmt.Println("Creating a book")
	collection := q.DB.Database("library").Collection("books")

	_, err := collection.InsertOne(context.Background(), b)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) UpdateBook(id primitive.ObjectID, b *models.Book) error {
	collection := q.DB.Database("library").Collection("books")

	_, err := collection.ReplaceOne(context.Background(), primitive.M{"_id": id}, b)
	if err != nil {
		return err
	}

	return nil
}

func (q *BookQueries) DeleteBook(id primitive.ObjectID) error {
	collection := q.DB.Database("library").Collection("books")

	_, err := collection.DeleteOne(context.Background(), primitive.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
