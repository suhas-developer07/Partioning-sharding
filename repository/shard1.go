package repository

import (
	"database/sql"
	"log"
)

type PostgresRepositoryShard1 struct{
	db *sql.DB
}

func NewPostgresRepositoryShard1(db *sql.DB) *PostgresRepositoryShard1{
	return &PostgresRepositoryShard1{db:db}
}

func (shard1 *PostgresRepositoryShard1) InitShard1() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS user_analytics_p1 (
			id BIGINT,
			session_time INT,
			country TEXT,
			created_at TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS user_analytics_p2 (
			id BIGINT,
			session_time INT,
			country TEXT,
			created_at TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS user_analytics_p3 (
			id BIGINT,
			session_time INT,
			country TEXT,
			created_at TIMESTAMP
		)`,
	}

	for i,query := range queries{
		if _,err := shard1.db.Exec(query);err!=nil{
			log.Fatalf("unable to create table user_analytics_p%d: %v\n",i+1,err)
			return err
		}
	}

	return nil
}