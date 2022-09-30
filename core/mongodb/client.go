package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type adapter struct {
	conn Connection
}

func New(username string, password string, uri string, maxBulkItemsSize int) (Client, error) {
	c, err := NewConn(username, password, uri, maxBulkItemsSize)
	if err != nil {
		return nil, err
	}
	return &adapter{
		conn: c,
	}, nil
}

func (a *adapter) Ping(ctx context.Context) error {
	return a.conn.Ping(ctx)
}

func (a *adapter) Disconnect(ctx context.Context) error {
	return a.conn.Disconnect(ctx)
}

func (a *adapter) InsertOne(
	ctx context.Context,
	dbName,
	collName string,
	doc interface{},
	opts ...*options.InsertOneOptions,
) (primitive.ObjectID, error) {
	return a.conn.InsertOne(ctx, dbName, collName, doc, opts...)
}

func (a *adapter) CountDocuments(
	ctx context.Context,
	dbName,
	collName string,
	filter interface{},
	opts ...*options.CountOptions,
) (int, error) {
	return a.conn.CountDocuments(ctx, dbName, collName, filter, opts...)
}

func (a *adapter) Find(ctx context.Context, dbName string, collectionName string,
	filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	return a.conn.Find(ctx, dbName, collectionName, filter, opts...)
}

func (a *adapter) FindOne(ctx context.Context, dbName string, collectionName string,
	filter interface{}, opts ...*options.FindOneOptions) (SingleResult, error) {
	return a.conn.FindOne(ctx, dbName, collectionName, filter, opts...)
}

func (a *adapter) BulkWrite(ctx context.Context, dbName string, collectionName string,
	operations []WriteModel, opts ...*options.BulkWriteOptions) (*BulkWriteResult, error) {

	return a.conn.BulkWrite(ctx, dbName, collectionName, operations, opts...)
}
