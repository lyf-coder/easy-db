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
	client   *mongo.Client
	config *connect.Config
}

// Insert return (insert ID) string,error  if error,return will be ""
func (mongodb Mongodb) Insert(ctx context.Context, collectionName string,doc interface{}) (string,error){
	insertIDs,err := mongodb.Inserts(ctx, collectionName, []interface{}{doc})
	if err!=nil{
		return "",err
	}
	return insertIDs[0],err
}
// InsertWithOptions return (insert ID) string,error if error,return will be ""
func (mongodb Mongodb) InsertWithOptions(ctx context.Context, collectionName string,doc interface{} ,  opts *easyOptions.InsertOpts) (string,error){
	insertIDs,err := mongodb.InsertsWithOptions(ctx, collectionName,[]interface{}{doc}, opts)
	if err!=nil{
		return "",err
	}
	return insertIDs[0],err
}

// Inserts return (insert IDs) []string,error
func (mongodb Mongodb) Inserts(ctx context.Context,collectionName string,docs []interface{}) ([]string,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)
	// insert doc to collection
	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	res, err := collection.InsertMany(ctx, docs)

	if err != nil {
		log.Println(err)
	}
	// handle _id
	var ids []string
	for _,insertID := range res.InsertedIDs{
		ids = append(ids, getObjectIDString(insertID))
	}
	return ids,err
}
// InsertsWithOptions return (insert IDs) []string,error
func (mongodb Mongodb) InsertsWithOptions(ctx context.Context,collectionName string,docs []interface{} ,  opts *easyOptions.InsertOpts) ([]string,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)
	// insert doc to collection
	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	res, err := collection.InsertMany(ctx, docs, &opts.InsertManyOptions)

	if err != nil {
		log.Println(err)
	}
	// handle _id
	var ids []string
	for _,insertID := range res.InsertedIDs{
		ids = append(ids, getObjectIDString(insertID))
	}
	return ids,err
}


// Find return Entity,error
func (mongodb Mongodb) Find(ctx context.Context,collectionName string, filter interface{}, opts easyOptions.FindOpts) (*entity.Entity,error){
	// limit 1
	opts.SetLimit(-1)
	findsResult,err := mongodb.Finds(ctx, collectionName, filter, opts)
	if err!=nil{
		return new(entity.Entity),err
	}
	return findsResult[0],err
}

// Finds return []Entity,error
func (mongodb Mongodb) Finds(ctx context.Context,collectionName string, filter interface{}, opts easyOptions.FindOpts) ([]*entity.Entity,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
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
		resultEntity.Set("_id", getObjectIDString(resultEntity.Get("_id")))
		results = append(results, resultEntity)
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	return results,err
}

// Update just update at most one document in the collection and return *result.UpdateResult,error
func (mongodb Mongodb) Update(ctx context.Context,collectionName string,filter interface{}, update interface{})(*result.UpdateResult,error) {
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mUpdateResult,err := collection.UpdateOne(ctx, filter, update)
	if err!=nil{
		log.Println(err)
	}
	updateResult := result.UpdateResult{}
	updateResult.UpdateResult = *mUpdateResult

	return &updateResult,err
}

// UpdateWithOptions just update at most one document in the collection and return *result.UpdateResult,error
func (mongodb Mongodb) UpdateWithOptions(ctx context.Context,collectionName string,filter interface{}, update interface{}, opts easyOptions.UpdateOpts)(*result.UpdateResult,error) {
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mUpdateResult,err := collection.UpdateOne(ctx, filter, update, &opts.UpdateOptions)
	if err!=nil{
		log.Println(err)
	}
	updateResult := result.UpdateResult{}
	updateResult.UpdateResult = *mUpdateResult

	return &updateResult,err
}

// Updates return *result.UpdateResult,error
func (mongodb Mongodb) Updates(ctx context.Context,collectionName string,filter interface{}, update interface{}) (*result.UpdateResult,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mUpdateResult,err := collection.UpdateMany(ctx, filter, update)
	if err!=nil{
		log.Println(err)
	}
	updateResult := result.UpdateResult{}
	updateResult.UpdateResult = *mUpdateResult

	return &updateResult,err
}

// UpdatesWithOptions return *result.UpdateResult,error
func (mongodb Mongodb) UpdatesWithOptions(ctx context.Context,collectionName string,filter interface{}, update interface{},opts easyOptions.UpdateOpts) (*result.UpdateResult,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mUpdateResult,err := collection.UpdateMany(ctx, filter, update, &opts.UpdateOptions)
	if err!=nil{
		log.Println(err)
	}
	updateResult := result.UpdateResult{}
	updateResult.UpdateResult = *mUpdateResult

	return &updateResult,err
}

//Delete at most one document from the collection.  return DeleteResult,error
func (mongodb Mongodb) Delete(ctx context.Context,collectionName string,filter interface{}) (*result.DeleteResult,error) {
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mDeleteResult,err := collection.DeleteOne(ctx, filter)
	if err!=nil{
		log.Println(err)
	}
	deleteResult := result.DeleteResult{}
	deleteResult.DeleteResult = *mDeleteResult

	return &deleteResult,err
}
//DeleteWithOptions at most one document from the collection.  return DeleteResult,error
func (mongodb Mongodb) DeleteWithOptions(ctx context.Context,collectionName string,filter interface{},opts easyOptions.DeleteOpts) (*result.DeleteResult,error) {
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mDeleteResult,err := collection.DeleteOne(ctx, filter,&opts.DeleteOptions)
	if err!=nil{
		log.Println(err)
	}
	deleteResult := result.DeleteResult{}
	deleteResult.DeleteResult = *mDeleteResult

	return &deleteResult,err
}

// Deletes return *result.DeleteResult,error
func (mongodb Mongodb) Deletes(ctx context.Context,collectionName string,filter interface{}) (*result.DeleteResult,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mDeleteResult,err := collection.DeleteMany(ctx, filter)
	if err!=nil{
		log.Println(err)
	}
	deleteResult := result.DeleteResult{}
	deleteResult.DeleteResult = *mDeleteResult

	return &deleteResult,err
}
// DeletesWithOptions return *result.DeleteResult,error
func (mongodb Mongodb) DeletesWithOptions(ctx context.Context,collectionName string,filter interface{},opts easyOptions.DeleteOpts) (*result.DeleteResult,error){
	// collectionName's collection
	collection := mongodb.client.Database(mongodb.config.DatabaseName).Collection(collectionName)

	if ctx == nil {
		var cancelFunc context.CancelFunc
		ctx,cancelFunc = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
	}

	mDeleteResult,err := collection.DeleteMany(ctx, filter, &opts.DeleteOptions)
	if err!=nil{
		log.Println(err)
	}
	deleteResult := result.DeleteResult{}
	deleteResult.DeleteResult = *mDeleteResult

	return &deleteResult,err
}
// Close Db connect
func (mongodb Mongodb) Close(ctx context.Context) error{
	err := mongodb.client.Disconnect(ctx)
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
	mongodb := Mongodb{client:client, config:config}
	return &mongodb
}

// getObjectIDString
func getObjectIDString(id interface{}) string {
	if v, ok := id.(primitive.ObjectID); ok {
		return v.Hex()
	}
	return ""
}
