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

	//model.Create(db, "Teste")

	// Pegue um ObjectID do banco
	model.Remove(db, "5c0f01a3d7e18b364b64336f")
}
