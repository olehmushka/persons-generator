package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection interface {
	Ping(context.Context) error
	Disconnect(context.Context) error
	InsertOne(context.Context, string, string, any, ...*options.InsertOneOptions) (primitive.ObjectID, error)
	InsertMany(ctx context.Context, dbName string, collName string, docs []any, opts ...*options.InsertManyOptions) ([]primitive.ObjectID, error)
	CountDocuments(context.Context, string, string, any, ...*options.CountOptions) (int, error)
	Find(context.Context, string, string, any, ...*options.FindOptions) (Cursor, error)
	FindOne(context.Context, string, string, any, ...*options.FindOneOptions) (SingleResult, error)
	BulkWrite(context.Context, string, string, []WriteModel, ...*options.BulkWriteOptions) (*BulkWriteResult, error)
	UpdateOne(ctx context.Context, dbName string, collName string, filter, update any, opts ...*options.UpdateOptions) (*UpdateResult, error)
	Truncate(ctx context.Context, dbName, collName string) error
	DeleteOne(ctx context.Context, dbName string, collName string, filter any, opts ...*options.DeleteOptions) (int, error)
	DeleteMany(ctx context.Context, dbName string, collName string, filter any, opts ...*options.DeleteOptions) (int, error)
}

type Client interface {
	Connection
}

type Cursor interface {
	ID() int64
	Next(context.Context) bool
	TryNext(context.Context) bool
	Decode(any) error
	Err() error
	Close(context.Context) error
	All(context.Context, any) error
	RemainingBatchLength() int
}

type SingleResult interface {
	Decode(any) error
	Err() error
}

type WriteModel interface {
	mongo.WriteModel
}
