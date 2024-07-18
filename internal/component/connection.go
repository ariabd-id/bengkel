package component

import (
	"bengkel/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GetDatabaseConnection(conf config.Config) *sql.DB {
	dsn := fmt.Sprintf(""+
		"host=%s "+
		"port=%s "+
		"customer=%s "+
		"password=%s "+
		"dbname=%s "+
		"sslmode=disable ",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.User,
		conf.DB.Pass,
		conf.DB.Name,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Error when opening connection to database: %s", err.Error())
	}

	return db
}
