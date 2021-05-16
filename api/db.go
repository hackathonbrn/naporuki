package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// User struct
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Phone        string             `bson:"phone" json:"phone"`
	PasswordHash string             `bson:"password_hash" json:"password_hash"`
	Subjects     []string           `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Achievements []string           `bson:"achievements,omitempty" json:"achievements,omitempty"`
	Grades       []float32          `bson:"grades,omitempty" json:"grades,omitempty"`
	Rating       float32            `bson:"rating,omitempty" json:"rating,omitempty"`
}

// Profile struct
type Profile struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	User User               `bson:"user" json:"user"`
	Desc string             `bson:"desc" json:"desc"`
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

func getAllProfiles() ([]Profile, error) {
	collection := client.Database("testing").Collection("profiles")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("cannot execute find: %v", err)
	}
	defer cur.Close(context.TODO())
	profiles := make([]Profile, 0)
	for cur.Next(context.TODO()) {
		var p Profile
		err := cur.Decode(&p)
		if err != nil {
			return nil, fmt.Errorf("cannot decode result: %v", err)
		}
		profiles = append(profiles, p)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("something happened: %v", err)
	}

	return profiles, nil
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

func getUserProfileByPhone(phone string) (*Profile, error) {
	var p Profile
	collection := client.Database("testing").Collection("profiles")
	err := collection.FindOne(context.TODO(), bson.D{{Key: "user.phone", Value: phone}}).Decode(&p)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("profile does not exist")
	} else if err != nil {
		return nil, fmt.Errorf("cannot execute find: %v", err)
	}

	return &p, nil
}

func addUser(u User) (primitive.ObjectID, error) {
	collection := client.Database("testing").Collection("users")
	res, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("cannot insert data: %v", err)
	}

	p := Profile{
		User: u,
	}

	if err := addProfile(p); err != nil {
		return primitive.ObjectID{}, fmt.Errorf("cannot insert data: %v", err)
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func addProfile(p Profile) error {
	collection := client.Database("testing").Collection("profiles")
	_, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		return fmt.Errorf("cannot insert data: %v", err)
	}

	return nil
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
