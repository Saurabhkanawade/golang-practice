package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://mongo:omsairam@cluster0.rk5loza.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "Watchlist"

// most important
var collection *mongo.Collection

//connect with mongodb

func init() {

	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkErrorNil(err)

	fmt.Println(" Mongo DB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance

	fmt.Println("Collection instance is ready ")

}

//mongodb helper file

// insert one record
func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)
	checkErrorNil(err) //handling the error

	fmt.Println("Inserted one movie into the database ", inserted.InsertedID)
}

// update the record
func updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkErrorNil(err)

	fmt.Println("Modified the count : ", result.ModifiedCount)
}

// delete the record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkErrorNil(err)

	fmt.Println("Movie get delet with delete counr: ", deleteCount)
}

// delete all records
func deleteAllMovie() {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	checkErrorNil(err)

	fmt.Println("Number of movie deleted :", deleteResult.DeletedCount)
}

// get all movies from the database
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	checkErrorNil(err)

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		checkErrorNil(err)

		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

// Actual controller -file

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allmovies := getAllMovies()
	json.NewEncoder(w).Encode(allmovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAdWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	deleteAllMovie()

}

func checkErrorNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
