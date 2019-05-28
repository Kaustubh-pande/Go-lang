package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/mongodb/mongo-go-driver/bson"
	//"github.com/mongodb/mongo-go-driver/bson/primitive"
	//"github.com/mongodb/mongo-go-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo"
)

const (
	//hosts      = "localhost:27017"
	database   = "quotesapi"
	username   = ""
	password   = ""
	collection = "quotes"
)

var client *mongo.Client

type Quotes struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Quote  string             `json:"quote,omitempty" bson:"quote,omitempty"`
	Author string             `json:"author,omitempty" bson:"author,omitempty"`
}

func CreateQuote(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var quotes Quotes
	_ = json.NewDecoder(request.Body).Decode(&quotes)
	collection := client.Database("database").Collection("collection")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, quotes)
	json.NewEncoder(response).Encode(result)
}

func GetQuote(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var quotes Quotes
	collection := client.Database("database").Collection("collection")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Quotes{ID: id}).Decode(&quotes)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(quotes)
}

func GetQuotearr(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var Quotesarr []Quotes
	collection := client.Database("database").Collection("collection")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var quotes Quotes
		cursor.Decode(&quotes)
		Quotesarr = append(Quotesarr, quotes)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(Quotesarr)
}

func sayhelloName(response http.ResponseWriter, request *http.Request) {
	request.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println("hello")
	fmt.Println(request.Form) // print information on server side.
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(response, "Hello User!") // write data to response
}
func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("1.....")
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := mux.NewRouter()
	fmt.Println("2.....")
	http.HandleFunc("/", sayhelloName)
	router.HandleFunc("/quotes", CreateQuote).Methods("POST")
	router.HandleFunc("/Quote", GetQuotearr).Methods("GET")
	router.HandleFunc("/Quote/{id}", GetQuote).Methods("GET")
	log.Fatal(http.ListenAndServe(":9090", router))
	fmt.Println("3.....")
}
