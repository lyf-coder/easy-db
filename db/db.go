/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package db // import "github.com/lyf-coder/easy-db/db"

import (
	"context"
	"github.com/lyf-coder/easy-db/connect"
	"github.com/lyf-coder/easy-db/mongo"
	"github.com/lyf-coder/easy-db/options"
	"github.com/lyf-coder/easy-db/result"
	"github.com/lyf-coder/entity"
)
// store Db - databaseName is key
var dbMap = make(map[string]Db)

// New Db by config
func New(config *connect.Config) Db{
	var db Db
	// first find db in dbMap by DatabaseName
	db = dbMap[config.DatabaseName]
	// find
	if db!=nil{
		return db
	}
	// not find in dbMap - New
	switch config.DbType {
	case connect.MONGODB:
		// init mongodb
		db = mongo.New(config)
	}
	dbMap[config.DatabaseName] = db
	return db
}

// Db
type Db interface {
	// Close Db connect
	Close(ctx context.Context) error

	// Insert return id
	Insert(ctx context.Context,collectionName string,doc interface{}) (string,error)
	// InsertWithOptions return id
	InsertWithOptions(ctx context.Context,collectionName string,doc interface{},opts *options.InsertOpts) (string,error)

	// Inserts return ids
	Inserts(ctx context.Context,collectionName string,docs []interface{}) ([]string,error)
	// InsertsWithOptions return ids
	InsertsWithOptions(ctx context.Context,collectionName string,docs []interface{} , opts *options.InsertOpts) ([]string,error)

	// Find return Entity
	Find(ctx context.Context,collectionName string, filter interface{}, opts options.FindOpts) (*entity.Entity,error)
	// Find return []Entity
	Finds(ctx context.Context,collectionName string, filter interface{}, opts options.FindOpts) ([]*entity.Entity,error)

	// Update just update at most one document in the collection and return *result.UpdateResult,error
	Update(ctx context.Context,collectionName string,filter interface{}, update interface{})(*result.UpdateResult,error)
	// UpdateWithOptions just update at most one document in the collection and return *result.UpdateResult,error
	UpdateWithOptions(ctx context.Context,collectionName string,filter interface{}, update interface{}, opts options.UpdateOpts)(*result.UpdateResult,error)

	// Updates return *result.UpdateResult,error
	Updates(ctx context.Context,collectionName string,filter interface{}, update interface{}) (*result.UpdateResult,error)
	// UpdatesWithOptions return *result.UpdateResult,error
	UpdatesWithOptions(ctx context.Context,collectionName string,filter interface{}, update interface{},opts options.UpdateOpts) (*result.UpdateResult,error)

	//Delete at most one document from the collection.  return DeleteResult,error
	Delete(ctx context.Context,collectionName string,filter interface{}) (*result.DeleteResult,error)
	//DeleteWithOptions at most one document from the collection.  return DeleteResult,error
	DeleteWithOptions(ctx context.Context,collectionName string,filter interface{},opts options.DeleteOpts) (*result.DeleteResult,error)

	// Deletes return *result.DeleteResult,error
	Deletes(ctx context.Context,collectionName string,filter interface{}) (*result.DeleteResult,error)
	// DeletesWithOptions return *result.DeleteResult,error
	DeletesWithOptions(ctx context.Context,collectionName string,filter interface{},opts options.DeleteOpts) (*result.DeleteResult,error)
}