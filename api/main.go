package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Greeting struct {
	Name     string
	Template string
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

func main() {
	t := Greeting{
		Name:     "Slava",
		Template: "Hello, ",
	}

	collection := client.Database("testing").Collection("greetings")
	_, err := collection.InsertOne(context.TODO(), t)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var g Greeting
		filter := bson.D{{Key: "name", Value: "Slava"}}
		err = collection.FindOne(context.TODO(), filter).Decode(&g)
		if err == mongo.ErrNoDocuments {
			fmt.Fprintf(w, "Record does not exist")
		} else if err != nil {
			fmt.Fprint(w, err.Error())
		}

		fmt.Fprint(w, g.Template, g.Name)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
