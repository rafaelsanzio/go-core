package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/rafaelsanzio/go-core/pkg/applog"
	"github.com/rafaelsanzio/go-core/pkg/errs"
)

func (s *Store) DeleteOne(ctx context.Context, collection string, id string) errs.AppError {
	col := s.client.Database(dbName).Collection(collection)
	_, err := col.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return errs.ErrMongoDeleteOne.Throwf(applog.Log, errs.ErrFmt, err)
	}

	return nil
}
