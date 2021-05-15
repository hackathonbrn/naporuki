package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// User struct
type User struct {
	ID           interface{} `bson:"_id,omitempty"`
	Name         string      `bson:"name"`
	Phone        string      `bson:"phone"`
	PasswordHash string      `bson:"password_hash"`
	Subjects     []string    `bson:"subjects,omitempty"`
	Achievements []string    `bson:"achievements,omitempty"`
	Grades       []float32   `bson:"grades,omitempty"`
	Rating       float32     `bson:"rating,omitempty"`
}

var client = setDBConnection()

func setDBConnection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECT_URL")))
	if err != nil {
		log.Fatal("bad db connect url: ", err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getAllUsers() ([]User, error) {
	collection := client.Database("testing").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("cannot execute find: %v", err)
	}
	defer cur.Close(context.TODO())
	users := make([]User, 0)
	for cur.Next(context.TODO()) {
		var u User
		err := cur.Decode(&u)
		if err != nil {
			return nil, fmt.Errorf("cannot decode result: %v", err)
		}
		users = append(users, u)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("something happened: %v", err)
	}

	return users, nil
}

func getUserByPhone(phone string) (*User, error) {
	var user User
	collection := client.Database("testing").Collection("users")
	err := collection.FindOne(context.TODO(), bson.D{{Key: "phone", Value: phone}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("user does not exist")
	} else if err != nil {
		return nil, fmt.Errorf("cannot execute find: %v", err)
	}

	return &user, nil
}

func addUser(u User) (interface{}, error) {
	collection := client.Database("testing").Collection("users")
	res, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, fmt.Errorf("cannot insert data: %v", err)
	}

	return res.InsertedID, nil
}

func addTestTeacher() error {
	t := User{
		Name:         "teacher",
		Phone:        "89029995361",
		Subjects:     []string{"Алгебра", "Информатика"},
		PasswordHash: "some hash",
		Rating:       4.9,
	}

	collection := client.Database("testing").Collection("users")
	_, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		return fmt.Errorf("cannot insert data: %v", err)
	}

	return nil
}

func addTestStudent() error {
	s := User{
		Name:         "student",
		Phone:        "89005001111",
		PasswordHash: "some hash",
		Grades:       []float32{4.0, 3.0},
		Achievements: []string{"the first one", "cool guy"},
	}

	collection := client.Database("testing").Collection("users")
	_, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		return fmt.Errorf("cannot insert data: %v", err)
	}

	return nil
}
