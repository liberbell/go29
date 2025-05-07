package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"myapp/internal/models"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	version  string
	DB       models.DBModel
}

func (app application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Println(fmt.Sprintf("Starting Back end Server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func OpenDB(dsn string) (*sql.DB, error) {
	// db, err := sql.Open("mysql", dsn)
	db, err := sql.Open("mysql", "james:secret@/widgets?parseTime=true&tls=false")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return db, nil
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.db.dsn, "dsn", "", "Database DSN connection string")
	flag.StringVar(&cfg.env, "env", "development", "Application environment{development|production|maintenance}")

	flag.Parse()
	fmt.Println("hello world")

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := OpenDB(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
		version:  version,
		DB:       models.DBModel{DB: conn},
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}
