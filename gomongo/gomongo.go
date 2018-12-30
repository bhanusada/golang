package gomongo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/x/bsonx"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Environment struct {
	Profile string `json:"profile"`
	Dev     Dev    `json:"dev"`
	Local   Local  `json:"local"`
}

type Dev struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Param    string `json:"param"`
}

type Local struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

var client *mongo.Client

func buildConnectionString() string {

	uri := "mongodb://"

	var j map[string]interface{}

	envfile, _ := ioutil.ReadFile("../gomongo/environment.json")

	json.Unmarshal(envfile, &j)

	system := j["profile"].(string)

	env := j[system].(map[string]interface{})

	if env["username"] != "" && env["password"] != "" {
		uri = uri + fmt.Sprintf("%s", env["username"]) + ":" + fmt.Sprintf("%s", env["password"]) + "@"
	}

	uri = uri + fmt.Sprintf("%s", env["host"]) + ":" + fmt.Sprintf("%s", env["port"])

	if env["param"] != "" {
		uri = uri + "/?" + fmt.Sprintf("%s", env["param"])
	}

	fmt.Println(uri)

	return uri
}

func init() {

	uri := buildConnectionString()

	var err error

	client, err = mongo.NewClient(uri)
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
func DetermineDocFlow(doctype string) bool {
	collection := client.Database("mydb").Collection("DocTypes")
	//collection := client.Database("qars1db3").Collection("DocTypes")
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

	return doc.Lookup("SRC_SYS").NullOK()
}
