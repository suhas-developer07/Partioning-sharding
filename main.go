package main

import (
	"log"
	"net/http"

	routes "github.com/suhas-developer07/Partioning-sharding/Routes"
	"github.com/suhas-developer07/Partioning-sharding/db"
	"github.com/suhas-developer07/Partioning-sharding/repository"
)

func main(){

	db.Connect()

	shard1Repo := repository.NewPostgresRepositoryShard1(db.Shard1)

	if err:=  shard1Repo.InitShard1(); err != nil {
		log.Fatalln("shard1 init failed:", err)
	}

	shard2Repo := repository.NewPostgresRepositoryShard2(db.Shard2)

	if err := shard2Repo.InitShard2(); err !=nil {
		log.Fatalln("shard2 init failed:", err)
	}

	router := routes.MountRoutes()

	http.ListenAndServe(":8080",router)
}

