package db

import (
	"appointments-system/ports"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresRepo struct {
	DB *sqlx.DB
}

var repo *postgresRepo
var once sync.Once

func Get() ports.Repository {
	once.Do(func() {
		source := fmt.Sprintf(
			`host=%s 
			port=%s 
			user=%s 
			password=%s 
			dbname=%s 
			sslmode=%s`,
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL_MODE"),
		)

		db, err := sqlx.Connect("postgres", source)
		if err != nil {
			log.Panicf("unable to connect to DB: %v", err.Error())
		}

		repo = &postgresRepo{
			DB: db,
		}
	})

	return repo
}
