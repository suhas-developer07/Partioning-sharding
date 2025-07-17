package repository

import (
	"database/sql"
	"log"
)

type PostgresRepositoryShard2 struct{
	db *sql.DB
}

func NewPostgresRepositoryShard2(db *sql.DB) *PostgresRepositoryShard2{
	return &PostgresRepositoryShard2{db:db}
}

func (shard2 *PostgresRepositoryShard2) InitShard2() error{
	queries := []string{
		`CREATE TABLE IF NOT EXISTS user_analytics_p4 (
			id BIGINT,
			session_time INT,
			country TEXT,
			created_at TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS user_analytics_p5 (
			id BIGINT,
			session_time INT,
			country TEXT,
			created_at TIMESTAMP
		)`,
		
	}

	for i,query := range queries {
		if _,err := shard2.db.Exec(query);err!=nil{
			log.Fatalf("unable to create a table user_analytics_p%d : %v\n", i+4,err)
			return err
		}
	}

	return nil
}