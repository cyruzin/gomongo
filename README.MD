# gomongo

Repositório de estudo para testar o novo driver do MongoDB para Golang.

Como testar:
  - Na pasta raíz, use o comando docker-compose up -d para subir o container do MongoDB
  - Coloque o IP do container na linha 37 do arquivo "model.go"
  - No terminal digite "go run *.go" para testar
  - No "main.go" tem uma linha comentada para testar a função de remover