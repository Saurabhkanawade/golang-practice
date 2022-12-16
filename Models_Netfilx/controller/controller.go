package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
	"github.com/saurabhkanawade/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/contains"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"


const connectionString = "mongodb+srv://mongo:omsairam@cluster0.rk5loza.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const colName = "Watchlist"

// most important
var collection *mongo.Collection

//connect with mongodb

 func init() {

	//client option
	clientOption := options.client().ApplyURL(connectionString)

	//connect to mongodb
  client ,err:=mongo.Connect(context.TODO(),clientOption)
  checkErrorNil(err)
 

   fmt.Println(" Mongo DB connection success")

   collection = client.Database(dbName).Collection(colName)

   //collection instance

   fmt.Println("Collection instance is ready ")


   //mongodb helper file 

   //insert one record
   func insertOneRecord(movie model.Netfilx){
	  
	inserted , err := collection.InsertOne(context.Background(),movie)
	checkErrorNil(err) //handling the error

	fmt.Println("Inserted one movie into the database ",inserted.InsertedID)
   }

   // update the record
   func updateRecord(movieId string){
	id,_ :=  primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}

	result,err := collection.UpdateOne(context.Background(),filter,update)
	checkErrorNil(err)
	
	fmt.Println("Modified the count : ",result.ModifiedCount)
   }

   // delete the record
   func deleteOneRecord(movieId string){
	id,_:=primitive.ObjectIDfromHex(movieId)
	filter:= bson.M{"_id":id}
	deleteCount,err:=collection.DeleteOne(context.Background(),filter)
	checkErrorNil(err)

	fmt.Println("Movie get delet with delete counr: ",deleteCount)
   }

   //delete all records
   func deleteAllRecords(){
   deleteResult,err=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	checkErrorNil(err)
   }

   fmt.Println("Number of movie deleted :",deleteResult.deleteCount)
}

func checkErrorNil(err error){
	if err!=nil{
		log.Fatal(err)
	   }
}

