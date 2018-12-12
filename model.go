package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var ctx = context.Background()

// DBConn struct
type DBConn struct {
	*mongo.Database
}

// ObjectID struct
type ObjectID struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`
}

// ToDo struct
type ToDo struct {
	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	Completed bool      `bson:"completed,omitempty" json:"completed,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// Conn cria uma conexão com o mongo e devolve a referência
// para Database e um erro.
func (conn *DBConn) Conn(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.NewClient("mongodb://172.21.0.3:27017")

	if err != nil {
		log.Fatalf("Não foi possível conectar com o Mongo: %v", err)
	}

	err = client.Connect(ctx)

	if err != nil {
		log.Fatalf("O cliente do Mongo não pôde se conectar ao contexto de segundo plano: %v", err)
	}

	todoDB := client.Database("gomongo")

	return todoDB, nil
}

// Create insere uma nova tarefa na coleção.
// Aceita um nome como parâmetro.
func (conn *DBConn) Create(db *mongo.Database, t string) {
	task := ToDo{
		Name:      t,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res, err := db.Collection("todo").InsertOne(
		ctx,
		task,
	)

	if err != nil {
		log.Fatalf("Não foi possível criar o registro: %s", err)
	}

	fmt.Println(res.InsertedID)
}

// Remove elimina uma tarefa da coleção.
// Aceita um ID como parâmetro.
func (conn *DBConn) Remove(db *mongo.Database, oID string) {

	id, err := primitive.ObjectIDFromHex(oID)

	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Collection("todo").DeleteOne(
		ctx,
		ObjectID{ID: id},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.DeletedCount)
}
