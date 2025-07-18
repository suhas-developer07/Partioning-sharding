package db

import (
	"database/sql"
	"log"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Shard1 *sql.DB
var Shard2 *sql.DB

func Connect(){
	var err error
	
	Shard1,err = sql.Open("pgx","postgres://suhas:secrete123@localhost:5432/shard1")

	if err!=nil{
		log.Fatalf("Failed to connect shard1",err)
	}

	Shard2,err = sql.Open("pgx","postgres://suhas:secrete123@localhost:5433/shard2")

	if err!=nil{
		log.Fatalln("Failed to connect shard2")
	}
}