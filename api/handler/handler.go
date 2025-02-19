package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"main.go/api/middleware"
	"main.go/internal/database"
	"main.go/internal/model"
)

var client = database.Dbconnect()

var ServerStart = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server and route is working")

	slog.Info("server and router working")
	middleware.SuccessResponse("server is running", w)
})

var RegisterUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user model.UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), w)
	}

	collection := client.Database("golang_api").Collection("authentication")

	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), w)
		return
	}

	fmt.Println("Inserted document ID:", res.InsertedID)

	middleware.SuccessResponse("Updated", w)

})

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user model.UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), w)
	}

	collection := client.Database("golang_api").Collection("authentication")

	updateerr := collection.FindOne(context.TODO(), bson.M{"mobilenumber": user.MobileNumber}).Decode(&user)

	if updateerr != nil {
		middleware.ServerErrResponse(updateerr.Error(), w)
		return
	}

	userMap := map[string]interface{}{
		"firstname":    user.Firstname,
		"lastname":     user.Lastname,
		"age":          user.Age,
		"mobilenumber": user.MobileNumber,
	}
	middleware.GetDataReponse(userMap, w)

})

var UserUpdate = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user model.UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), w)
	}
	collection := client.Database("golang_api").Collection("authentication")

	filter := bson.M{"mobilenumber": user.MobileNumber}
	updatename := bson.M{"$set": bson.M{"firstname": user.Firstname}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.UpdateOne(ctx, filter, updatename)
	if err != nil {
		middleware.ErrorResponse(err.Error(), w)
	}
	fmt.Printf("result matched count %d \n", result.MatchedCount)
	middleware.SuccessResponse("User Sucesully updated", w)
})


var DeleteUser = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user model.UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		middleware.ServerErrResponse(err.Error(), w)
	}

	collection := client.Database("golang_api").Collection("authentication")

	filter := bson.M{"mobilenumber": user.MobileNumber}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		middleware.ErrorResponse(err.Error(), w)
	}
	fmt.Printf("result matched count %d \n", result.DeletedCount)
	middleware.SuccessResponse("User Sucesully Deleted", w)
})