package controller

import (
	"context"
	"fmt"
	"log"
)

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

//most imp

var collection *mango.Collection

// connect with mongodb

func init() {
	//client option
	clientOption := options.client().ApplyURI(connectionString)

	//connect to db
	client, err := mongo.Connect(context.TODO(), clientOption)
	checkErrorNil(err)

	fmt.Println("Mangodb connection is success")

	collection = client.Database(dbName).Collection(colName)

	//collecion instance is ready for me

	fmt.Println("Collection is ready for the instance")

}

func checkErrorNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
