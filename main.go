package main

import (
	"context"
	"log"
)

var model *DBConn

func main() {

	ctx := context.Background()
	c, cancel := context.WithCancel(ctx)

	defer cancel()

	db, err := model.Conn(c)

	if err != nil {
		log.Fatalf("A configuração do banco de dados falhou: %v", err)
	}

	model.Create(db, "Fazer a feira")

	// Pegue um ObjectID do banco
	// Ex: 5c105670e09ecd4b2281793b
	// model.Remove(db, "object_id")
}
