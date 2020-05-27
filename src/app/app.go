package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

type App struct {
	Cassandra *gocql.Session
	Router    *gin.Engine
}

func (a *App) StartApp() {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "artisan"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	// defer session.Close()
	if err != nil {
		panic(err)
	}
	a.Cassandra = session
	// a.Cassandra.Close()
	a.Router = gin.Default()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	logs := log.New(os.Stdout, "auth-service ", log.LstdFlags)

	server := &http.Server{
		Addr:         addr,
		Handler:      a.Router,
		ErrorLog:     logs,
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logs.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logs.Println("Recieved terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
