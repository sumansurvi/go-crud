package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// Generate a connection string without a password
func connectionString() string {
	host := viper.GetString("db_host")  // or "127.0.0.1"
	port := viper.GetInt("db_port")     // default PostgreSQL port
	user := viper.GetString("db_user")  // your PostgreSQL username
	dbname := viper.GetString("dbname") // your database name

	// Return the formatted connection string
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
}

// Initialize PostgreSQL connection
func initDB() (*sql.DB, error) {
	var err error
	connStr := connectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to the database: " + err.Error())
		return nil, err
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		logger.Error("Could not ping the database: " + err.Error())
		return nil, err
	}
	logger.Info("Successfully connected to the database")
	return db, nil
}
