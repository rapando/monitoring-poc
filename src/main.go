package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rapando/monitoring-poc/src/api"
	"github.com/rapando/monitoring-poc/src/pkg/conn"
	"github.com/rapando/monitoring-poc/src/pkg/log"
	"github.com/rs/cors"
)

func main() {
	var err = godotenv.Load()
	if err != nil {
		panic("failed to read .env")
	}

	log.InitLogger()

	time.Sleep(time.Second * 10)

	err = conn.DbConnect()
	if err != nil {
		log.Errorf("init", "failed to connect to db because %v", err)
	}

	log.Infof("init", "starting out")

	var router = chi.NewRouter()
	var _cors = cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           300,
	})
	router.Use(_cors.Handler)
	api.InitRouter(router)

	log.Infof("init", "starting API at %s", time.Now().Format(time.Kitchen))

	// start API
	err = http.ListenAndServe("0.0.0.0:5001", router)
	if err != nil {
		log.Errorf("init", "failed to initiate API because %v", err)
	}

}
