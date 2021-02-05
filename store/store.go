package store

import (
	"context"
	"log"

	"github.com/AJGherardi/ManageBot/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// TODO: Implment locks

// DB defines a interface for manageing application data
type DB interface {
	GetAllServers() []types.ServerData
	OpenServer(guildID string) types.ServerData
	CreateServer(types.ServerData)
	DeleteServer(guildID string)
	CloseServerWithReplacment(guildID string, replacement types.ServerData)
}

// MongoDB holds and abstracts access to the mongo database and its collections
type MongoDB struct {
	servers *mongo.Collection
}

// GetAllServers returns a slice of all servers
func (d *MongoDB) GetAllServers() []types.ServerData {
	// Get all documents
	cursor, err := d.servers.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	// Make servers slice
	servers := []types.ServerData{}
	for cursor.Next(context.Background()) {
		// Decode into server struct
		var server types.ServerData
		if err = cursor.Decode(&server); err != nil {
			log.Fatal(err)
		}
		// Add to server slice
		servers = append(servers, server)
	}
	return servers
}

// OpenServer returns the server with the given guild id
func (d *MongoDB) OpenServer(guildID string) types.ServerData {
	// Get document and decode
	var server types.ServerData
	d.servers.FindOne(context.Background(), types.ServerData{GuildID: guildID}).Decode(&server)
	return server
}

// CreateServer inserts a server
func (d *MongoDB) CreateServer(server types.ServerData) {
	d.servers.InsertOne(context.Background(), server)
}

// DeleteServer removes the server at the given guild id
func (d *MongoDB) DeleteServer(guildID string) {
	d.servers.DeleteOne(context.Background(), types.ServerData{GuildID: guildID})
}

// CloseServerWithReplacment replaces the server at the given guild id
func (d *MongoDB) CloseServerWithReplacment(guildID string, replacement types.ServerData) {
	d.servers.ReplaceOne(context.Background(), types.ServerData{GuildID: guildID}, replacement)
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
	mongoDb := MongoDB{servers: servers}
	return &mongoDb
}
