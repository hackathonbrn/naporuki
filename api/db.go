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

// Teacher struct
type Teacher struct {
	ID           interface{} `bson:"_id,omitempty"`
	Name         string      `bson:"name"`
	Phone        string      `bson:"phone"`
	Subjects     []string    `bson:"subjects"`
	PasswordHash string      `bson:"password_hash"`
	Rating       float32     `bson:"rating,omitempty"`
}

// Student struct
type Student struct {
	ID           interface{} `bson:"_id,omitempty"`
	Name         string      `bson:"name"`
	Phone        string      `bson:"phone"`
	Achievements []string    `bson:"achievements,omitempty"`
	PasswordHash string      `bson:"password_hash"`
	Grades       []float32   `bson:"grades,omitempty"`
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

func getAllTeachers() ([]Teacher, error) {
	collection := client.Database("testing").Collection("teachers")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("cannot execute find: %v", err)
	}
	defer cur.Close(context.TODO())
	teachers := make([]Teacher, 0)
	for cur.Next(context.TODO()) {
		var t Teacher
		err := cur.Decode(&t)
		if err != nil {
			return nil, fmt.Errorf("cannot decode result: %v", err)
		}
		teachers = append(teachers, t)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("something happened: %v", err)
	}

	return teachers, nil
}

func getAllStudents() ([]Student, error) {
	collection := client.Database("testing").Collection("students")
	cur, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("cannot execute find: %v", err)
	}
	defer cur.Close(context.TODO())
	students := make([]Student, 0)
	for cur.Next(context.TODO()) {
		var s Student
		err := cur.Decode(&s)
		if err != nil {
			return nil, fmt.Errorf("cannot decode result: %v", err)
		}
		students = append(students, s)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("something happened: %v", err)
	}

	return students, nil
}

func addTestTeacher() error {
	t := Teacher{
		Name:         "teacher",
		Phone:        "89029995361",
		Subjects:     []string{"Алгебра", "Информатика"},
		PasswordHash: "some hash",
		Rating:       4.9,
	}

	collection := client.Database("testing").Collection("teachers")
	_, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		return fmt.Errorf("cannot insert data: %v", err)
	}

	return nil
}

func addTestStudent() error {
	s := Student{
		Name:         "student",
		Phone:        "89005001111",
		PasswordHash: "some hash",
		Grades:       []float32{4.0, 3.0},
		Achievements: []string{"the first one", "cool guy"},
	}

	collection := client.Database("testing").Collection("students")
	_, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		return fmt.Errorf("cannot insert data: %v", err)
	}

	return nil
}
