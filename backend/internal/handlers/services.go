package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"nextjs-golang/internal/db"
)

type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
}

func GetServices(w http.ResponseWriter, r *http.Request) {
	collection := db.GetCollection("modernApp", "services")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching services", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var services []Service
	if err = cursor.All(context.TODO(), &services); err != nil {
		http.Error(w, "Error decoding services", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(services)
}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var service Service
	json.NewDecoder(r.Body).Decode(&service)

	collection := db.GetCollection("modernApp", "services")
	_, err := collection.InsertOne(context.TODO(), service)
	if err != nil {
		http.Error(w, "Error creating service", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Service created successfully"})
}
