package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Contact struct {
	Userone  string `json:"userone" bson:"userone"`
	Usertwo  string `json:"usertwo" bson:"usertwo"`
	Timestamp    time.Time     `json:"timestamp" bson:"timestamp"`
}

type User struct {
	Id           string        `json:"id" bson:"id"`
	Name        string        `json:"name" bson:"title"`
	DOB			 string       `json:"dob" bson:"dob"`
	Phone        string        `json:"phone" bson:"phone"`
	Email        string        `json:"email" bson:"email"`
	Timestamp    time.Time     `json:"timestamp" bson:"timestamp"`
}



func CreateUser(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type","application/json")
	var user User
	_= json.NewDecoder(request.body).Decode(&user)
	collection := client.Database("dheeraj").collection("user")
	ctx,_ := context.WithTimeout(content.Background(), 10*time.Second)
	result,_ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)

}

func getUser(response http.ResponseWriter, request *http.Request){
	var id;
	id = User.ID;
	if(request.ID == id)
	return request;
}

func findUser(w http.ResponseWriter, r *http.Request) {
	
}


func handleRequests() {
	http.HandleFunc("/users", createUser)
	http.HandleFunc("/users/:Id", findUser)
	http.HandleFunc("/users/:Id", getUser)
	log.Fatal(http.ListenAndServe(":5000",nil))	
}


func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	handleRequests()	
}
