package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"main.go/api/middleware"
	"main.go/internal/database"
	"main.go/internal/database/cache"
	"main.go/internal/model"
)

const (
	RedisAddr     string = "localhost:6379"
	RedisPassword string = ""
	RedisDB       int    = 0
)

var client = database.Dbconnect()
var c = cache.Init(RedisAddr, RedisPassword, RedisDB)

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

	redisKey := fmt.Sprint(user.MobileNumber)

	cachedUser, err := c.Get(context.TODO(), redisKey)
	if err == nil { // Cache hit
		w.WriteHeader(http.StatusOK)
		var userMap map[string]interface{}
		json.Unmarshal([]byte(cachedUser), &userMap)

		middleware.GetRedisDataReponse(userMap, w)
		return
	}

	fmt.Println("monbo db is not calling")

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

	userJSON, _ := json.Marshal(userMap)
	log.Println(fmt.Sprint(redisKey))
	if err := c.Set(context.TODO(), fmt.Sprint(redisKey), userJSON, time.Hour); err != nil {
		log.Println("Error: Value is not stored")
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

	redisKey := fmt.Sprint(user.MobileNumber)
	delCount, err := c.Del(ctx, redisKey).Result()
	if err != nil {
		log.Fatalf("Failed to delete key: %v", err)
	}
	if delCount > 0 {
		fmt.Println("Key deleted successfully!")
	} else {
		fmt.Println("Key not found!")
	}
	fmt.Printf("result matched count %d \n", result.DeletedCount)
	middleware.SuccessResponse("User Sucesully Deleted", w)
})
