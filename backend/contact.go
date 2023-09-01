package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	Id      *primitive.ObjectID `bson:"_id"`
	Name    string              `json:"name" bson:"name"`
	Email   string              `json:"email" bson:"email"`
	Message string              `json:"message" bson:"message"`
}

func postContact(res http.ResponseWriter, req *http.Request) {
	contactCollection := db.Collection("contact")
	if contactCollection == nil {
		http.Error(res, "Database error", http.StatusInternalServerError)
		return
	}

	var contact Contact
	err := json.NewDecoder(req.Body).Decode(&contact)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = contactCollection.InsertOne(ctx, contact)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
