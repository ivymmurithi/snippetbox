package main

import (
	"database/sql"
	"flag"
	"html/template"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ivymmurithi/snippetbox/pkg/models/mysql"
	"github.com/joho/godotenv"
	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog 		*log.Logger
	infoLog 		*log.Logger
	session 		*sessions.Session
	snippets 		*mysql.SnippetModel
	templateCache 	map[string]*template.Template
}

func main() {
	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
	}
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	secretKey := os.Getenv("SECRET_KEY")
	sslCertPath := os.Getenv("SSL_CERT_PATH")
	sslKeyPath := os.Getenv("SSL_KEY_PATH")

	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := fmt.Sprintf("web:%s@/snippetbox?parseTime=true", databasePassword)
	secret := flag.String("secret", fmt.Sprintf("%s", secretKey), "Secret Key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	app := &application{
		errorLog: 		errorLog,
		infoLog: 		infoLog,
		session:		session,
		snippets: 		&mysql.SnippetModel{DB: db},
		templateCache: 	templateCache,
	}

	srv := &http.Server{
		Addr:	*addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS(sslCertPath, sslKeyPath)
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err =db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}