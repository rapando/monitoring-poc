package api

import (
	"context"
	"net/http"

	"github.com/rapando/monitoring-poc/src/models"
	"github.com/rapando/monitoring-poc/src/pkg/conn"
	"github.com/rapando/monitoring-poc/src/pkg/helpers"
	"github.com/rapando/monitoring-poc/src/pkg/log"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var db = conn.GetDB()
	var q = models.New(db)
	var ctx = context.Background()
	var requestID = r.Context().Value(RequestIDKey).(string)
	log.Infof(requestID, "request to add data")

	var err = q.CreateRandomData(ctx, models.CreateRandomDataParams{
		X: helpers.GenerateRandomStr(),
		Y: helpers.GenerateRandomStr(),
	})
	if err != nil {
		log.Warnf(requestID, "failed to add data because %v", err)
		Response(w, http.StatusBadRequest, map[string]any{
			"ok": false,
			"id": requestID,
		})
		return
	}

	log.Infof(requestID, "data added successfully")
	Response(w, http.StatusCreated, map[string]any{
		"ok": true,
		"id": requestID,
	})

}

func CountHandler(w http.ResponseWriter, r *http.Request) {
	var db = conn.GetDB()
	var q = models.New(db)
	var ctx = context.Background()
	var requestID = r.Context().Value(RequestIDKey).(string)
	log.Infof(requestID, "request to count data")

	count, err := q.CountData(ctx)
	if err != nil {
		log.Warnf(requestID, "failed to count data because %v", err)
		Response(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"count": 0,
			"id":    requestID,
		})
		return
	}

	log.Infof(requestID, "request to count data was successful. %d", count)
	Response(w, http.StatusCreated, map[string]any{
		"ok":    true,
		"count": count,
		"id":    requestID,
	})

}
