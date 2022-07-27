package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Album struct {
	ID     int64   `bson:"id"`
	Title  string  `bson:"title"`
	Artist string  `bson:"artist"`
	Price  float32 `bson:"price"`
}

func main() {
	insert()
	update()
	find()
	findall()
	remove()
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb+srv://test:test@cluster0.awkve.mongodb.net/?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database("recordings"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx := context.Background()

	_, err = db.Collection("album").InsertOne(ctx, Album{ID: 5, Title: "Hari Yang Cerah", Artist: "Peterpan", Price: 50000})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("album").InsertOne(ctx, Album{ID: 6, Title: "Sebuah Nama Sebuah Cerita", Artist: "Peterpan", Price: 50000})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert success!")
}

func find() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("album").Find(ctx, bson.M{"id": 1})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Album, 0)
	for csr.Next(ctx) {
		var row Album
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Title  :", result[0].Title)
		fmt.Println("Artist :", result[0].Artist)
		fmt.Println("Price  :", result[0].Price)
	}
}

func findall() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("album").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]Album, 0)
	for csr.Next(ctx) {
		var row Album
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		for _, res := range result {
			fmt.Println("Title  :", res.Title)
			fmt.Println("Artist :", res.Artist)
			fmt.Println("Price  :", res.Price)
		}
	}
}

func update() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"id": 2}
	var changes = Album{ID: 2, Title: "Bintang Di surga", Artist: "Peterpan", Price: 60000}
	_, err = db.Collection("album").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func remove() {
	ctx := context.Background()
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"id": 2}
	_, err = db.Collection("album").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}
