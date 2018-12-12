package main

import (
	"context"
	"log"
)

var model *DBConn

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	db, err := model.Conn(ctx)

	if err != nil {
		log.Fatalf("A configuração do banco de dados falhou: %v", err)
	}

	model.Create(db, "Fazer a feira")

	// Pegue um ObjectID do banco
	//Ex: 5c105670e09ecd4b2281793b
	//model.Remove(db, "object_id")
}
