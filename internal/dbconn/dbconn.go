package dbconn

import (
	"Work2Rabbit/database"
	"Work2Rabbit/internal/config"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *database.Queries

func Init(conf *config.Config) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresDatabase,
	)

	psqlDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Couldn't connect to database. %v", err)
		os.Exit(0)
	}

	DB = database.New(psqlDB)
}
