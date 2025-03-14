package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func IsEven(x int) bool {
	return x%2 == 0
}

func PrintAtInterval(duration time.Duration) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // Ensure the ticker is stopped when done

	// Set up a timeout for the ticker (e.g., run for 10 seconds)
	timeout := time.After(duration)

	for {
		select {
		case <-ticker.C: // This block runs every time the ticker ticks
			fmt.Println("Ticker ticked:", time.Now())
		case <-timeout: // This block runs after the timeout
			fmt.Println("Timeout reached. Stopping ticker.")
			return
		}
	}
}

func DemoGoRoutines(wg *sync.WaitGroup, howmany int) {
	for i := range howmany {
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Printf("Running go routine %v\n", i)
			wg.Done()
		}()
	}
}

// Global variable for the connection pool
var db *sql.DB

// Initialize the connection pool
func initDB() error {
	// Set up connection string
	connStr := "user=postgres password=root dbname=postgres host=localhost sslmode=disable"
	var err error

	// Open the connection pool (max 10 connections, max idle 5)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("could not open database connection: %v", err)
	}

	// Set connection pool properties (optional)
	db.SetMaxOpenConns(10) // Max open connections to the database
	db.SetMaxIdleConns(5)  // Max idle connections in the pool

	// Verify that the database connection is valid
	if err := db.Ping(); err != nil {
		return fmt.Errorf("could not ping database: %v", err)
	}

	fmt.Println("Database connection pool initialized successfully.")
	return nil
}

// Function to fetch PostgreSQL version
func getDbVersion() (string, error) {
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return "", fmt.Errorf("could not query database version: %v", err)
	}
	return version, nil
}

func main() {

	// start := time.Now()
	// fmt.Printf("Starting my program at %v", start)

	// var wg sync.WaitGroup
	// wg.Add(100000)
	// DemoGoRoutines(&wg, 100000)
	// wg.Wait()

	// fmt.Printf("Ended my program after %v", time.Since(start))

	// go PrintAtInterval(10 * time.Second)

	err := initDB()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /go-health", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Running"))
	})

	mux.HandleFunc("GET /go-rand", func(w http.ResponseWriter, r *http.Request) {

		response, err := json.Marshal(struct{ Number int }{Number: rand.Intn(200)})
		if err != nil {
			http.Error(w, fmt.Sprintf("Some error occured %v", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	mux.HandleFunc("GET /go-test-even/{num}", func(w http.ResponseWriter, r *http.Request) {

		strNum := r.PathValue("num")
		intNum, err := strconv.Atoi(strNum)
		if err != nil {
			http.Error(w, fmt.Sprintf("Given input %s is not a valid number", strNum), http.StatusBadRequest)
		}

		response, err := json.Marshal(struct{ IsEven bool }{IsEven: IsEven(intNum)})
		if err != nil {
			http.Error(w, fmt.Sprintf("Some error occured %v", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	mux.HandleFunc("/go-db", func(w http.ResponseWriter, r *http.Request) {
		version, err := getDbVersion()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "PostgreSQL version: %s", version)
	})

	// wg.Wait()

	http.ListenAndServe(":8080", mux)

}
