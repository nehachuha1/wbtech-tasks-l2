package main

import "fmt"

type connectionUrl string

type Connection interface {
	executeQuery(string)
}

type Redis struct {
	connectionUrl
}

func (r *Redis) executeQuery(query string) {
	fmt.Printf("Executing query in redis: %s\n", query)
}

func NewRedis() *Redis {
	return &Redis{connectionUrl: "connection to Redis DB"}
}

type Postgres struct {
	connectionUrl
}

func (p *Postgres) executeQuery(query string) {
	fmt.Printf("Executing query in postgres: %s\n", query)
}

func NewPostgres() *Postgres {
	return &Postgres{connectionUrl: "connection to Postgres"}
}

func getConnection(typeOfConnection string) (Connection, error) {
	if typeOfConnection == "redis" {
		return NewRedis(), nil
	} else if typeOfConnection == "postgres" {
		return NewPostgres(), nil
	}

	return nil, fmt.Errorf("wrong type of connection")
}

func main() {
	postgres, _ := getConnection("postgres")
	redis, _ := getConnection("redis")

	postgres.executeQuery("some query")
	redis.executeQuery("some query")
}
