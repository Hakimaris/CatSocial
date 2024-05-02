package database

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v4/pgxpool"
    "github.com/joho/godotenv"
)

var dbPool *pgxpool.Pool

func Connect() error {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        return err
    }

    // Get database configuration from environment variables
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbParam := os.Getenv("DB_PARAM")

    // Create database connection string
    connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbParam)

    // Create connection pool
    pool, err := pgxpool.Connect(context.Background(), connString)
    if err != nil {
        return err
    }

    // Store the connection pool
    dbPool = pool

    // Print a message indicating successful connection
    fmt.Println("Successfully connected to the database")

    return nil
}

func GetDBPool() *pgxpool.Pool {
    return dbPool
}

func Close() {
    if dbPool != nil {
        dbPool.Close()
    }
}
