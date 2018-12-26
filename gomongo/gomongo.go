package gomongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client

func init() {
	/*file, err := os.Open("sample.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	read, readerr := ioutil.ReadAll(file)
	if readerr != nil {
		log.Fatal(readerr)
	}*/
	var err error
	client, err = mongo.NewClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo initialized")

}

//DetermineDocFlow method lookup for available document types from Mongodb
func DetermineDocFlow(doctype string) bsonx.Doc {
	collection := client.Database("mydb").Collection("DocTypes")
	/*res, err := collection.InsertOne(context.Background(), bson.M{"data": read})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID

	fmt.Println(id)*/
	//oid, _ := primitive.ObjectIDFromHex("5c1fb4b4a64c4f21d6693a67")
	fmt.Printf("I`m in determine doc flow %s", doctype)
	fmt.Println()
	doc := bsonx.Doc{}
	//filter := bson.D{{"DOCTYPE", doctype}} //Find
	filter := bson.M{"DOCTYPE": doctype}
	retrieve := collection.FindOne(context.Background(), filter)
	//defer retrieve.Close(context.Background())
	retrieve.Decode(&doc)

	fmt.Println(doc.Lookup("SRC_SYS").NullOK())
	/*
		for retrieve.Next(context.Background()) {
			raw, err := retrieve.DecodeBytes()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(raw)
		} */

	return doc
}
