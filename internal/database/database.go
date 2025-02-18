package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func Dbconnect() *mongo.Client {
	
	clientOptions := options.Client().ApplyURI("mongodb+srv://jatin_dharaiya:Jatin%4008102000@demo.havv7ft.mongodb.net/golang_api")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Connection Failed to DB")
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pingErr := client.Ping(ctx, nil)

	if pingErr != nil {
		log.Fatal("Failed to connect to MongoDB", err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}