package main

import "fmt"

type SQLite struct{}

func (s SQLite) ClearDatabase() {
	fmt.Println("SQLite: started cleaning")
	fmt.Println("Processing...")
	fmt.Println("SQLite: cleared database")
}

func (s SQLite) NewRow() {
	fmt.Println("SQLite: Added new row")
}

func (s SQLite) MakeMigrations() {
	fmt.Println("SQLite: made migrations")
}

type Redis struct{}

func (r Redis) ClearDatabase() {
	fmt.Println("Redis: started cleaning")
	fmt.Println("Processing...")
	fmt.Println("Redis: cleared database")
}

func (r Redis) MakeMigrations() {
	fmt.Println("No migrations in redis")
}

func (r Redis) Add() {
	fmt.Println("Redis: added new key-pair value")
}

type PostgresSQL struct{}

func (p PostgresSQL) ClearDatabase() {
	fmt.Println("PostgresSQL: started cleaning")
	fmt.Println("Processing...")
	fmt.Println("PostgresSQL: cleared database")
}

func (p PostgresSQL) MakeMigrations() {
	fmt.Println("SQLite: made migrations")
}

func (p PostgresSQL) InsertNewRow() {
	fmt.Println("PostgresSQL: added new row")
}

type DatabaseConnection interface {
	ClearDatabase()
	MakeMigrations()
}

type DatabaseBuilderAdapter struct {
	DatabaseType string
	DatabaseConn interface{}
}

func BuildDatabase(dbType string) *DatabaseBuilderAdapter {
	dbAdapter := &DatabaseBuilderAdapter{DatabaseType: dbType}

	switch dbType {
	case "SQLite":
		dbAdapter.DatabaseConn = SQLite{}
	case "PostgresSQL":
		dbAdapter.DatabaseConn = PostgresSQL{}
	case "Redis":
		dbAdapter.DatabaseConn = Redis{}
	}

	return dbAdapter
}

func (a DatabaseBuilderAdapter) CreateNewRow() {
	switch a.DatabaseType {
	case "SQLite":
		dbConn, _ := (a.DatabaseConn).(SQLite)
		dbConn.NewRow()
	case "PostgresSQL":
		dbConn, _ := (a.DatabaseConn).(PostgresSQL)
		dbConn.InsertNewRow()
	case "Redis":
		dbConn, _ := (a.DatabaseConn).(Redis)
		dbConn.Add()
	}
}

func main() {
	dbBuilder := BuildDatabase("Redis")

	dbBuilder.CreateNewRow()
}
