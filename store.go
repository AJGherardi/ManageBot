package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DB holds and abstracts access to the database and its collections
type DB struct {
	servers *mongo.Collection
}

// Server holds all information that is needed to manage a discord guild
type Server struct {
	GuildID string `bson:"guildID,omitempty"`
	Name    string `bson:"name,omitempty"`
	// CommandHistory
	// Tickets
	// Warnings
}

// GetAllServers returns a slice of all servers
func (d *DB) GetAllServers() []Server {
	// Get all documents
	cursor, err := d.servers.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	// Make servers slice
	servers := []Server{}
	for cursor.Next(context.Background()) {
		// Decode into server struct
		var server Server
		if err = cursor.Decode(&server); err != nil {
			log.Fatal(err)
		}
		// Add to server slice
		servers = append(servers, server)
	}
	return servers
}

// AddServer inserts a server
func (d *DB) AddServer(guildID, name string) {
	d.servers.InsertOne(context.Background(), Server{
		GuildID: guildID,
		Name:    name,
	})
}

// RemoveServer removes the server at the given guild id
func (d *DB) RemoveServer(guildID string) {
	d.servers.DeleteOne(context.Background(), Server{GuildID: guildID})
}

// OpenDB returns the DB struct which is used to manage application data
func OpenDB() DB {
	// Get db client
	db, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// Connect with timeout
	db.Connect(context.TODO())
	// Test using ping
	db.Ping(context.TODO(), readpref.Primary())
	// Get collection ref
	servers := db.Database("main").Collection("servers")
	return DB{servers: servers}
}
