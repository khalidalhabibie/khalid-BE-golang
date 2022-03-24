package database

import (
	"fmt"
	"gokes/pkg/utils"

	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

func PostgreSQLConnection() (*gorm.DB, error) {
	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(postgresConnURL), &gorm.Config{

		// Logger: logger.Default.LogMode(logger.Info),
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		// log.Println(err)
		// panic(err)
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error to assign db to sql db ")
	}

	sqlDB.SetMaxOpenConns(maxConn)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := sqlDB.Ping(); err != nil {
		defer sqlDB.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil

}

// PostgreSQLConnection func for connection to PostgreSQL database.
// func PostgreSQLConnection() (*sqlx.DB, error) {
// 	// Define database connection settings.
// 	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
// 	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
// 	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

// 	// Build PostgreSQL connection URL.
// 	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Define database connection for PostgreSQL.
// 	db, err := sqlx.Connect("pgx", postgresConnURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("error, not connected to database, %w", err)
// 	}

// 	// Set database connection settings:
// 	// 	- SetMaxOpenConns: the default is 0 (unlimited)
// 	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
// 	// 	- SetConnMaxLifetime: 0, connections are reused forever
// 	db.SetMaxOpenConns(maxConn)
// 	db.SetMaxIdleConns(maxIdleConn)
// 	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

// 	// Try to ping database.
// 	if err := db.Ping(); err != nil {
// 		defer db.Close() // close database connection
// 		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
// 	}

// 	return db, nil
// }
