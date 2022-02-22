package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rafaelsanzio/go-core/pkg/api"
	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
	"github.com/rafaelsanzio/go-core/pkg/store"
)

func main() {
	log.Println("starting up...")
	store.GetStore()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		_ = errs.ErrMongoConnect.Throwf(applog.Log, errs.ErrFmt, err)
	}

	ctx, cancFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancFunc()

	err = client.Connect(ctx)
	if err != nil {
		_ = errs.ErrMongoConnect.Throwf(applog.Log, errs.ErrFmt, err)
	}

	defer func() {
		err = client.Disconnect(ctx)
		if err != nil {
			_ = errs.ErrMongoConnect.Throwf(applog.Log, errs.ErrFmt, err)
		}
	}()

	log.Println("MongoDB server is healthy.")
	log.Fatal(http.ListenAndServe(":8000", api.NewRouter()))
}
