/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package mongo // import "github.com/lyf-coder/easy-db/mongo"

import (
	"context"
	"github.com/lyf-coder/easy-db/connect"
	easyOptions "github.com/lyf-coder/easy-db/options"
	"github.com/lyf-coder/easy-db/result"
	"github.com/lyf-coder/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
	"time"
)

// Mongodb is mongo database's struct
type Mongodb struct {
	Client *mongo.Client
	Config *connect.Config
}

// Insert return (insert ID) string,error
// if error,return will be ""
func (mongodb Mongodb) Insert(ctx context.Context, collectionName string, doc interface{}, opts *easyOptions.InsertOneOpts) (string, error) {
	// collectionName's collection
	c := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)
	// insert doc to collection
	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	if opts == nil {
		opts = new(easyOptions.InsertOneOpts)
	}
	r, err := c.InsertOne(ctx, doc, &opts.InsertOneOptions)
	if err != nil {
		return "", err
	}
	return GetObjectIDString(r.InsertedID), err
}

// Inserts return (insert IDs) []string,error
func (mongodb Mongodb) Inserts(ctx context.Context, collectionName string, docs []interface{}, opts *easyOptions.InsertOpts) ([]string, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)
	// insert doc to collection
	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	if opts == nil {
		opts = new(easyOptions.InsertOpts)
	}
	res, err := collection.InsertMany(ctx, docs, &opts.InsertManyOptions)

	if err != nil {
		log.Println(err)
	}
	// handle _id
	var ids []string
	for _, insertID := range res.InsertedIDs {
		ids = append(ids, GetObjectIDString(insertID))
	}
	return ids, err
}

// Find return Entity,error
func (mongodb Mongodb) Find(ctx context.Context, collectionName string, filter interface{}, opts *easyOptions.FindOpts) (*entity.Entity, error) {
	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	// limit 1
	opts.SetLimit(-1)
	findsResult, err := mongodb.Finds(ctx, collectionName, filter, opts)
	if err != nil {
		return nil, err
	}
	if len(findsResult) > 0 {
		return findsResult[0], err
	}
	return nil, err
}

// Finds return []Entity,error
func (mongodb Mongodb) Finds(ctx context.Context, collectionName string, filter interface{}, opts *easyOptions.FindOpts) ([]*entity.Entity, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	if opts == nil {
		opts = new(easyOptions.FindOpts)
	}
	cur, err := collection.Find(ctx, filter, &opts.FindOptions)
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(ctx)

	// store return data
	var results []*entity.Entity

	for cur.Next(ctx) {
		var resultMap map[string]interface{}
		err := cur.Decode(&resultMap)
		if err != nil {
			log.Println(err)
		}
		// covert to entity
		resultEntity := entity.New(resultMap)
		// covert _id from ObjectID type to string type
		resultEntity.Set("_id", GetObjectIDString(resultEntity.Get("_id")))
		results = append(results, resultEntity)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return results, err
}

// Update just update at most one document in the collection and return *result.UpdateResult,error
func (mongodb Mongodb) Update(ctx context.Context, collectionName string, filter interface{}, update interface{}, opts *easyOptions.UpdateOpts) (*result.UpdateResult, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	if opts == nil {
		opts = new(easyOptions.UpdateOpts)
	}

	mUpdateResult, err := collection.UpdateOne(ctx, filter, update, &opts.UpdateOptions)
	if err != nil {
		log.Println(err)
	}

	return &result.UpdateResult{UpdateResult: *mUpdateResult}, err
}

// Updates return *result.UpdateResult,error
func (mongodb Mongodb) Updates(ctx context.Context, collectionName string, filter interface{}, update interface{}, opts *easyOptions.UpdateOpts) (*result.UpdateResult, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	if opts == nil {
		opts = new(easyOptions.UpdateOpts)
	}
	mUpdateResult, err := collection.UpdateMany(ctx, filter, update, &opts.UpdateOptions)
	if err != nil {
		log.Println(err)
	}

	return &result.UpdateResult{UpdateResult: *mUpdateResult}, err
}

//Delete at most one document from the collection.  return DeleteResult,error
func (mongodb Mongodb) Delete(ctx context.Context, collectionName string, filter interface{}, opts *easyOptions.DeleteOpts) (*result.DeleteResult, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	if opts == nil {
		opts = new(easyOptions.DeleteOpts)
	}
	mDeleteResult, err := collection.DeleteOne(ctx, filter, &opts.DeleteOptions)
	if err != nil {
		log.Println(err)
	}
	return &result.DeleteResult{DeleteResult: *mDeleteResult}, err
}

// DeletesWithOptions return *result.DeleteResult,error
func (mongodb Mongodb) Deletes(ctx context.Context, collectionName string, filter interface{}, opts *easyOptions.DeleteOpts) (*result.DeleteResult, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	if opts == nil {
		opts = new(easyOptions.DeleteOpts)
	}

	mDeleteResult, err := collection.DeleteMany(ctx, filter, &opts.DeleteOptions)
	if err != nil {
		log.Println(err)
	}

	return &result.DeleteResult{DeleteResult: *mDeleteResult}, err
}

// Count return int64,error
func (mongodb Mongodb) Count(ctx context.Context, collectionName string, filter interface{}, opts *easyOptions.CountOpts) (int64, error) {
	// collectionName's collection
	collection := mongodb.Client.Database(mongodb.Config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx, cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}
	if opts == nil {
		opts = new(easyOptions.CountOpts)
	}

	val, err := collection.CountDocuments(ctx, filter, &opts.CountOptions)
	if err != nil {
		log.Println(err)
	}
	return val, err
}

// Close Db connect
func (mongodb Mongodb) Close(ctx context.Context) error {
	err := mongodb.Client.Disconnect(ctx)
	if err != nil {
		log.Println(err)
	}
	return err
}

// getUri is get connect mongodb uri func
func getUri(config *connect.Config) string {
	var buf strings.Builder
	buf.WriteString("mongodb://")
	buf.WriteString(config.UserName)
	buf.WriteString(":")
	buf.WriteString(config.Password)
	buf.WriteString("@")
	buf.WriteString(config.Host)
	buf.WriteString(":")
	buf.WriteString(config.Port)
	buf.WriteString("/")
	buf.WriteString(config.DatabaseName)
	buf.WriteString("?")

	// handle options
	for option := range config.Options {
		buf.WriteString(option)
		buf.WriteString("=")
		buf.WriteString(config.Options[option])
		buf.WriteString("&")
	}

	return buf.String()
}

// New create Mongodb struct instance
func New(config *connect.Config) *Mongodb {
	// define a 10s context
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	// connect db to get client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getUri(config)))

	if err != nil {
		log.Println(err)
	}
	mongodb := Mongodb{Client: client, Config: config}
	return &mongodb
}

// getObjectIDString
func GetObjectIDString(id interface{}) string {
	if v, ok := id.(primitive.ObjectID); ok {
		return v.Hex()
	}
	return ""
}
